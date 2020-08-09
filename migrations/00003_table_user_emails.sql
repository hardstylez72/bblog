-- +goose Up

create table if not exists bb.user_emails (
     id uuid primary key,
     user_id uuid references bb.users(id),
     created_at timestamp default now(),
     is_active bool,
     address varchar(256)
);

-- +goose Down

drop table if exists bb.user_emails;