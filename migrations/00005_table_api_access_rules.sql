-- +goose Up

create table bb.api_access_rules
(
    id          serial                not null
        constraint user_role_urls_pk
            primary key,
    role_id     int references bb.roles(id),
    url_pattern varchar(256)          not null,
    method      varchar(32)           not null,
    description text,
    is_exclude  boolean default false not null
);

-- +goose Down

drop table if exists bb.api_access_rules;