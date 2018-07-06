-- TRAINS

CREATE TABLE trains (
   id uuid DEFAULT uuid_generate_v1(),
   departure_time TIME,
   departure_date DATE,
   arrival_time TIME,
   arrival_date DATE,
   departure_city VARCHAR(30) NOT NULL,
   arrival_city VARCHAR(30) NOT NULL,
   train_type VARCHAR(30) NOT NULL,
   car_type VARCHAR(30) NOT NULL,
   price VARCHAR(10) NOT NULL,
   PRIMARY KEY (id)
);

INSERT INTO trains
  (id,  departure_time, departure_date, arrival_time, arrival_date, departure_city, arrival_city, train_type, car_type, price)
  VALUES ( uuid_generate_v1(), '10:23:54', '2004-10-19', '10:23:54', '2004-10-19', 'Lviv', 'Kyiv', 'electric', 'coupe', '200uah');
INSERT INTO trains
  (id,  departure_time, departure_date, arrival_time, arrival_date, departure_city, arrival_city, train_type, car_type, price)
  VALUES ( uuid_generate_v1(), '10:23:54', '2004-10-19', '10:23:54', '2004-10-19', 'Lviv', 'Kyiv', 'electric', 'coupe', '200uah');
INSERT INTO trains
  (id,  departure_time, departure_date, arrival_time, arrival_date, departure_city, arrival_city, train_type, car_type, price)
  VALUES ( uuid_generate_v1(), '10:23:54', '2004-10-19', '10:23:54', '2004-10-19', 'Lviv', 'Kyiv', 'electric', 'coupe', '200uah');
INSERT INTO trains
  (id,  departure_time, departure_date, arrival_time, arrival_date, departure_city, arrival_city, train_type, car_type, price)
  VALUES ( uuid_generate_v1(), '10:23:54', '2004-10-19', '10:23:54', '2004-10-19', 'Lviv', 'Kyiv', 'electric', 'coupe', '200uah');


-- TRIP_TRAINS

CREATE TABLE trips_trains(
  id uuid DEFAULT uuid_generate_v1(),
  trips_id uuid REFERENCES trips (trips_id) ON DELETE CASCADE,
  trains_id uuid REFERENCES trains (id) ON DELETE CASCADE,
  PRIMARY KEY (id)
);
