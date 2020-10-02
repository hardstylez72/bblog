-- +goose Up

create table if not exists ad.users (
    id int primary key,
    auth_type varchar(256) not null,
    external_id varchar(256) not null,
    
    email varchar(256) null,
    login varchar(256) null,
    name varchar(256) null,

    created_at timestamp default now() not null,
    updated_at timestamp default null,
    deleted_at timestamp default null
);

-- +goose Down

drop table if exists ad.users;