
create table if not exists bb.phones (
    id uuid primary key,
    user_id uuid references bb.users(id),
    created_at timestamp default now(),
    is_active bool,
    number varchar(256)
)