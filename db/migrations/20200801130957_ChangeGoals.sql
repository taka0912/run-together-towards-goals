-- +goose Up
-- SQL in this section is executed when the migration is applied.

ALTER TABLE goals ADD display_flag tinyint(2) default 0 AFTER goal_name;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

ALTER TABLE goals DROP COLUMN display_flag;
