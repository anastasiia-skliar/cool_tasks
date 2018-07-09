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


-- TRIP_TRAINS

CREATE TABLE trips_trains(
  id uuid DEFAULT uuid_generate_v1(),
  trips_id uuid REFERENCES trips (trips_id) ON DELETE CASCADE,
  trains_id uuid REFERENCES trains (id) ON DELETE CASCADE,
  PRIMARY KEY (id)
);

CREATE TABLE trips(
  trips_id uuid DEFAULT uuid_generate_v1(),
  users_id uuid REFERENCES users (id) ON DELETE CASCADE,
  PRIMARY KEY (trips_id)
);
