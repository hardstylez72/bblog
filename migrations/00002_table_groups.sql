-- +goose Up

create table if not exists ad.groups (
     id  varchar(24) primary key,
     description text not null,
     created_at timestamp default now() not null,
     updated_at timestamp default null,
     deleted_at timestamp default null
);

INSERT INTO ad.groups (id, description) VALUES ('ADMIN', 'allowed to use all methods') on conflict do nothing;
INSERT INTO ad.groups (id, description) VALUES ('GUEST', 'allowed to use some GET methods') on conflict do nothing;
INSERT INTO ad.groups (id, description) VALUES ('USER', 'allowed to use some methods') on conflict do nothing;

-- +goose Down

drop table if exists ad.groups;