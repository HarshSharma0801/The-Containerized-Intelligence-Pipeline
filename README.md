# 🚀 The Containerized Intelligence Pipeline

## 📋 Scenario

As a DevOps engineer in a modern tech company, you're tasked with deploying a scalable microservices architecture that demonstrates containerization, service orchestration, and database persistence. This project showcases a complete intelligence pipeline with Node.js Express, Go Gin, and PostgreSQL working together in a containerized environment.

## 🎯 Core Learning Objectives

- Containerization fundamentals with Docker
- Microservices architecture design
- Service orchestration with Docker Compose
- Database persistence and volume management
- Health checks and service dependencies
- Environment-based configuration management
- Cross-service communication patterns

## 🛠️ Tech Stack & Rationale

| Technology                                                                                                     | Purpose             | Rationale                                                         |
| -------------------------------------------------------------------------------------------------------------- | ------------------- | ----------------------------------------------------------------- |
| ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)                   | Containerization    | Industry standard for application containerization and deployment |
| ![Docker Compose](https://img.shields.io/badge/Docker%20Compose-2496ED?style=flat&logo=docker&logoColor=white) | Orchestration       | Simplifies multi-container application management                 |
| ![Node.js](https://img.shields.io/badge/Node.js-339933?style=flat&logo=node.js&logoColor=white)                | API Gateway         | Fast, event-driven runtime perfect for API orchestration          |
| ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)                               | Computation Service | High-performance language ideal for CPU-intensive calculations    |
| ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=flat&logo=postgresql&logoColor=white)       | Database            | Robust, ACID-compliant database with excellent Docker support     |
| ![Gin](https://img.shields.io/badge/Gin-00ADD8?style=flat&logo=go&logoColor=white)                             | Go Web Framework    | Lightweight, fast HTTP framework for Go services                  |
| ![Express](https://img.shields.io/badge/Express-000000?style=flat&logo=express&logoColor=white)                | Node.js Framework   | Minimal, flexible web application framework                       |
| ![GitHub Actions](https://img.shields.io/badge/GitHub%20Actions-2088FF?style=flat&logo=github-actions&logoColor=white) | CI/CD Pipeline       | Tight integration with GitHub, simplest way to create CI/CD pipelines |

## 📋 Implementation Steps

### Step 1: Environment Configuration Setup

Set up environment variables for secure configuration management

<img width="940" height="460" alt="Screenshot 2025-07-27 at 6 55 41 PM" src="https://github.com/user-attachments/assets/9e07ec9f-f5fa-4f3b-8d89-5522bb861ed3" />

```bash
# Copy environment template
cp env.example .env

# Configure your environment variables
# - Database credentials
# - Service ports
# - Security settings
```

### Step 2: Docker Compose File 

Configure `docker-compose.yml` for all the services configure PostgreSQL with persistent volume and initialization scripts in `database` directory 

<img width="940" height="753" alt="Screenshot 2025-07-27 at 6 56 03 PM" src="https://github.com/user-attachments/assets/204f0c41-fbc5-4c3c-856d-dfa7edb4f5a6" />


```sql
-- Automatic table creation on container startup
CREATE TABLE IF NOT EXISTS process_logs (
    id SERIAL PRIMARY KEY,
    time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    processing_time INTERVAL NOT NULL
);
```

### Step 3: Spinning up Ec2 (Large Size Recommended for High Computation) 

Spin up Ec2 on AWS and install various elements as root user

1. `Docker`
2. `git`
3. `docker-compose` 

```bash
# Update system
yum update -y        # For Amazon Linux
# OR
apt update && apt upgrade -y  # For Ubuntu

# Install Docker
yum install docker -y
# OR
apt install docker.io -y

# Start Docker
systemctl start docker
systemctl enable docker

# Install docker-compose
curl -L "https://github.com/docker/compose/releases/download/2.24.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

# Install git
yum install docker -y

# Confirm installation
docker --version
docker-compose --version

```

### Step 4: Attach , Mount and Fomrat EBS Volume for Database for Ec2

Mount and Format attached EBS Volume onto Ec2 on a directory for database storage example below this volume xvdb is mounted on /mnt/xvdb directory

```bash
# Commands

# Check the Attached Volume
df -h

# Format the volume attached on Ec2
sudo mkfs.ext4 /dev/xvdb 

# Mount on a directory
sudo mkdir -p /mnt/xvdb
sudo mount /dev/xvdb /mnt/xvdb
```

<img width="855" height="199" alt="Screenshot 2025-07-27 at 7 22 25 PM" src="https://github.com/user-attachments/assets/dc37db18-4987-4b14-b24e-b99171463b21" />

### Step 5: Create deploy User and setup Permissions

create a deploy user and setup permission as it will be responsible for deploying (good practice you can do it via root user too)

```bash

# Add a user
sudo adduser deploy

# Add user in user groups for docker and wheel (run sudo commands)
usermod -aG docker deploy
usermod -aG wheel deploy

# login deploy
su deploy

```

setup ssh key

```bash
# Generate ssh key
ssh-keygen
cd ~/.ssh
```

In `authorized_keys` remove the restrictions related to ssh 

give permission to postgres service (you can find the id once its created on docker-compose) for mounted directory example (/mnt/xvdb/postgres-data) so that it can access and write into it


### Step 6: Clone git repository in deploy user and setup Env 

```bash
# Login as deploy user
su deploy

# Go to ~ directory and clone repository
cd ~ 
git clone https://github.com/HarshSharma0801/The-Containerized-Intelligence-Pipeline.git

# Cd into Repository and setup env

cd The-Containerized-Intelligence-Pipeline 
vim env 
```

## Step 7: Start The Containers and verify

```bash

# start docker compose 
docker-compose up -d --build 

# Verify using 
docker ps
```

<img width="1017" height="133" alt="Screenshot 2025-07-27 at 7 41 37 PM" src="https://github.com/user-attachments/assets/bacee4e7-beb6-42da-83e8-792eb687d60a" />


## Step 7: Setup CI/CD Pipeline

1. Create `./github/workflows/deploy.yml`

<img width="1017" height="701" alt="Screenshot 2025-07-27 at 7 43 09 PM" src="https://github.com/user-attachments/assets/a4841dc8-03db-43eb-af8d-005f114ec25a" />

2. Drive Github Secrets 

`EC2_SSH_KEY` -> private ssh key for deploy user found in .ssh folder in deploy user (needed for ssh into ec2 as deploy user without password)
`EC2_HOST` -> ec2 Host found in ec2 dashboard
`EC2_USER` -> deploy
`ENV_FILE` -> check the env file in `.env.example`

3. Push to main branch to see the actions tab 

## 🏗️ Architecture Overview

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Client/User   │    │  Node.js API    │    │  Go Compute     │
│                 │────│  Gateway        │────│  Service        │
│  Port: ANY      │    │  Port: 3000     │    │  Port: 8086     │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                              │
                              ▼
                       ┌─────────────────┐
                       │  PostgreSQL     │
                       │  Database       │
                       │  Port: 5432     │
                       └─────────────────┘
```

## 📊 API Endpoints

### Health Checks

```bash
# Node.js service health
GET http://localhost:3000/health

# Go service health
GET http://localhost:8086/health
```

### Main Pipeline Endpoint

```bash
# Trigger computation pipeline
GET http://localhost:3000/calculate
```

**Response:**

```json
{
  "result": {
    "time": 0.05234,
    "operation": "prime_calculation",
    "processedAt": "2024-01-15T10:30:00Z"
  },
  "processingTime": 150,
  "timestamp": "2024-01-15T10:30:00.000Z"
}
```

## 🗃️ Database Schema

```sql
CREATE TABLE process_logs (
    id SERIAL PRIMARY KEY,
    time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    processing_time INTERVAL NOT NULL
);
```

## 🔧 Service Configuration

### Environment Variables

| Variable            | Default  | Description                 |
| ------------------- | -------- | --------------------------- |
| `NODE_PORT`         | 3000     | Node.js API Gateway port    |
| `GO_PORT`           | 8086     | Go computation service port |
| `POSTGRES_USER`     | postgres | Database username           |
| `POSTGRES_PASSWORD` | password | Database password           |
| `POSTGRES_DB`       | logs     | Database name               |
| `POSTGRES_PORT`     | 5432     | Database port               |

### Docker Services

- **nodejs-server**: http://localhost:3000
- **go-server**: http://localhost:8086
- **postgres-db**: localhost:5432

## 💾 Data Persistence

PostgreSQL data is persisted using Docker volumes:

- **Volume Mount**: `/mnt/xvdb/postgres-data:/var/lib/postgresql/data`
- **Init Scripts**: `./database/init.sql` automatically executed on first run

## 🔄 Service Dependencies

```yaml
nodejs-server:
  depends_on:
    postgres-db:
      condition: service_healthy
    go-server:
      condition: service_healthy
```

## 🏥 Health Monitoring

All services include health checks:

- **Interval**: 30 seconds
- **Timeout**: 10 seconds
- **Retries**: 3 attempts
- **Restart Policy**: unless-stopped

## 🛑 Managing the Deployment

### Start Services

```bash
docker-compose up --build
```

### Stop Services

```bash
docker-compose down
```

### Remove Everything (including volumes)

```bash
docker-compose down -v
```

### View Logs

```bash
# All services
docker-compose logs

# Specific service
docker-compose logs nodejs-server
docker-compose logs go-server
docker-compose logs postgres-db
```

### Scale Services

```bash
# Scale Go service for higher computation load
docker-compose up --scale go-server=3
```

## 🔍 Troubleshooting

### Common Issues

1. **Port conflicts**: Ensure ports 3000, 8086, and 5432 are available
2. **Environment variables**: Verify `.env` file exists and is properly configured
3. **Docker permissions**: Ensure Docker daemon is running and accessible
4. **Volume permissions**: Check filesystem permissions for PostgreSQL data directory

### Debug Commands

```bash
# Check container status
docker-compose ps

# Inspect container logs
docker-compose logs -f [service-name]

# Execute commands in running containers
docker-compose exec nodejs-server sh
docker-compose exec go-server sh
docker-compose exec postgres-db psql -U postgres -d logs
```

This containerized intelligence pipeline demonstrates modern DevOps practices with microservices, containerization, and automated orchestration - perfect for learning deployment fundamentals and scaling strategies.
