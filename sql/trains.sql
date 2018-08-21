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
  VALUES ( uuid_generate_v1(), '11:23:54', '2018-07-20', '18:23:54', '2018-07-22', 'Lviv', 'Odessa', 'electric', 'coupe', '200uah');
INSERT INTO trains
  (id,  departure_time, departure_date, arrival_time, arrival_date, departure_city, arrival_city, train_type, car_type, price)
  VALUES ( uuid_generate_v1(), '10:23:54', '2018-07-21', '17:23:54', '2018-07-24', 'Kyiv', 'Moscow', 'electric', 'coupe', '190uah');
INSERT INTO trains
  (id,  departure_time, departure_date, arrival_time, arrival_date, departure_city, arrival_city, train_type, car_type, price)
  VALUES ( uuid_generate_v1(), '12:23:54', '2018-07-22', '16:23:54', '2018-07-23', 'Lviv', 'Kyiv', 'electric', 'coupe', '225uah');
INSERT INTO trains
  (id,  departure_time, departure_date, arrival_time, arrival_date, departure_city, arrival_city, train_type, car_type, price)
  VALUES ( uuid_generate_v1(), '15:23:54', '2018-07-23', '20:23:54', '2018-07-25', 'Lviv', 'Kharkiv', 'electric', 'coupe', '320uah');

SELECT * FROM trains WHERE (departure_city = 'Lviv' AND arrival_city = 'Kyiv' AND (departure_time >= '10:00:00' AND departure_time <= '17:00:00'));

-- TRIP_TRAINS

CREATE TABLE trips_trains(
  id uuid DEFAULT uuid_generate_v1(),
  trip_id uuid REFERENCES trips (trip_id) ON DELETE CASCADE ON UPDATE CASCADE,
  train_id uuid REFERENCES trains (id) ON DELETE CASCADE ON UPDATE CASCADE,
  PRIMARY KEY (id)
);

SELECT trains.*, museums.* FROM trips LEFT JOIN trips_trains ON trips_trains.trip_id = trips.trip_id LEFT JOIN trains ON trips_trains.train_id = trains.id
LEFT JOIN trips_museums ON trips_museums.trip_id = trips.trip_id LEFT JOIN museums ON trips_museums.museum_id = museums.id WHERE trips.user_id = 'a46889d4-839f-11e8-8d89-c01885c5bc39';

SELECT trains.* FROM trips, (SELECT * FROM trains, trips_trains WHERE trips_trains.train_id = trains.id) AS trains WHERE trains.trip_id = trips.trip_id AND trips.user_id = 'a46889d4-839f-11e8-8d89-c01885c5bc39';
