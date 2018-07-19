CREATE DATABASE cool_tasks;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
id uuid DEFAULT uuid_generate_v1(),
name VARCHAR(34) NOT NULL,
login VARCHAR(34) NOT NULL,
password VARCHAR(16) NOT NULL,
PRIMARY KEY (id)
);
CREATE TABLE tasks (
id uuid DEFAULT uuid_generate_v1(),
user_id uuid REFERENCES users (id) ON DELETE CASCADE,
name VARCHAR(34) NOT NULL,
time TIMESTAMP,
created_at TIMESTAMP,
updated_at TIMESTAMP,
description TEXT,
PRIMARY KEY(id)
);
