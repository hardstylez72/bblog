-- +goose Up

create table if not exists ad.users_routes (
   route_id int references ad.routes (id),
   user_id bigint references ad.users (id),
   is_excluded bool default false not null
);

-- +goose Down

drop table if exists ad.users_routes;