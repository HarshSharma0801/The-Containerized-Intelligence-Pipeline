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

## 📋 Implementation Steps

### Step 1: Environment Configuration Setup

Set up environment variables for secure configuration management

<img width="860" height="597" alt="Environment Setup" src="https://github.com/user-attachments/assets/9d6ccd31-aceb-4e36-9518-87c4a4028b73" />

```bash
# Copy environment template
cp env.example .env

# Configure your environment variables
# - Database credentials
# - Service ports
# - Security settings
```

### Step 2: Database Initialization

Configure PostgreSQL with persistent volume and initialization scripts

<img width="983" height="459" alt="Database Setup" src="https://github.com/user-attachments/assets/a65b742e-fcdd-41ed-ac57-c82696684288" />

```sql
-- Automatic table creation on container startup
CREATE TABLE IF NOT EXISTS process_logs (
    id SERIAL PRIMARY KEY,
    time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    processing_time INTERVAL NOT NULL
);
```

### Step 3: Go Computation Service

Build and deploy the high-performance Go service for prime calculations

<img width="860" height="658" alt="Go Service Deployment" src="https://github.com/user-attachments/assets/49f27a90-df8e-4faa-a1da-cb9f85fda3b2" />

```bash
# Go service handles CPU-intensive computations
# - Concurrent prime number calculations
# - Health check endpoints
# - Performance metrics
```

### Step 4: Node.js API Gateway

Deploy the Express server as the main entry point and orchestrator

<img width="860" height="597" alt="Node.js Gateway" src="https://github.com/user-attachments/assets/0267e2d4-1ff7-4880-86e6-dae6938383bf" />

```bash
# Node.js service coordinates:
# - API request handling
# - Service-to-service communication
# - Database logging
# - Response aggregation
```

### Step 5: Container Orchestration

Launch the complete microservices stack with Docker Compose

<img width="860" height="597" alt="Container Orchestration" src="https://github.com/user-attachments/assets/5264b737-e3a0-45f9-bb45-83892848df97" />

```bash
# Start all services with dependencies
docker-compose up --build
```

### Step 6: Service Health Verification

Verify all services are running and communicating properly

<img width="860" height="597" alt="Health Checks" src="https://github.com/user-attachments/assets/06028d53-4803-40cb-8feb-8664985c0df0" />

```bash
# Check all service endpoints
curl http://localhost:3000/health
curl http://localhost:8086/health
```

## 🚀 Quick Start

```bash
# 1. Clone and navigate to project
git clone <repository-url>
cd The-Containerized-Intelligence-Pipeline

# 2. Set up environment
cp env.example .env
# Edit .env with your preferred settings

# 3. Deploy the complete stack
docker-compose up --build

# 4. Verify deployment
curl http://localhost:3000/health
curl http://localhost:8086/health

# 5. Test the pipeline
curl http://localhost:3000/calculate
```

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
