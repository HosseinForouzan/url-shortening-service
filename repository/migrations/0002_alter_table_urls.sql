-- +migrate Up
ALTER TABLE urls
ADD visits integer DEFAULT 0;

-- +migrate Down
ALTER TABLE urls
DROP COLUMN visits;