-- +migrate Up
CREATE TABLE urls(
    id SERIAL PRIMARY KEY,
    long_url Text NOT NULL,
    short_code VARCHAR(11) UNIQUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
   
);

-- +migrate Down
DROP TABLE urls;