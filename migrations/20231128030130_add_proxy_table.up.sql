create table proxy(
    id       int auto_increment
        primary key,
    addr varchar(200) not null,
    enabled tinyint not null,
    tcp_fingerprint varchar(100),
    country varchar(50),
    server_name varchar(50) not null,
    is_busy tinyint not null,
    order_id int,
    speed int,
    external_ip varchar(50) not null,
    user varchar(50) not null,
    password varchar(50) not null,
    rotation_period tinyint,
    rented_at datetime,
    rent_finish datetime,
    protocol varchar(50) not null,
    price int not null,
    constraint addr_unique_index
        unique (addr),
    constraint server_unique_index
        unique (server_name)
);