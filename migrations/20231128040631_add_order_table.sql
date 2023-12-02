-- +goose Up
CREATE TABLE proxy_order(
    id       SERIAL NOT NULL,
    order_date TIMESTAMP NOT NULL,
    account_id INT NOT NULL,
    order_expiration_date TIMESTAMP NOT NULL,
    proxies_number INT NOT NULL,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE order;