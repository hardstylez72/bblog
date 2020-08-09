-- +goose Up

create schema if not exists bb;

-- +goose Down

DROP schema if exists bb;