CREATE TABLE users (
    first_name      TEXT    NOT NULL,
    last_name       TEXT    NOT NULL,
    email           TEXT    NOT NULL UNIQUE
);

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