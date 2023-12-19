-- +goose Up
ALTER TABLE refresh_token
ALTER COLUMN token TYPE VARCHAR(200),
ALTER COLUMN user_agent TYPE VARCHAR(200);

-- +goose Down
ALTER TABLE refresh_token
ALTER COLUMN token TYPE VARCHAR(50),
ALTER COLUMN user_agent TYPE VARCHAR(50);