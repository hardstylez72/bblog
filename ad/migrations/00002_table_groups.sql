-- +goose Up

create table if not exists ad.groups (
     id serial primary key,
     code text not null unique,
     description text not null,
     created_at timestamp default now() not null,
     updated_at timestamp default null,
     deleted_at timestamp default null
);

-- +goose Down

drop table if exists ad.groups;