-- +goose Up
-- +goose StatementBegin
CREATE TABLE account(
    id       SERIAL NOT NULL,
    name VARCHAR(50) NOT NULL,
    hashed_password VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    telephone VARCHAR(50) NULL,
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    enabled BOOLEAN NOT NULL,
    chat_id INT, 
    PRIMARY KEY (id)
);
ALTER TABLE account ADD CONSTRAINT account_uniq_name UNIQUE (name);
ALTER TABLE account ADD CONSTRAINT account_uniq_email UNIQUE (email);
ALTER TABLE account ADD CONSTRAINT account_uniq_telephone UNIQUE (telephone);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE account DROP CONSTRAINT account_uniq_name;
ALTER TABLE account DROP CONSTRAINT account_uniq_email;
ALTER TABLE account DROP CONSTRAINT account_uniq_telephone;
DROP TABLE account;
-- +goose StatementEnd