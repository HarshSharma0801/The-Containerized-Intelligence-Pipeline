require("dotenv").config();
const express = require("express");
const axios = require("axios");
const { Pool } = require("pg");

const app = express();
const port = process.env.NODE_PORT || 3000;

// Database configuration from environment variables
const pool = new Pool({
  user: process.env.POSTGRES_USER || "postgres",
  host: process.env.POSTGRES_HOST || "postgres-db",
  database: process.env.POSTGRES_DB || "logs",
  password: process.env.POSTGRES_PASSWORD || "password",
  port: process.env.POSTGRES_PORT || 5432,
});

// Go server configuration
const GO_SERVER_HOST = process.env.GO_SERVER_HOST || "go-server";
const GO_SERVER_PORT = process.env.GO_SERVER_PORT || 8086;

app.use(express.json());

app.get("/health", (req, res) => {
  res.json({ status: "Node.js server is running" });
});

app.get("/calculate", async (req, res) => {
  const startTime = Date.now();

  try {
    console.log(`Starting calculation process`);

    const response = await axios.get(
      `http://${GO_SERVER_HOST}:${GO_SERVER_PORT}/compute`
    );

    const endTime = Date.now();
    const processingTime = endTime - startTime;

    const result = await pool.query(
      "INSERT INTO process_logs (time, processing_time) VALUES ($1, $2) RETURNING process_number",
      [new Date(), String(response.data.time)]
    );

    const processNumber = result.rows[0].process_number;
    console.log(`Process ${processNumber} completed in ${processingTime}ms`);

    res.json({
      processNumber: processNumber,
      result: response.data,
      processingTime: processingTime,
      timestamp: new Date(),
    });
  } catch (error) {
    console.error("Error calling Go server:", error.message);
    res.status(500).json({
      error: "Failed to process calculation",
    });
  }
});

app.listen(port, () => {
  console.log(`Node.js server running on port ${port}`);
});
