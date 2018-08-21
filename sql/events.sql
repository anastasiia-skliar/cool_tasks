CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE events (
id uuid DEFAULT uuid_generate_v1(),
title VARCHAR (64) NOT NULL,
category VARCHAR (64) NOT NULL,
town VARCHAR (64) NOT NULL,
date DATE ,
price INT ,
PRIMARY KEY (id)
);

CREATE TABLE trips_events (
id uuid DEFAULT uuid_generate_v1(),
event_id uuid REFERENCES events(id) ON  DELETE  CASCADE,
trip_id uuid REFERENCES trips(trip_id) ON  DELETE  CASCADE,
PRIMARY KEY (id)
);

INSERT INTO events (id, title, category, town, date, price)
 VALUES (uuid_generate_v1(),'Good Traditions of Galicia Fair','fair','Pustomyty District, Viniava village','2018-08-19','0');
 INSERT INTO events (id, title, category, town, date, price)
 VALUES (uuid_generate_v1(),'ZaxidFest','festival','Horodok district, Rodatychi village','2018-08-24','700');
INSERT INTO events (id, title, category, town, date, price)
 VALUES (uuid_generate_v1(),' IT Arena','conference','Lviv','2018-09-28','3405');
INSERT INTO events (id, title, category, town, date, price)
 VALUES (uuid_generate_v1(),'Kacheli','entertaiment','Kyiv','2018-07-16','150');
INSERT INTO events (id, title, category, town, date, price)
 VALUES (uuid_generate_v1(),'Jazz on the beach','concert','Kyiv','2018-08-09','350');
INSERT INTO events (id, title, category, town, date, price)
 VALUES (uuid_generate_v1(),'Hey, you, hello','theatre','Kyiv','2018-07-22','100');
INSERT INTO events (id, title, category, town, date, price)
 VALUES (uuid_generate_v1(),'Zedd','concert','Kyiv','2018-07-19','649');
