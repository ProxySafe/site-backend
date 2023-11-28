create table order(
    id       int auto_increment
        primary key,
    order_date datetime not null,
    account_id int not null,
    order_expiration_date datetime not null,
    proxies_number int not null  
);