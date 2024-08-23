create table conversations
(
    id         bigint not null auto_increment,
    created_at timestamp       not null,
    primary key (id)
) engine = InnoDB;