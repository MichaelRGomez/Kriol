-- Filename :migrations/000006_add_schools_indexes.down.sql
drop index if exists schools_name_idx;
drop index if exists schools_level_idx;
drop index if exists schools_mode_idx;
