CREATE TABLE users (
    id              INT     PRIMARY KEY,
    first_name      TEXT    NOT NULL,
    last_name       TEXT    NOT NULL,
    email           TEXT    NOT NULL UNIQUE,
    created_date    TIMESTAMP NOT NULL,
    username        TEXT NOT NULL UNIQUE,
    phone_no        INT     NOT NULL UNIQUE,
    gender          TEXT    NOT NULL,
    profile_picture VARCHAR(1000) NOT NULL,
    bio             VARCHAR(1000) NOT NULL
   
);