-- Active: 1691754991722@@127.0.0.1@5435@proxysafe
-- +goose Up
ALTER TABLE proxy
DROP COLUMN rented_at,
DROP COLUMN rent_finish,
DROP COLUMN is_busy,
ADD COLUMN is_busy SMALLINT;

-- +goose Down
ALTER TABLE proxy
ADD COLUMN rented_at TIMESTAMP,
ADD COLUMN rent_finish TIMESTAMP
DROP COLUMN is_busy,
ADD COLUMN is_busy BOOLEAN;
