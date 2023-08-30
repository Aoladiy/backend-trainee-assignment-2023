create database if not exists dockermysql;
use dockermysql;
drop table if exists segments_users_log;
drop table if exists segments_users;
drop table if exists segments;
drop table if exists users;
create table users
(
    id int(15) not null auto_increment,
    primary key (id)
);
create table segments
(
    id   int(15)      not null auto_increment,
    slug varchar(255) not null unique,
    primary key (id)
);
create table segments_users
(
    user_id         int(15) not null,
    segment_id      int(15) not null,
    expiration_time DATETIME,
    foreign key (user_id) references users (id) on delete cascade,
    foreign key (segment_id) references segments (id) on delete cascade,
    primary key (user_id, segment_id)
);
create table segments_users_log
(
    id         int(15)      not null auto_increment primary key,
    user_id    int(15)      not null,
    segment_id int(15)      not null,
    action     varchar(255) not null,
    datetime   DATETIME     not null default now(),
    foreign key (user_id) references users (id) on delete cascade,
    foreign key (segment_id) references segments (id) on delete cascade
);
