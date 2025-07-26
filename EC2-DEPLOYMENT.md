# EC2 Deployment Guide

## Prerequisites

### EC2 Instance Requirements

- **Instance Type**: t3.medium or larger (2 vCPU, 4GB RAM minimum)
- **OS**: Amazon Linux 2 or Ubuntu 20.04+
- **Security Group**: Open ports 22 (SSH), 80 (HTTP), 443 (HTTPS), 3000 (Node.js), 8086 (Go)
- **Storage**: 20GB+ EBS volume

### Required Software

- Docker
- Docker Compose
- Git

## Step-by-Step Deployment

### 1. Launch and Configure EC2 Instance

```bash
# Update system
sudo yum update -y  # Amazon Linux
# OR
sudo apt update && sudo apt upgrade -y  # Ubuntu

# Install Docker
sudo yum install -y docker  # Amazon Linux
# OR
sudo apt install -y docker.io  # Ubuntu

# Start Docker service
sudo systemctl start docker
sudo systemctl enable docker

# Add user to docker group
sudo usermod -a -G docker $USER
newgrp docker

# Install Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Install Git
sudo yum install -y git  # Amazon Linux
# OR
sudo apt install -y git  # Ubuntu
```

### 2. Clone and Setup Project

```bash
# Clone the repository
git clone <your-repo-url>
cd The-Containerized-Intelligence-Pipeline

# Create production environment file
cp env.example .env
```

### 3. Configure Environment Variables

Edit the `.env` file with production values:

```bash
nano .env
```

**Production `.env` Configuration:**

```bash
# Database Configuration
POSTGRES_USER=postgres
POSTGRES_PASSWORD=your_super_secure_password_123!
POSTGRES_DB=logs
POSTGRES_HOST=postgres-db
POSTGRES_PORT=5432

# Node.js Server Configuration
NODE_ENV=production
NODE_PORT=3000
GO_SERVER_HOST=go-server
GO_SERVER_PORT=8086

# Go Server Configuration
GO_PORT=8086
GIN_MODE=release

# Security (Generate strong values)
JWT_SECRET=your_jwt_secret_256_bit_key_here
API_KEY=your_api_key_here
```

**⚠️ Security Best Practices:**

1. **Strong Passwords**: Use complex passwords (20+ characters)
2. **Secret Generation**:

   ```bash
   # Generate strong JWT secret
   openssl rand -base64 32

   # Generate API key
   openssl rand -hex 32
   ```

3. **File Permissions**: Restrict access to .env file
   ```bash
   chmod 600 .env
   ```

### 4. Deploy Application

```bash
# Build and start services
docker-compose -f docker-compose.prod.yml --env-file .env up --build -d

# Check status
docker-compose -f docker-compose.prod.yml ps

# View logs
docker-compose -f docker-compose.prod.yml logs -f
```

### 5. Configure Reverse Proxy (Optional but Recommended)

Install Nginx for reverse proxy:

```bash
# Install Nginx
sudo yum install -y nginx  # Amazon Linux
# OR
sudo apt install -y nginx  # Ubuntu

# Create Nginx configuration
sudo tee /etc/nginx/conf.d/microservices.conf << 'EOF'
server {
    listen 80;
    server_name your-domain.com;  # Replace with your domain

    # Node.js API
    location /api/ {
        proxy_pass http://localhost:3000/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # Go Server API
    location /compute/ {
        proxy_pass http://localhost:8086/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # Health checks
    location /health {
        proxy_pass http://localhost:3000/health;
    }
}
EOF

# Test and start Nginx
sudo nginx -t
sudo systemctl start nginx
sudo systemctl enable nginx
```

### 6. SSL/HTTPS Setup (Recommended)

```bash
# Install Certbot for Let's Encrypt
sudo yum install -y certbot python3-certbot-nginx  # Amazon Linux
# OR
sudo apt install -y certbot python3-certbot-nginx  # Ubuntu

# Get SSL certificate
sudo certbot --nginx -d your-domain.com

# Auto-renewal
sudo crontab -e
# Add: 0 12 * * * /usr/bin/certbot renew --quiet
```

## Environment Variable Best Practices

