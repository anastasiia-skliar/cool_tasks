CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE trips_museums (
museum_id uuid REFERENCES museums(id) DELETE  ON  CASCADE,
trip_id uuid REFERENCES trips(id) DELETE ON CASCADE,
id uuid DEFAULT uuid_generate_v1(),
PRIMARY KEY (id)
);


CREATE TABLE  museums (
id id uuid DEFAULT uuid_generate_v1(),
name VARCHAR(34) NOT NULL,
location VARCHAR (20) NOT NULL,
price INT NOT NULL,
opened_at INT NOT NULL,
closed_at INT NOT NULL,
museum_type VARCHAR(34) NOT NULL,
additional_info VARCHAR(60) NOT NULL
)