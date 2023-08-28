use dockermysql;
drop table if exists users;
drop table if exists segments;
drop table if exists segments_users;
create table users
(
    id int(15) not null auto_increment,
    primary key (id)
);
create table segments
(
    id   int(15)      not null auto_increment,
    slug varchar(255) not null,
    primary key (id)
);
create table segments_users
(
    user_id    int(15) not null,
    segment_id int(15) not null,
    foreign key (user_id) references users (id),
    foreign key (segment_id) references segments (id),
    primary key (user_id, segment_id)
)