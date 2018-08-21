CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE  museums (
id uuid DEFAULT uuid_generate_v1(),
name VARCHAR(34) NOT NULL,
location VARCHAR (20) NOT NULL,
price INT,
opened_at TIME NOT NULL,
closed_at TIME NOT NULL,
museum_type VARCHAR(34) NOT NULL,
additional_info VARCHAR(60) NOT NULL,
PRIMARY KEY (id)
);

CREATE TABLE trips_museums (
museum_id uuid REFERENCES museums(id) ON DELETE CASCADE,
trip_id uuid REFERENCES trips(trip_id) ON DELETE CASCADE,
id uuid DEFAULT uuid_generate_v1(),
PRIMARY KEY (id)
);

INSERT INTO museums (name, location, price, opened_at, closed_at, museum_type, additional_info) VALUES
('Arsenal museum', 'Lviv', 40, '10:00:00', '20:00:00', 'History', 'Half-price for students');
INSERT INTO museums (name, location, price, opened_at, closed_at, museum_type, additional_info) VALUES
('Art gallery', 'Lviv', 50, '09:00:00', '18:00:00', 'Art', 'Closed on Monday');
INSERT INTO museums (name, location, price, opened_at, closed_at, museum_type, additional_info) VALUES
('Chocolate museum', 'Lviv', 0, '10:30:00', '20:00:00', 'Specialty', 'You can buy chocolate');
INSERT INTO museums (name, location, price, opened_at, closed_at, museum_type, additional_info) VALUES
('Aviation museum', 'Kiev', 100, '12:00:00', '18:00:00', 'Specialty', 'Open air');
INSERT INTO museums (name, location, price, opened_at, closed_at, museum_type, additional_info) VALUES
('St. Michael''s Cathedrale', 'Kiev', 0, '08:00:00', '21:00:00', 'Sacred & Religious Sites', 'Can climb the bell tower');
INSERT INTO museums (name, location, price, opened_at, closed_at, museum_type, additional_info) VALUES
('Golden gate', 'Kiev', 60, '10:00:00', '20:00:00', 'History', 'Located in the center of the city');
