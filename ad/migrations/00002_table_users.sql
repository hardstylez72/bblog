-- +goose Up

create table if not exists ad.users (
    id bigserial primary key,
    external_id varchar(256) not null,
    is_system bool not null,

    name varchar(256) null,
    description text null,
    email varchar(256) null,
    phone varchar(256) null,

    created_at timestamp default now() not null,
    updated_at timestamp default null,
    deleted_at timestamp default null
);

-- +goose Down

drop table if exists ad.users;