DROP INDEX IF EXISTS idx_tables_restaurant_id cascade;
DROP INDEX IF EXISTS idx_reservations_profile_id cascade;

DROP TABLE IF EXISTS "reservations" cascade;
DROP TABLE IF EXISTS "profiles" cascade;
DROP TABLE IF EXISTS "tables" cascade;
DROP TABLE IF EXISTS "restaurants" cascade;

CREATE TABLE "restaurants"
(
    id serial primary key,
    google_id text,
    address text,
    description text,
    tags text[],
    img_url text,
    phone_number text,
    email text,
    website_url text,
    geoposition text,
    rating int
);

CREATE TABLE "tables"
(
    id serial primary key,
    restaurant_id int references "restaurants" (id) on delete cascade not null,
    floor int,
    pos_x int,
    pos_y int,
    places int
);

CREATE TABLE "profiles"
(
    id serial primary key,
    firebase_id text,
    name text,
    surname text,
    date_of_birth date,
    sex text,
    phone_number text,
    email text,
    password text,
    avatar_url text
);

CREATE TABLE "reservations"
(
    id serial primary key,
    table_id int references "tables" (id) on delete cascade not null,
    profile_id int references "profiles" (id) on delete cascade not null,
    reservation_date date,
    cell_id int,
    num_of_cells int,
    comment text
);

--TODO: Недавно просмотренные, Избранное (rest_id, cli_id), Рейтинг (rest_id, cli_id, value)

CREATE INDEX IF NOT EXISTS idx_tables_restaurant_id on tables using btree(restaurant_id);
CREATE INDEX IF NOT EXISTS idx_reservations_profile_id on reservations using btree(profile_id);