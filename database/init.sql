\c logs

-- Create the process_logs table with corrected data types
CREATE TABLE IF NOT EXISTS process_logs (
    id SERIAL PRIMARY KEY,
    time TIMESTAMP
    WITH
        TIME ZONE NOT NULL DEFAULT NOW(),
        -- Storing duration as an INTERVAL is best for calculations
        processing_time INTERVAL NOT NULL
);