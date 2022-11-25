-- Filename :migrations/000009_add_permissions.up.sql

create table if not exists permissions(
    id bigserial primary key,
    code text not null
);

--create a linking table that links users to permissions
-- this is an example of a many to many relationship
create table if not exists users_permissions(
    user_id bigint not null references users (id) on delete cascade,
    permission_id bigint not null references permissions (id) on delete cascade,
    primary key (user_id, permission_id)
);

insert into permissions (code) 
values ('schools:read'), ('schools:write');