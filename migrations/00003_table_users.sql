-- +goose Up

create table if not exists bb.users (
    id uuid primary key,
    registered_at timestamp default now() not null,
    external_auth_type varchar(256) not null,
    external_id varchar(256) not null,
    
    email varchar(256) null,
    login varchar(256) null,
    name varchar(256) null,

    is_banned bool not null default false,
    role_id int references bb.roles(id)
);

-- +goose Down

drop table if exists bb.users;