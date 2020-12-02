-- +goose Up

create table if not exists ad.users_groups (
   user_id int  references ad.users (id),
   group_id varchar(24) references ad.groups (id),
   created_at timestamp not null default now(),
   deleted_at timestamp default null
);

-- +goose Down

drop table if exists ad.users_groups;