-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
id uuid DEFAULT uuid_generate_v1(),
name VARCHAR(34) NOT NULL,
login VARCHAR(34) NOT NULL,
password chkpass,
role VARCHAR(16),
PRIMARY KEY (id)
);
INSERT INTO users(id,name,login,password,role) VALUES (uuid_generate_v1(),'John','admin','admin','admin');
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;