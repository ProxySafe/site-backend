create table country_proxy(
    id       int auto_increment
        primary key,
    proxy_id int not null,
    country varchar(50) not null
);