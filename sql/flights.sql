CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE trips_flights (
flight_id uuid REFERENCES flights(id) ON DELETE CASCADE,
trip_id uuid REFERENCES trips(trip_id) ON DELETE CASCADE,
id uuid DEFAULT uuid_generate_v1(),
PRIMARY KEY (id)
);

CREATE TABLE  flights (
id uuid DEFAULT uuid_generate_v1(),
departure_city VARCHAR (30) NOT NULL,
departure_time TIME,
departure_date DATE,
arrival_city VARCHAR (30) NOT NULL,
arrival_time TIME,
arrival_date DATE,
price INT NOT NULL,
PRIMARY KEY (id)
);

INSERT INTO flights
  (id,  departure_city, departure_time, departure_date, arrival_city, arrival_time, arrival_date, price)
  VALUES (uuid_generate_v1(),'Lviv', '10:11:26', '2018-06-16', 'Kyiv','10:12:22', '2018-07-19', '300');
INSERT INTO flights
  (id,  departure_city, departure_time, departure_date, arrival_city, arrival_time, arrival_date, price)
  VALUES (uuid_generate_v1(),'Sokal', '11:11:16', '2018-05-12', 'Lviv','09:10:16', '2018-05-13', '150');
INSERT INTO flights
  (id,  departure_city, departure_time, departure_date, arrival_city, arrival_time, arrival_date, price)
  VALUES (uuid_generate_v1(),'Kovel', '09:08:10', '2018-06-01', 'Germany','08:11:12', '2018-06-02', '700');
INSERT INTO flights
  (id,  departure_city, departure_time, departure_date, arrival_city, arrival_time, arrival_date, price)
  VALUES (uuid_generate_v1(),'Tokyo', '06:02:18', '2018-06-03', 'Kyiv','08:10:10', '2018-06-04', '1200');
