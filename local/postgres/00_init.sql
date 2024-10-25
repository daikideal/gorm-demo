-- 
-- migrations
CREATE TABLE IF NOT EXISTS bytea_samples (
    id VARCHAR(255) PRIMARY KEY,
    data BYTEA NOT NULL
);