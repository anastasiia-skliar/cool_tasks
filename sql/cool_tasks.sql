--
-- PostgreSQL database dump
--

-- Dumped from database version 10.4 (Debian 10.4-2.pgdg90+1)
-- Dumped by pg_dump version 10.4 (Ubuntu 10.4-0ubuntu0.18.04)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

CREATE DATABASE cool_tasks WITH TEMPLATE = template0 ENCODING = 'SQL_ASCII' LC_COLLATE = 'C' LC_CTYPE = 'C';


ALTER DATABASE cool_tasks OWNER TO postgres;

\connect cool_tasks


-- Dumped from database version 10.4 (Ubuntu 10.4-0ubuntu0.18.04)
-- Dumped by pg_dump version 10.4 (Ubuntu 10.4-0ubuntu0.18.04)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'SQL_ASCII';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- Name: chkpass; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS chkpass WITH SCHEMA public;


--
-- Name: EXTENSION chkpass; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION chkpass IS 'data type for auto-encrypted passwords';


--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: events; Type: TABLE; Schema: public; Owner: postgres
--
CREATE TABLE public.events (
    id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    title character varying(64) NOT NULL,
    category character varying(64) NOT NULL,
    town character varying(64) NOT NULL,
    date date,
    price integer
);


ALTER TABLE public.events OWNER TO postgres;

--
-- Name: flights; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.flights (
    id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    departure_city character varying(30) NOT NULL,
    departure_time time without time zone,
    departure_date date,
    arrival_city character varying(30) NOT NULL,
    arrival_time time without time zone,
    arrival_date date,
    price integer NOT NULL
);


ALTER TABLE public.flights OWNER TO postgres;

--
-- Name: hotels; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.hotels (
    id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    name character varying(255) NOT NULL,
    class character varying(1),
    capacity integer,
    rooms_left integer,
    floors integer,
    max_price character varying(10) NOT NULL,
    city_name character varying(255) NOT NULL,
    address character varying(255) NOT NULL
);


ALTER TABLE public.hotels OWNER TO postgres;

--
-- Name: museums; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.museums (
    id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    name character varying(34) NOT NULL,
    location character varying(20) NOT NULL,
    price integer,
    opened_at time without time zone NOT NULL,
    closed_at time without time zone NOT NULL,
    museum_type character varying(34) NOT NULL,
    additional_info character varying(60) NOT NULL
);


ALTER TABLE public.museums OWNER TO postgres;

--
-- Name: restaurants; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.restaurants (
    id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    name character varying(100) NOT NULL,
    location character varying(100) NOT NULL,
    stars integer,
    prices integer,
    description text
);


ALTER TABLE public.restaurants OWNER TO postgres;

--
-- Name: tasks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tasks (
    id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    user_id uuid,
    name character varying(34) NOT NULL,
    "time" timestamp without time zone,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    description text
);


ALTER TABLE public.tasks OWNER TO postgres;

--
-- Name: trains; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.trains (
    id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    departure_time time without time zone,
    departure_date date,
    arrival_time time without time zone,
    arrival_date date,
    departure_city character varying(30) NOT NULL,
    arrival_city character varying(30) NOT NULL,
    train_type character varying(30) NOT NULL,
    car_type character varying(30) NOT NULL,
    price character varying(10) NOT NULL
);


ALTER TABLE public.trains OWNER TO postgres;

--
-- Name: trips; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.trips (
    trip_id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    user_id uuid
);


ALTER TABLE public.trips OWNER TO postgres;

--
-- Name: trips_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.trips_events (
    id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    event_id uuid,
    trip_id uuid
);


ALTER TABLE public.trips_events OWNER TO postgres;

--
-- Name: trips_flights; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.trips_flights (
    flight_id uuid,
    trip_id uuid,
    id uuid DEFAULT public.uuid_generate_v1() NOT NULL
);


ALTER TABLE public.trips_flights OWNER TO postgres;

--
-- Name: trips_hotels; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.trips_hotels (
    id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    trip_id uuid,
    hotels_id uuid
);


ALTER TABLE public.trips_hotels OWNER TO postgres;

