-- +goose Up

create table if not exists ad.users (
    id bigserial primary key,
    external_id varchar(256) not null unique,

    created_at timestamp default now() not null,
    updated_at timestamp default null,
    deleted_at timestamp default null
);

-- +goose Down

drop table if exists ad.users;