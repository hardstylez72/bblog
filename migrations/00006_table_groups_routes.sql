-- +goose Up

create table if not exists ad.groups_routes (
   id int references ad.routes (id),
   group_id varchar(24) references ad.groups (id),
   created_at timestamp not null default now(),
   deleted_at timestamp default null
);

-- +goose Down

drop table if exists ad.groups_routes;