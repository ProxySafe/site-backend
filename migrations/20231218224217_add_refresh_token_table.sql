-- +goose Up
-- +goose StatementBegin
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
ALTER TABLE refresh_token ADD CONSTRAINT refresh_token_uniq_account_id UNIQUE (account_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE account DROP CONSTRAINT refresh_token_uniq_account_id;
DROP TABLE refresh_token;
-- +goose StatementEnd