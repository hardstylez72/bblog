-- +goose Up

create table if not exists ad.users (
    id bigserial primary key,
    external_id varchar(256) null,
    is_business bool,

    login varchar(128) null,

    last_name varchar(256) null,
    first_name varchar(256) null,
    middle_name varchar(256) null,

    description text null,

    email varchar(256) null,
    phone varchar(256) null
);

-- +goose Down

drop table if exists ad.users;