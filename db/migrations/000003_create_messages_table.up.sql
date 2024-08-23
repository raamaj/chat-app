create table messages
(
    id         bigint not null auto_increment,
    conversation_id   bigint not null,
    sender_id       bigint not null,
    content   text not null,
    send_at timestamp       not null,
    primary key (id)
) engine = InnoDB;