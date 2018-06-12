CREATE DATABASE cool_tasks;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE cool_tasks.users ( 
id uuid DEFAULT uuid_generate_v1(),
name VARCHAR(34) NOT NULL,
login VARCHAR(34) NOT NULL,
password VARCHAR(16) NOT NULL,
PRIMARY KEY (id)
);

CREATE TABLE cool_tasks.tasks (
id uuid DEFAULT uuid_generate_v1(),
user_id uuid REFERENCES users (id),
name VARCHAR(34) NOT NULL,
time TIMESTAMP,
created_at TIMESTAMP,
updated_at TIMESTAMP,
description TEXT,
PRIMARY KEY(id)
);

