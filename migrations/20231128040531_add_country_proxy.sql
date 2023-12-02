-- +goose Up
CREATE TABLE country_proxy(
    id       SERIAL NOT NULL,
    proxy_id INT NOT NULL,
    country VARCHAR(50) NOT NULL,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE country_proxy;