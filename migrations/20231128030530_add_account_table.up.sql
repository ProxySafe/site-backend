create table account(
    id       int auto_increment
        primary key,
    name varchar(50) not null,
    hashed_password varchar(100) not null,
    email varchar(100) not null,
    telephone varchar(50) not null,
    created_at datetime not null,
    deleted_at datetime,
    enabled tinyint not null,
    chat_id int, 
    constraint name_unique_index
        unique (name),
    constraint email_unique_index
        unique (email),
    constraint telephone_unique_index
        unique (telephone)
);