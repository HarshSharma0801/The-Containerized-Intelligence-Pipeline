-- Connect to the logs database
\connect logs

-- Create process_logs table inside logs database
CREATE TABLE IF NOT EXISTS process_logs (
    id SERIAL PRIMARY KEY,
    process_number SERIAL NOT NULL,
    time TIMESTAMP NOT NULL,
    processing_time VARCHAR NOT NULL
);
