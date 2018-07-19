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
  VALUES ( uuid_generate_v1(), 'Крива липа', 'Lviv', 4, 3, 'Кулінарна студія «Крива Липа» – це авторська кухня без ГМО. Справжні кулінарні шедеври тільки найкращої якості та зі свіжих продуктів від знаних майстрів своєї справи');
INSERT INTO restaurants
  (id,  name, location, stars, prices, description)
  VALUES ( uuid_generate_v1(), 'Криівка', 'Lviv', 5, 5, 'Автентичний заклад, оздоблений у вигляді польової криївки УПА, знаходиться у підвалі одного з будинків');
INSERT INTO restaurants
  (id,  name, location, stars, prices, description)
  VALUES ( uuid_generate_v1(), 'Живий хліб', 'Lviv', 5, 3, 'Хліб та булочки тут готують на натуральних заквасках з італійського борошна. Для круасанів використовують французьке масло.');
INSERT INTO restaurants
  (id,  name, location, stars, prices, description)
  VALUES ( uuid_generate_v1(), 'Фан-бар Банка', 'Lviv', 5, 4, 'Концептуальний демократичний бар, де вперше в Україні всі страви та напої подаються виключно у традиційних скляних банках. Все повинно бути в банках – в барі заборонені пляшки, тарілки, чарки, склянки й інший подібний посуд.');


-- TRIP_RESTAURANTS

CREATE TABLE trips_restaurants(
  id uuid DEFAULT uuid_generate_v1(),
  trip_id uuid REFERENCES trips (trip_id) ON DELETE CASCADE,
  restaurant_id uuid REFERENCES restaurants (id) ON DELETE CASCADE,
  PRIMARY KEY (id)
);
