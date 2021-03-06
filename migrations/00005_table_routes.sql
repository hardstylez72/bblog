-- +goose Up

create table if not exists ad.routes (
   id int primary key,
   path varchar(1024) not null,
   method varchar(10) not null,
   entity varchar(128) not null,
   description text not null,
   updated_at timestamp default null,
   created_at timestamp not null default now(),
   deleted_at timestamp default null
);

-- +goose Down

drop table if exists ad.routes;