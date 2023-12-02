-- +goose Up
-- +goose StatementBegin
CREATE TABLE proxy(
    id       SERIAL NOT NULL,
    addr VARCHAR(200) NOT NULL,
    enabled SMALLINT NOT NULL,
    tcp_fingerprint VARCHAR(100),
    country VARCHAR(50),
    server_name VARCHAR(50) NOT NULL,
    is_busy BOOLEAN NOT NULL,
    order_id INT,
    speed INT,
    external_ip VARCHAR(50) NOT NULL,
    user_name VARCHAR(50) NOT NULL,
    proxy_password VARCHAR(50) NOT NULL,
    rotation_period SMALLINT,
    rented_at TIMESTAMP,
    rent_finish TIMESTAMP,
    protocol VARCHAR(50) NOT NULL,
    price INT NOT NULL,
    PRIMARY KEY(id)
);
ALTER TABLE proxy ADD CONSTRAINT proxy_uniq_name UNIQUE (addr);
ALTER TABLE proxy ADD CONSTRAINT proxy_uniq_server_name UNIQUE (server_name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE proxy DROP CONSTRAINT proxy_uniq_name;
ALTER TABLE proxy DROP CONSTRAINT proxy_uniq_server_name;
DROP TABLE proxy;
-- +goose StatementEnd