### 1. Environment File Structure

```bash
# .env files hierarchy
.env                    # Production secrets (never commit)
.env.example           # Template with dummy values (commit this)
.env.local             # Local development overrides
.env.staging           # Staging environment
```

### 2. Docker Compose Environment Loading

```yaml
# docker-compose.prod.yml
services:
  nodejs-server:
    env_file:
      - .env
    environment:
      - NODE_ENV=${NODE_ENV}
      - DATABASE_URL=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}
```

### 3. Application Environment Loading

**Node.js** (using dotenv package):

```javascript
// Add to package.json
"dependencies": {
  "dotenv": "^16.0.0"
}

// Add to server.js (top of file)
require('dotenv').config();

// Usage
const dbConfig = {
  user: process.env.POSTGRES_USER,
  password: process.env.POSTGRES_PASSWORD,
  host: process.env.POSTGRES_HOST,
  port: process.env.POSTGRES_PORT,
  database: process.env.POSTGRES_DB,
};
```

**Go** (using environment variables):

```go
import "os"

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}

// Usage
port := getEnv("GO_PORT", "8086")
dbHost := getEnv("POSTGRES_HOST", "localhost")
```

### 4. Security Considerations

1. **Never commit `.env` files**

   ```bash
   # Add to .gitignore
   .env
   .env.local
   .env.production
   ```

2. **Use AWS Systems Manager Parameter Store** (Advanced):

   ```bash
   # Store secrets in AWS Parameter Store
   aws ssm put-parameter --name "/app/postgres/password" --value "secret" --type "SecureString"

   # Retrieve in application
   aws ssm get-parameter --name "/app/postgres/password" --with-decryption
   ```

3. **Environment Variable Validation**:

   ```javascript
   // Node.js validation
   const requiredEnvVars = ["POSTGRES_PASSWORD", "JWT_SECRET"];
   const missingVars = requiredEnvVars.filter(
     (varName) => !process.env[varName]
   );

   if (missingVars.length > 0) {
     console.error("Missing required environment variables:", missingVars);
     process.exit(1);
   }
   ```

## Monitoring and Maintenance

### 1. Health Checks

```bash
# Check application health
curl http://your-ec2-ip:3000/health
curl http://your-ec2-ip:8086/health

# Check database
docker exec postgres-db pg_isready -U postgres
```

### 2. Logs Management

```bash
# View real-time logs
docker-compose -f docker-compose.prod.yml logs -f

# View specific service logs
docker-compose -f docker-compose.prod.yml logs nodejs-server

# Log rotation (add to crontab)
0 0 * * * docker system prune -f
```

### 3. Backup Database

```bash
# Create backup script
#!/bin/bash
BACKUP_DIR="/home/ec2-user/backups"
DATE=$(date +%Y%m%d_%H%M%S)

docker exec postgres-db pg_dump -U postgres logs > "${BACKUP_DIR}/logs_backup_${DATE}.sql"

# Keep only last 7 days of backups
find ${BACKUP_DIR} -name "logs_backup_*.sql" -mtime +7 -delete
```

### 4. Auto-scaling Considerations

For production scaling:

- Use Application Load Balancer (ALB)
- Auto Scaling Groups
- ECS or EKS for container orchestration
- RDS for managed database
- ElastiCache for caching

## API Usage Examples

```bash
# Health check
curl https://your-domain.com/health

# Calculate endpoint
curl -X POST https://your-domain.com/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"data": [10, 20, 30, 40, 50]}'

# Direct Go server access
curl https://your-domain.com/compute/health
```

## Troubleshooting

### Common Issues

1. **Container fails to start**:

   ```bash
   docker-compose -f docker-compose.prod.yml logs [service-name]
   ```

2. **Environment variables not loading**:

   ```bash
   docker exec [container-name] env | grep POSTGRES
   ```

3. **Database connection issues**:

   ```bash
   docker exec nodejs-server ping postgres-db
   ```

4. **Port conflicts**:
   ```bash
   sudo netstat -tulpn | grep :3000
   ```

This deployment guide ensures your microservices run securely and efficiently on EC2 with proper environment variable management!
