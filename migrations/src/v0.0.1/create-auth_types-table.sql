
create table if not exists bb.auth_types (
    id uuid primary key default uuid_in(md5(random()::text || clock_timestamp()::text)::cstring),
    code varchar(256) not null
);

insert into bb.auth_types (id, code) values ('d7c8e05b-156b-4e86-811f-30b26bb81240','google');
insert into bb.auth_types (id, code) values ('b65cb084-095a-4b48-9913-6fcbb1afa21e','yandex');
insert into bb.auth_types (id, code) values ('e97892e1-7507-440a-80fa-e43b74aff469','github');