--
-- Name: trips_museums; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.trips_museums (
    museum_id uuid,
    trip_id uuid,
    id uuid DEFAULT public.uuid_generate_v1() NOT NULL
);


ALTER TABLE public.trips_museums OWNER TO postgres;

--
-- Name: trips_restaurants; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.trips_restaurants (
    id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    trip_id uuid,
    restaurant_id uuid
);


ALTER TABLE public.trips_restaurants OWNER TO postgres;

--
-- Name: trips_trains; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.trips_trains (
    id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    trip_id uuid,
    train_id uuid
);


ALTER TABLE public.trips_trains OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    name character varying(34) NOT NULL,
    login character varying(34) NOT NULL,
    password public.chkpass,
    role character varying(16)
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Data for Name: events; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.events (id, title, category, town, date, price) FROM stdin;
9badbade-85cb-11e8-b86f-c01885c5bc39	Good Traditions of Galicia Fair	fair	Pustomyty District, Viniava village	2018-08-19	0
9badbadf-85cb-11e8-b86f-c01885c5bc39	ZaxidFest	festival	Horodok district, Rodatychi village	2018-08-24	700
9badbae0-85cb-11e8-b86f-c01885c5bc39	 IT Arena	conference	Lviv	2018-09-28	3405
9badbae1-85cb-11e8-b86f-c01885c5bc39	Kacheli	entertaiment	Kyiv	2018-07-16	150
9badbae2-85cb-11e8-b86f-c01885c5bc39	Jazz on the beach	concert	Kyiv	2018-08-09	350
9badbae3-85cb-11e8-b86f-c01885c5bc39	Hey, you, hello	theatre	Kyiv	2018-07-22	100
9badbae4-85cb-11e8-b86f-c01885c5bc39	Zedd	concert	Kyiv	2018-07-19	649
\.


--
-- Data for Name: flights; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.flights (id, departure_city, departure_time, departure_date, arrival_city, arrival_time, arrival_date, price) FROM stdin;
7640e5dc-85cb-11e8-b86f-c01885c5bc39	Lviv	10:11:26	2018-06-16	Kyiv	10:12:22	2018-07-19	300
7640e5dd-85cb-11e8-b86f-c01885c5bc39	Sokal	11:11:16	2018-05-12	Lviv	09:10:16	2018-05-13	150
7640e5de-85cb-11e8-b86f-c01885c5bc39	Kovel	09:08:10	2018-06-01	Germany	08:11:12	2018-06-02	700
7640e5df-85cb-11e8-b86f-c01885c5bc39	Tokyo	06:02:18	2018-06-03	Kyiv	08:10:10	2018-06-04	1200
\.


--
-- Data for Name: hotels; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.hotels (id, name, class, capacity, rooms_left, floors, max_price, city_name, address) FROM stdin;
5e25825a-85cb-11e8-b86f-c01885c5bc39	Hotel Ukraine	3	1000	218	12	3200uah	Kyiv	Vulytsya Instytutsʹka 4
5e25825b-85cb-11e8-b86f-c01885c5bc39	Lviv	4	1450	200	9	3480uah	Lviv	Prospect V. Chornovil, 7
5e25825c-85cb-11e8-b86f-c01885c5bc39	Citadel Inn	5	1234	0	9	4000uah	Lviv	Hrabovskoho Street, 11
5e25825d-85cb-11e8-b86f-c01885c5bc39	Nota bene	3	750	49	4	1380uah	Lviv	Valer'yana Polishchuka St, 78
5e25825e-85cb-11e8-b86f-c01885c5bc39	Astoria Hotel	4	900	390	6	4000uah	Lviv	Hrabovskoho Street, 11
\.


