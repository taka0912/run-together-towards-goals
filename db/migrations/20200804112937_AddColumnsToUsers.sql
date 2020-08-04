-- +goose Up
-- SQL in this section is executed when the migration is applied.

ALTER TABLE `users`
    ADD COLUMN age              varchar(60),
    ADD COLUMN age_display_flag tinyint(2) default 0,
    ADD COLUMN address          varchar(60),
    ADD COLUMN birth_place      varchar(60),
    ADD COLUMN hobby            varchar(255),
    ADD COLUMN occupation       varchar(255),
    ADD COLUMN strong_point     varchar(255),
    ADD COLUMN skill            varchar(255);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

ALTER TABLE users
    DROP COLUMN age,
    DROP COLUMN age_display_flag,
    DROP COLUMN address,
    DROP COLUMN birth_place,
    DROP COLUMN hobby,
    DROP COLUMN occupation,
    DROP COLUMN strong_point,
    DROP COLUMN skill;
