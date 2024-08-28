create table articles
(
    id bigint not null auto_increment,
    title varchar(255) not null,
    image_url varchar(255),
    content text,
    author varchar(50),
    like_count int,
    comment_count int,
    created_at timestamp not null,
    primary key (id)
);