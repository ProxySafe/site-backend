create table tcp_proxy(
    id       int auto_increment
        primary key,
    proxy_id int not null,
    fingerprint varchar(50) not null
);