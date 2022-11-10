-- Filename :migrations/000005_add_schools_check_constraint.up.sql
alter table schools add constraint mode_length_check check (array_length(mode, 1) between 1 and 5);