-- +goose Up
CREATE TABLE tcp_proxy(
    id       SERIAL NOT NULL,
    proxy_id INT NOT NULL,
    fingerprint VARCHAR(50) NOT NULL,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE tcp_proxy;