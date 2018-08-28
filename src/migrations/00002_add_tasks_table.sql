-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE tasks (
id uuid DEFAULT uuid_generate_v1(),
user_id uuid REFERENCES users (id) ON DELETE CASCADE,
name VARCHAR(34) NOT NULL,
time TIMESTAMP,
created_at TIMESTAMP,
updated_at TIMESTAMP,
description TEXT,
completed BOOLEAN DEFAULT false,
PRIMARY KEY(id)
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE tasks;