--
-- Data for Name: museums; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.museums (id, name, location, price, opened_at, closed_at, museum_type, additional_info) FROM stdin;
45caf12c-85cb-11e8-b86f-c01885c5bc39	Arsenal museum	Lviv	40	10:00:00	20:00:00	History	Half-price for students
45caf12d-85cb-11e8-b86f-c01885c5bc39	Art gallery	Lviv	50	09:00:00	18:00:00	Art	Closed on Monday
45caf12e-85cb-11e8-b86f-c01885c5bc39	Chocolate museum	Lviv	0	10:30:00	20:00:00	Specialty	You can buy chocolate
45caf12f-85cb-11e8-b86f-c01885c5bc39	Aviation museum	Kiev	100	12:00:00	18:00:00	Specialty	Open air
45caf130-85cb-11e8-b86f-c01885c5bc39	St. Michael's Cathedrale	Kiev	0	08:00:00	21:00:00	Sacred & Religious Sites	Can climb the bell tower
45caf131-85cb-11e8-b86f-c01885c5bc39	Golden gate	Kiev	60	10:00:00	20:00:00	History	Located in the center of the city
\.


--
-- Data for Name: restaurants; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.restaurants (id, name, location, stars, prices, description) FROM stdin;
a3b6d6f8-86a6-11e8-9a39-d4bed959082a	Крива липа	Lviv	4	3	Кулінарна студія «Крива Липа» – це авторська кухня без ГМО. Справжні кулінарні шедеври тільки найкращої якості та зі свіжих продуктів від знаних майстрів своєї справи
a3b6d6f9-86a6-11e8-9a39-d4bed959082a	Криівка	Lviv	5	5	Автентичний заклад, оздоблений у вигляді польової криївки УПА, знаходиться у підвалі одного з будинків
a3b6d6fa-86a6-11e8-9a39-d4bed959082a	Живий хліб	Lviv	5	3	Хліб та булочки тут готують на натуральних заквасках з італійського борошна. Для круасанів використовують французьке масло.
a3b6d6fb-86a6-11e8-9a39-d4bed959082a	Фан-бар Банка	Lviv	5	4	Концептуальний демократичний бар, де вперше в Україні всі страви та напої подаються виключно у традиційних скляних банках. Все повинно бути в банках – в барі заборонені пляшки, тарілки, чарки, склянки й інший подібний посуд.
\.


--
-- Data for Name: tasks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tasks (id, user_id, name, "time", created_at, updated_at, description) FROM stdin;
\.


--
-- Data for Name: trains; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trains (id, departure_time, departure_date, arrival_time, arrival_date, departure_city, arrival_city, train_type, car_type, price) FROM stdin;
2ee64f92-85cb-11e8-b86f-c01885c5bc39	11:23:54	2018-07-20	18:23:54	2018-07-22	Lviv	Odessa	electric	coupe	200uah
2ee64f93-85cb-11e8-b86f-c01885c5bc39	10:23:54	2018-07-21	17:23:54	2018-07-24	Kyiv	Moscow	electric	coupe	190uah
2ee64f94-85cb-11e8-b86f-c01885c5bc39	12:23:54	2018-07-22	16:23:54	2018-07-23	Lviv	Kyiv	electric	coupe	225uah
2ee64f95-85cb-11e8-b86f-c01885c5bc39	15:23:54	2018-07-23	20:23:54	2018-07-25	Lviv	Kharkiv	electric	coupe	320uah
\.


--
-- Data for Name: trips; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trips (trip_id, user_id) FROM stdin;
\.


--
-- Data for Name: trips_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trips_events (id, event_id, trip_id) FROM stdin;
\.


--
-- Data for Name: trips_flights; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trips_flights (flight_id, trip_id, id) FROM stdin;
\.


--
-- Data for Name: trips_hotels; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trips_hotels (id, trip_id, hotels_id) FROM stdin;
\.


--
-- Data for Name: trips_museums; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trips_museums (museum_id, trip_id, id) FROM stdin;
\.


--
-- Data for Name: trips_restaurants; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trips_restaurants (id, trip_id, restaurant_id) FROM stdin;
\.


--
-- Data for Name: trips_trains; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trips_trains (id, trip_id, train_id) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, login, password, role) FROM stdin;
03dc3258-86a7-11e8-9a39-d4bed959082a	John	admin	:SrU4bmypbPpMo	admin
\.


--
-- Name: events events_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- Name: flights flights_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.flights
    ADD CONSTRAINT flights_pkey PRIMARY KEY (id);


--
-- Name: hotels hotels_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.hotels
    ADD CONSTRAINT hotels_pkey PRIMARY KEY (id);


