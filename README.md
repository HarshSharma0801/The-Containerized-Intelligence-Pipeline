# Containerized Intelligence Pipeline

A microservices architecture with Node.js Express, Go Gin, and PostgreSQL database.

## Architecture

- **Node.js Express Server** (Port 3000): Entry point that receives requests and calls the Go server
- **Go Gin Server** (Port 8080): Performs calculations and returns results
- **PostgreSQL Database** (Port 5432): Stores process logs with persistent volume

## Project Structure

```
├── nodejs-server/
│   ├── server.js
│   ├── package.json
│   └── Dockerfile
├── go-server/
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── database/
│   └── init.sql
└── docker-compose.yml
```

## Setup and Running

1. **Start all services:**

   ```bash
   docker-compose up --build
   ```

2. **Check service health:**

   ```bash
   # Node.js server
   curl http://localhost:3000/health

   # Go server
   curl http://localhost:8080/health
   ```

## API Usage

### Main Calculation Endpoint

**POST** `http://localhost:3000/calculate`

```json
{
  "data": [1, 2, 3, 4, 5],
  "processNumber": 123
}
```

**Response:**

```json
{
  "processNumber": 123,
  "result": 1.5811388300841898,
  "processingTime": 150,
  "timestamp": "2024-01-15T10:30:00.000Z"
}
```

### Example Usage

```bash
# Send calculation request
curl -X POST http://localhost:3000/calculate \
  -H "Content-Type: application/json" \
  -d '{"data": [10, 20, 30, 40, 50], "processNumber": 456}'
```

## How It Works

1. Node.js server receives a calculation request
2. Makes HTTP call to Go server with the data
3. Go server performs standard deviation calculation
4. Results are returned to Node.js server
5. Process time and number are logged to PostgreSQL database
6. Final response sent back to client

## Database Schema

```sql
CREATE TABLE process_logs (
  id SERIAL PRIMARY KEY,
  process_number SERIAL NOT NULL,
  time TIMESTAMP NOT NULL,
  processing_time_ms INTEGER NOT NULL
);
```

## Services

- **nodejs-server**: http://localhost:3000
- **go-server**: http://localhost:8080
- **postgres-db**: localhost:5432

## Volume

PostgreSQL data is persisted in the `postgres_data` Docker volume.

## Stopping Services

```bash
docker-compose down
```

To remove volumes as well:

```bash
docker-compose down -v
```
