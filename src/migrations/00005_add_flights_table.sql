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
  VALUES (uuid_generate_v1(),'Lviv', '2018-06-16 10:11:26', 'Kyiv', '2018-07-19 10:12:22', 300);
INSERT INTO flights(id,  departure_city, departure, arrival_city, arrival, price)
  VALUES (uuid_generate_v1(),'Sokal', '2018-05-12 11:11:16', 'Lviv', '2018-05-13 09:10:16', 150);
INSERT INTO flights(id,  departure_city, departure, arrival_city, arrival, price)
  VALUES (uuid_generate_v1(),'Kovel', '2018-06-01 09:08:10', 'Germany', '2018-06-02 08:11:12', 700);
INSERT INTO flights(id,  departure_city, departure, arrival_city, arrival, price)
  VALUES (uuid_generate_v1(),'Tokyo', '2018-06-03 06:02:18', 'Kyiv', '2018-06-04 08:10:10', 1200);

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

