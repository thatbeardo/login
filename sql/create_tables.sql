CREATE TABLE user_types{
    id      INT PRIMARY KEY,
    disc    TEXT NOT NULL,
}
CREATE TABLE users (
    id              INT     PRIMARY KEY,
    user_type_id    INT     NOT NULL,
    first_name      TEXT    NOT NULL,
    last_name       TEXT    NOT NULL,
    email           TEXT    NOT NULL UNIQUE,
    created_date    TIMESTAMP NOT NULL,
    username        TEXT NOT NULL UNIQUE,
    phone_no        INT     NOT NULL UNIQUE,
    gender          TEXT    NOT NULL,
    profile_picture VARCHAR(1000) NOT NULL,
    bio             VARCHAR(1000) NOT NULL,
    FOREIGN KEY (user_type_id) REFERENCES user_types(id)
);

CREATE TABLE creators{
    fanfit_user_id      INT     NOT NULL,
    payment_info        TEXT    NOT NULL,
    logo_picture        TEXT    NOT NULL,
    background_picture  TEXT    NOT NULL
};
CREATE TABLE consumers{
    fanfit_user_id      INT     FOREIGN KEY
};
    -- create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- id SERIAL PRIMARY  KEY     NOT NULL,
    -- user_type_id    INT     NOT NULL,
    -- first_name      TEXT    NOT NULL,
    -- last_name       TEXT    NOT NULL,
    -- email           TEXT    NOT NULL UNIQUE, 
    -- username        TEXT    NOT NULL UNIQUE, 
    -- phone_no        INT     NOT NULL,
    -- gender          VARCHAR(10)  NOT NULL, 
    -- profile_picture VARCHAR(10)  NOT NULL, 
    -- bio             VARCHAR(10)  NOT NULL   