CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE trips_flights (
flight_id uuid REFERENCES flights(id) DELETE  ON  CASCADE,
trip_id uuid REFERENCES trips(id) DELETE ON CASCADE,
id uuid DEFAULT uuid_generate_v1(),
PRIMARY KEY (id)
);

CREATE TABLE  flights (
id id uuid DEFAULT uuid_generate_v1(),
departure_city VARCHAR (30) NOT NULL,
departure_time TIME,
departure_date DATE,
arrival_city VARCHAR (30) NOT NULL,
arrival_time TIME,
arrival_date DATE,
price INT NOT NULL,
PRIMARY KEY (id)
);