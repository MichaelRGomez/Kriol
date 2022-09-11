--filename: kriol/backend/kriol/migrations/kriol_up.sql
create table if not exists intial_entries(
    kriol text not null,
    english text not null
);