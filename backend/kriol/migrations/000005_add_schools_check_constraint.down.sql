-- Filename :migrations/000005_add_schools_check_constraint.down.sql
alter table schools drop constraint if exists mode_lenght_check;