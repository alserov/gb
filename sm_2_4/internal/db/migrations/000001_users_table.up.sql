CREATE TABLE IF NOT EXISTS users(
    id text primary key unique ,
    username text unique ,
    password text ,
    age smallint

);

CREATE INDEX username_idx ON users(username);