DROP INDEX IF EXISTS idx_tables_restaurant_id cascade;
DROP INDEX IF EXISTS idx_reservations_profile_id cascade;

DROP TABLE IF EXISTS "reservations" cascade;
DROP TABLE IF EXISTS "profiles" cascade;
DROP TABLE IF EXISTS "tables" cascade;
DROP TABLE IF EXISTS "restaurants" cascade;

CREATE TABLE "restaurants"
(
    id serial primary key,
    google_id text default '',
    name text default '',
    description text default '',
    address text default '',
    img_url text default '',
    phone_number text default '',
    email text default '',
    website_url text default '',
    geoposition text default '',
    kitchen text default '',
    tags text[],
    rating int default 1,
    starts_at_cell_id int default 0,
    ends_at_cell_id int default 47
);

CREATE TABLE "tables"
(
    id serial primary key,
    restaurant_id int references "restaurants" (id) on delete cascade not null,
    floor int not null default 1,
    pos_x int default 0,
    pos_y int default 0,
    places int default 0
);

CREATE TABLE "profiles"
(
    id serial primary key,
    firebase_id text not null default '' unique,
    name text not null default '',
    surname text not null default '',
    date_of_birth text not null default '',
    sex text not null default '',
    phone_number text not null default '',
    email text not null default '',
    password text not null default '',
    avatar_url text not null default ''
);

CREATE TABLE "reservations"
(
    id serial primary key,
    table_id int references "tables" (id) on delete cascade not null,
    profile_id text references "profiles" (firebase_id) on delete cascade not null,
    reservation_date text default '',
    cells int[],
    comment text default '',
    num_of_guests int default 0
);

--TODO: Недавно просмотренные, Избранное (rest_id, cli_id), Рейтинг (rest_id, cli_id, value)

CREATE INDEX IF NOT EXISTS idx_tables_restaurant_id on tables using btree(restaurant_id);
--CREATE INDEX IF NOT EXISTS idx_reservations_profile_id on reservations using btree(profile_id);