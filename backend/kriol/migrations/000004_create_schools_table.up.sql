-- Filename :migrations/000004_create_schools_table.up.sql

create table if not exists schools(
    id bigserial PRIMARY KEY,
    createed_at timestamp (0) with time zone not null default now(),
    name text not null,
    level text not null,
    contact text not null,
    phone text not null,
    email text not null,
    website text not null,
    address text not null,
    mode text[] not null,
    version int not null default 1
);