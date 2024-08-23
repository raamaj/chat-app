create table users
(
    id         bigint not null auto_increment,
    username   varchar(20) not null,
    email       varchar(100) not null,
    password   varchar(50) not null,
    created_at timestamp       not null,
    primary key (id)
) engine = InnoDB;