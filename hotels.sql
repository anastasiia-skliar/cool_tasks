CREATE TABLE hotels (
id uuid DEFAULT uuid_generate_v1(),
name VARCHAR(255) NOT NULL,
class string,
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


CREATE TABLE trips_hotels(
  id uuid DEFAULT uuid_generate_v1(),
  trips_id uuid REFERENCES trips (trips_id) ON DELETE CASCADE,
  hotels_id uuid REFERENCES hotels (id) ON DELETE CASCADE,
PRIMARY KEY (id));
