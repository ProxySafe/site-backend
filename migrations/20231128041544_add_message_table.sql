-- +goose Up
CREATE TABLE message(
    id       SERIAL NOT NULL,
    account_id INT NOT NULL,
    send_time TIMESTAMP NOT NULL,
    text VARCHAR(500) NOT NULL,
    chat_id INT NOT NULL,
    message_index INT NOT NULL,
    PRIMARY KEY(id)
);


-- +goose Down
DROP TABLE message;