-- +goose Up

create table if not exists bb.users (
    id uuid primary key,
    registered_at timestamp default now(),
    external_id varchar(256) null,
    external_auth_type varchar(256),
    login varchar(256),
    first_name varchar(256) ,
    last_name varchar(256),
    middle_name varchar(256),
    is_banned bool
);

-- +goose Down

drop table if exists bb.users;