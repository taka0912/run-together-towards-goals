SET CHARSET UTF8;

DROP DATABASE IF EXISTS my_goal;

CREATE DATABASE IF NOT EXISTS my_goal DEFAULT CHARACTER SET utf8;

create table users (
    id integer primary key AUTO_INCREMENT,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    nickname varchar(255),
    password varchar(255),
    age integer,
    role integer
);

create index idx_users_deleted_at on users (deleted_at);

create table daily_kpts (
    id integer primary key AUTO_INCREMENT,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    user_id integer not null,
    keep varchar(255) not null,
    problem varchar(255) not null,
    try varchar(255) not null,
    good integer,
    fight integer
);

create index idx_daily_kpts_deleted_at on daily_kpts (deleted_at);

create table goals (
    id integer primary key AUTO_INCREMENT,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    user_id integer not null,
    genre_id integer not null,
    goal_name varchar(255) not null
);

create index idx_goals_deleted_at on goals (deleted_at);

create table kpt_reaction_histories (
    id integer primary key AUTO_INCREMENT,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    kpt_id integer not null,
    user_id integer not null,
    reaction integer not null
);

create index idx_kpt_reaction_histories_deleted_at on kpt_reaction_histories (deleted_at);

create table todo_lists (
    id integer primary key AUTO_INCREMENT,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    goal_id integer not null,
    required_elements varchar(255) not null,
    specific_goal varchar(255) not null,
    todo varchar(255) not null,
    limit_date date not null
);

create index idx_todo_lists_deleted_at on todo_lists (deleted_at);

create table genres (
    id integer primary key AUTO_INCREMENT,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    genre_name varchar(255) not null
);

create index idx_genres_deleted_at on genres (deleted_at);

INSERT INTO genres (id, created_at, updated_at, deleted_at, genre_name) VALUES (1, now(), now(), null, 'ダイエット');
INSERT INTO genres (id, created_at, updated_at, deleted_at, genre_name) VALUES (2, now(), now(), null, '筋トレ');
INSERT INTO genres (id, created_at, updated_at, deleted_at, genre_name) VALUES (3, now(), now(), null, '健康');
INSERT INTO genres (id, created_at, updated_at, deleted_at, genre_name) VALUES (4, now(), now(), null, 'プログラミング');
INSERT INTO genres (id, created_at, updated_at, deleted_at, genre_name) VALUES (5, now(), now(), null, '資格勉強');
INSERT INTO genres (id, created_at, updated_at, deleted_at, genre_name) VALUES (6, now(), now(), null, '副業');
INSERT INTO genres (id, created_at, updated_at, deleted_at, genre_name) VALUES (7, now(), now(), null, '語学');
INSERT INTO genres (id, created_at, updated_at, deleted_at, genre_name) VALUES (8, now(), now(), null, 'その他');
