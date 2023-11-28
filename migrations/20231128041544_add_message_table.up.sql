create table message(
    id       int auto_increment
        primary key,
    account_id int not null,
    send_time datetime not null,
    text varchar(500) not null,
    chat_id int not null,
    message_index int not null
);