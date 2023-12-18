-- +goose Up
CREATE TABLE refresh_token (
    id SERIAL NOT NULL,
    account_id INT NOT NULL,
    token VARCHAR(50) NOT NULL,
    expires TIMESTAMP NOT NULL,
    user_agent VARCHAR(50) NOT NULL,
    fingerprint INT NOT NULL,
    os VARCHAR(20) NOT NULL,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE refresh_token;