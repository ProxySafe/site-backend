-- +goose Up
CREATE TABLE protocol_proxy(
    id       SERIAL NOT NULL,
    proxy_id INT NOT NULL,
    protocol VARCHAR(50) NOT NULL,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE protocol_proxy;