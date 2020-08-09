-- +goose Up

create table if not exists bb.articles (
     id uuid primary key,
     body text not null,
     title text not null,
     preface text not null,
     user_id uuid not null references bb.users(id),
     created_at timestamp not null default now(),
     updated_at timestamp null default now(),
     deleted_at timestamp default null
);

-- +goose Down

drop table if exists bb.articles;