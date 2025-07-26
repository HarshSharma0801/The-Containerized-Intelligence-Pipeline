-- This command must be run separately in psql.
-- It will show an error if the database already exists, which is acceptable in a script.
CREATE DATABASE logs;

-- Connect to the new database. This is a psql meta-command, not SQL.
\c logs

-- Create the process_logs table with corrected data types
CREATE TABLE IF NOT EXISTS process_logs (
    id SERIAL PRIMARY KEY,
    process_number INTEGER NOT NULL,
    time TIMESTAMP
    WITH
        TIME ZONE NOT NULL DEFAULT NOW(),
        -- Storing duration as an INTERVAL is best for calculations
        processing_time INTERVAL NOT NULL
);