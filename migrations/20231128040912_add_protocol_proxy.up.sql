create table protocol_proxy(
    id       int auto_increment
        primary key,
    proxy_id int not null,
    protocol varchar(50) not null
);