-- +goose Up

create table if not exists ad.routes (
   id serial primary key,
   route varchar(1024) not null,
   method varchar(10) not null,
   description text not null,
   updated_at timestamp default null,
   created_at timestamp not null default now(),
   deleted_at timestamp default null
);

create unique index routes_id_index
    on ad.routes (method, route);

-- +goose Down

drop table if exists ad.routes;

