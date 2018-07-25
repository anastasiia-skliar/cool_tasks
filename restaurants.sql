-- RESTAURANTS

CREATE TABLE restaurants (
   id uuid DEFAULT uuid_generate_v1(),
   name VARCHAR(100) NOT NULL,
   location VARCHAR(100) NOT NULL,
   stars int,
   prices int,
   description TEXT,
   PRIMARY KEY (id)
);

INSERT INTO restaurants
  (id,  name, location, stars, prices, description)
  VALUES ( uuid_generate_v1(), 'Kryva Lypa', 'Lviv', 4, 3, 'Some info 1');
INSERT INTO restaurants
  (id,  name, location, stars, prices, description)
  VALUES ( uuid_generate_v1(), 'Kryivka', 'Lviv', 5, 5, 'Some info 2');
INSERT INTO restaurants
  (id,  name, location, stars, prices, description)
  VALUES ( uuid_generate_v1(), 'Jiviy hlib', 'Lviv', 5, 3, 'Some info 3');
INSERT INTO restaurants
  (id,  name, location, stars, prices, description)
  VALUES ( uuid_generate_v1(), 'Fun-bar Banka', 'Lviv', 5, 4, 'Some info 4');


-- TRIP_RESTAURANTS

CREATE TABLE trips_restaurants(
  id uuid DEFAULT uuid_generate_v1(),
  trip_id uuid REFERENCES trips (trip_id) ON DELETE CASCADE,
  restaurant_id uuid REFERENCES restaurants (id) ON DELETE CASCADE,
  PRIMARY KEY (id)
);
