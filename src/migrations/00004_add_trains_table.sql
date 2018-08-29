-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE trains (
   id uuid DEFAULT uuid_generate_v1(),
   departure TIMESTAMP ,
   arrival TIMESTAMP ,
   departure_city VARCHAR(30) NOT NULL,
   arrival_city VARCHAR(30) NOT NULL,
   train_type VARCHAR(30) NOT NULL,
   car_type VARCHAR(30) NOT NULL,
   price INT NOT NULL,
   PRIMARY KEY (id)
);

INSERT INTO trains
  (id,  departure, arrival, departure_city, arrival_city, train_type, car_type, price)
  VALUES ( uuid_generate_v1(), '2018-08-30 12:20:00', '2018-08-31 13:55:00', 'Lviv', 'Odessa', 'electric', 'coupe', 200);
INSERT INTO trains
  (id,  departure, arrival, departure_city, arrival_city, train_type, car_type, price)
  VALUES ( uuid_generate_v1(), '2018-08-30 13:20:00', '2018-08-30 16:55:00', 'Lviv', 'Kyiv', 'electric', 'coupe', 250);
INSERT INTO trains
  (id,  departure, arrival, departure_city, arrival_city, train_type, car_type, price)
  VALUES ( uuid_generate_v1(), '2018-08-31 12:50:00', '2018-09-01 13:55:00', 'Lviv', 'Moscow', 'electric', 'coupe', 320);
INSERT INTO trains
  (id,  departure, arrival, departure_city, arrival_city, train_type, car_type, price)
  VALUES ( uuid_generate_v1(), '2018-08-30 13:20:00', '2018-08-30 14:55:00', 'Lviv', 'Odessa', 'electric', 'coupe', 190);
INSERT INTO trains
  (id,  departure, arrival, departure_city, arrival_city, train_type, car_type, price)
  VALUES ( uuid_generate_v1(), '2018-08-31 17:20:00', '2018-08-31 20:55:00', 'Lviv', 'Odessa', 'electric', 'coupe', 230);
-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE trips_trains(
  id uuid DEFAULT uuid_generate_v1(),
  trip_id uuid REFERENCES trips (trip_id) ON DELETE CASCADE ON UPDATE CASCADE,
  train_id uuid REFERENCES trains (id) ON DELETE CASCADE ON UPDATE CASCADE,
  PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE trips_trains;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE trains;

