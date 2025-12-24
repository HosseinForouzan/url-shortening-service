-- +migrate Up
CREATE TABLE short_url(
    id SERIAL PRIMARY KEY,
    long_url Text NOT NULL,
    short_code VARCHAR(11),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
   
);

-- +migrate Down
DROP TABLE short_url;