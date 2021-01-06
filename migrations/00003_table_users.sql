-- +goose Up

create table if not exists ad.users (
    id int primary key,
    external_id varchar(256) not null
);

-- +goose Down

drop table if exists ad.users;