CREATE TABLE hotels (
id uuid DEFAULT uuid_generate_v1(),
name VARCHAR(255) NOT NULL,
class VARCHAR(1),
capacity INT,
rooms_left INT,
floors INT,
max_price VARCHAR(10) NOT NULL,
city_name VARCHAR(255) NOT NULL,
address VARCHAR(255) NOT NULL,
PRIMARY KEY (id));

INSERT INTO hotels
 (id,  name, class, capacity, rooms_left, floors, max_price, city_name,address)
  VALUES ( uuid_generate_v1(),'Hotel Ukraine', '3', '1000','218', '12','3200uah', 'Kyiv','Vulytsya Instytutsʹka 4');

INSERT INTO hotels
(id,  name, class, capacity, rooms_left, floors, max_price, city_name,address)
VALUES ( uuid_generate_v1(),'Lviv', '4', '1450','200', '9','3480uah', 'Lviv','Prospect V. Chornovil, 7');

INSERT INTO hotels
(id,  name, class, capacity, rooms_left, floors, max_price, city_name,address)
VALUES ( uuid_generate_v1(),'Citadel Inn', '5', '1234','0', '9','4000uah', 'Lviv','Hrabovskoho Street, 11');

INSERT INTO hotels
(id,  name, class, capacity, rooms_left, floors, max_price, city_name,address)
VALUES ( uuid_generate_v1(),'Nota bene','3','750', '49','4','1380uah', 'Lviv','Valer''yana Polishchuka St, 78');

INSERT INTO hotels
(id,  name, class, capacity, rooms_left, floors, max_price, city_name,address)
VALUES ( uuid_generate_v1(),'Astoria Hotel', '4', '900','390', '6','4000uah', 'Lviv','Hrabovskoho Street, 11');



CREATE TABLE trips_hotels(
  id uuid DEFAULT uuid_generate_v1(),
  trip_id uuid REFERENCES trips (trip_id) ON DELETE CASCADE,
  hotels_id uuid REFERENCES hotels (id) ON DELETE CASCADE,
PRIMARY KEY (id));
