-- +goose Up

create table if not exists ad.routes_groups (
   route_id int  references ad.routes (id) not null,
   group_id int references ad.groups (id) not null,
   UNIQUE  (route_id, group_id)
);

-- +goose Down

drop table if exists ad.routes_groups;