--
-- Name: museums museums_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.museums
    ADD CONSTRAINT museums_pkey PRIMARY KEY (id);


--
-- Name: restaurants restaurants_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.restaurants
    ADD CONSTRAINT restaurants_pkey PRIMARY KEY (id);


--
-- Name: tasks tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (id);


--
-- Name: trains trains_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trains
    ADD CONSTRAINT trains_pkey PRIMARY KEY (id);


--
-- Name: trips_events trips_events_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_events
    ADD CONSTRAINT trips_events_pkey PRIMARY KEY (id);


--
-- Name: trips_flights trips_flights_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_flights
    ADD CONSTRAINT trips_flights_pkey PRIMARY KEY (id);


--
-- Name: trips_hotels trips_hotels_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_hotels
    ADD CONSTRAINT trips_hotels_pkey PRIMARY KEY (id);


--
-- Name: trips_museums trips_museums_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_museums
    ADD CONSTRAINT trips_museums_pkey PRIMARY KEY (id);


--
-- Name: trips trips_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips
    ADD CONSTRAINT trips_pkey PRIMARY KEY (trip_id);


--
-- Name: trips_restaurants trips_restaurants_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_restaurants
    ADD CONSTRAINT trips_restaurants_pkey PRIMARY KEY (id);


--
-- Name: trips_trains trips_trains_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_trains
    ADD CONSTRAINT trips_trains_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: tasks tasks_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: trips_events trips_events_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_events
    ADD CONSTRAINT trips_events_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id) ON DELETE CASCADE;


--
-- Name: trips_events trips_events_trip_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_events
    ADD CONSTRAINT trips_events_trip_id_fkey FOREIGN KEY (trip_id) REFERENCES public.trips(trip_id) ON DELETE CASCADE;


--
-- Name: trips_flights trips_flights_flight_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_flights
    ADD CONSTRAINT trips_flights_flight_id_fkey FOREIGN KEY (flight_id) REFERENCES public.flights(id) ON DELETE CASCADE;


--
-- Name: trips_flights trips_flights_trip_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_flights
    ADD CONSTRAINT trips_flights_trip_id_fkey FOREIGN KEY (trip_id) REFERENCES public.trips(trip_id) ON DELETE CASCADE;


--
-- Name: trips_hotels trips_hotels_hotels_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_hotels
    ADD CONSTRAINT trips_hotels_hotels_id_fkey FOREIGN KEY (hotels_id) REFERENCES public.hotels(id) ON DELETE CASCADE;


--
-- Name: trips_hotels trips_hotels_trip_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_hotels
    ADD CONSTRAINT trips_hotels_trip_id_fkey FOREIGN KEY (trip_id) REFERENCES public.trips(trip_id) ON DELETE CASCADE;


--
-- Name: trips_museums trips_museums_museum_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_museums
    ADD CONSTRAINT trips_museums_museum_id_fkey FOREIGN KEY (museum_id) REFERENCES public.museums(id) ON DELETE CASCADE;


--
-- Name: trips_museums trips_museums_trip_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_museums
    ADD CONSTRAINT trips_museums_trip_id_fkey FOREIGN KEY (trip_id) REFERENCES public.trips(trip_id) ON DELETE CASCADE;


--
-- Name: trips_restaurants trips_restaurants_restaurant_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_restaurants
    ADD CONSTRAINT trips_restaurants_restaurant_id_fkey FOREIGN KEY (restaurant_id) REFERENCES public.restaurants(id) ON DELETE CASCADE;


--
-- Name: trips_restaurants trips_restaurants_trip_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_restaurants
    ADD CONSTRAINT trips_restaurants_trip_id_fkey FOREIGN KEY (trip_id) REFERENCES public.trips(trip_id) ON DELETE CASCADE;


--
-- Name: trips_trains trips_trains_train_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_trains
    ADD CONSTRAINT trips_trains_train_id_fkey FOREIGN KEY (train_id) REFERENCES public.trains(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: trips_trains trips_trains_trip_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips_trains
    ADD CONSTRAINT trips_trains_trip_id_fkey FOREIGN KEY (trip_id) REFERENCES public.trips(trip_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: trips trips_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trips
    ADD CONSTRAINT trips_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

