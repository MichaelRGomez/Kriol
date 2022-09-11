--filename: kriol/backend/kriol/migrations/kriol_down.sql
alter table intial_entries
drop column entry_date;