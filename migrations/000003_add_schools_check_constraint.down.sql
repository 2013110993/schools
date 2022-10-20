-- Filename: migrations/000001_add_schools_check_constraint.down.sql

ALTER TABLE schools ADD CONSTRAInT IF EXISTS mode_length_check;