-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE flights(
id uuid DEFAULT uuid_generate_v1(),
departure_city VARCHAR (30) NOT NULL,
departure TIMESTAMP ,
arrival_city VARCHAR (30) NOT NULL,
arrival TIMESTAMP ,
price INT NOT NULL,
PRIMARY KEY (id)
);

INSERT INTO flights(id,  departure_city, departure, arrival_city, arrival, price)
  VALUES (uuid_generate_v1(),'Lviv', '2018-08-31 17:20:00', 'Odessa',  '2018-08-31 18:55:00', 300);
INSERT INTO flights(id,  departure_city, departure, arrival_city, arrival, price)
  VALUES (uuid_generate_v1(),'Odessa', '2018-09-01 11:11:16', 'Lviv', '2018-09-01 09:10:16', 150);
INSERT INTO flights(id,  departure_city, departure, arrival_city, arrival, price)
  VALUES (uuid_generate_v1(),'Lviv', '2018-08-30 09:08:10', 'Moscow', '2018-08-30 15:11:12', 700);
INSERT INTO flights(id,  departure_city, departure, arrival_city, arrival, price)
  VALUES (uuid_generate_v1(),'Tokyo', '2018-08-30 06:02:18', 'Kyiv', '2018-08-30 08:10:10', 1200);

-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE trips_flights (
flight_id uuid REFERENCES flights(id) ON DELETE CASCADE,
trip_id uuid REFERENCES trips(trip_id) ON DELETE CASCADE,
id uuid DEFAULT uuid_generate_v1(),
PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE trips_flights;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE flights;

