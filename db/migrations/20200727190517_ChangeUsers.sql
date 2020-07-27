-- +goose Up
-- SQL in this section is executed when the migration is applied.

ALTER TABLE users DROP COLUMN age;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

ALTER TABLE users ADD age int AFTER password;
