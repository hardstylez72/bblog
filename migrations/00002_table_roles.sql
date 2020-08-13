-- +goose Up

create table if not exists bb.roles (
     id int primary key,
     code varchar(24) not null,
     description text null
);

INSERT INTO bb.roles (id, code, description) VALUES (0, 'ADMIN', null) on conflict do nothing;
INSERT INTO bb.roles (id, code, description) VALUES (1, 'GUEST', null) on conflict do nothing;
INSERT INTO bb.roles (id, code, description) VALUES (2, 'USER', null) on conflict do nothing;

-- +goose Down

drop table if exists bb.roles;