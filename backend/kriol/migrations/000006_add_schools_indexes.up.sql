-- Filename :migrations/000006_add_schools_indexes.up.sql
create index if not exists schools_name_idx on schools using gin(to_tsvector('simple', name));
create index if not exists schools_level_idx on schools using gin(to_tsvector('simple', level));
create index if not exists schools_mode_idx on schools using gin(mode);