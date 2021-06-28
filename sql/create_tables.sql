CREATE TABLE user_types (
    id      INT PRIMARY KEY,
    description    TEXT NOT NULL
);
CREATE TABLE users (
    id              SERIAL     PRIMARY KEY,
    user_type_id    INT     NOT NULL,
    first_name      TEXT    NOT NULL,
    last_name       TEXT    NOT NULL,
    email           TEXT    NOT NULL UNIQUE,
    created_date    TIMESTAMP NOT NULL DEFAULT Now(),
    username        TEXT    NULL UNIQUE,
    phone_no        TEXT    NULL UNIQUE,
    gender          TEXT    NULL,
    profile_picture VARCHAR(1000) NULL,
    bio             VARCHAR(1000) NULL,
    FOREIGN KEY (user_type_id) REFERENCES user_types(id)
);
CREATE TABLE creators (
    fanfit_user_id      INT     NOT NULL,
    payment_info        TEXT    NOT NULL,
    logo_picture        TEXT    NOT NULL,
    background_picture  TEXT    NOT NULL,
    PRIMARY KEY (fanfit_user_id),
    FOREIGN KEY (fanfit_user_id) REFERENCES users(id)
);
CREATE TABLE clients (
    fanfit_user_id      INT     NOT NULL,
    temp_field          TEXT        NULL,
    PRIMARY KEY (fanfit_user_id),
    FOREIGN KEY (fanfit_user_id) REFERENCES users(id)
);


INSERT INTO user_types(id, description) VALUES(0, 'Creator');
INSERT INTO user_types(id, description) VALUES(1, 'Clients');

-- Inserts a new creator
DO $$
DECLARE
    last_id INTEGER;
BEGIN
    INSERT INTO users(user_type_id, first_name, last_name, email, username, phone_no, gender, profile_picture, bio)
    VALUES(0, 'Scott', 'Mathison', 'scott@gmail.com', 'scottMath', '791-985-1998',
           'Male', 'public/scottMath/scott_mathison_profile_picture.png', 'Love working with bodyweight to build strenght and power!')
    RETURNING users.id INTO last_id;
    INSERT INTO creators(fanfit_user_id, payment_info, logo_picture, background_picture)
    VALUES(last_id, 'rwbrtb', 'test_image', 'public/scottMath/scott_mathison_background_image.png');
END $$;
DO $$
DECLARE
    last_id INTEGER;
BEGIN
    INSERT INTO users(user_type_id, first_name, last_name, email, username, phone_no, gender, profile_picture, bio)
    VALUES(0, 'Stefana', 'Hall', 'stefana@gmail.com', 'definingFitness', '443-890-1952',
           'Female', 'test_profile_picture', 'Love female body positivity!')
    RETURNING users.id INTO last_id;
    INSERT INTO creators(fanfit_user_id, payment_info, logo_picture, background_picture)
    VALUES(last_id, 'rwbrtb', 'test_image', 'test_image');
END $$;
DO $$
DECLARE
    last_id INTEGER;
BEGIN
    INSERT INTO users(user_type_id, first_name, last_name, email, username, phone_no, gender, profile_picture, bio)
    VALUES(0, 'Gabe', 'Dalessandro', 'gdalessa@usc.edu', 'italianstallion', '570-332-5722',
           'Male', 'test_profile_picture', 'Sports and living a healthy lifestyle!')
    RETURNING users.id INTO last_id;
    INSERT INTO creators(fanfit_user_id, payment_info, logo_picture, background_picture)
    VALUES(last_id, 'rwbrtb', 'test_image', 'test_image');
END $$;
-- Inserts a new client
DO $$
DECLARE
    last_id INTEGER;
BEGIN
    INSERT INTO users(user_type_id, first_name, last_name, email)
    VALUES(1, 'Gabe', 'Dalessandro', 'gabe@gmail.com')
      RETURNING users.id INTO last_id;
    INSERT INTO clients(fanfit_user_id)
    VALUES(last_id);
END $$;
DO $$
DECLARE
    last_id INTEGER;
BEGIN
    INSERT INTO users(user_type_id, first_name, last_name, email)
    VALUES(1, 'Harshil', 'Mavani', 'harshil@gmail.com')
      RETURNING users.id INTO last_id;
    INSERT INTO clients(fanfit_user_id)
    VALUES(last_id);
END $$;

-- Function to create clients: Finish implementing this and then just call it a bunch
-- CREATE OR REPLACE FUNCTION P_CREATE_NEW_CLIENT (
--  p_user_type_id  INTEGER,
--  p_first_name    TEXT,
--  p_last_name     TEXT,
--  p_email         TEXT)
--  RETURNS void
--  LANGUAGE plpgsql
--     AS
-- $$
-- DECLARE
--  new_id          INTEGER;
-- BEGIN
--  INSERT INTO users(user_type_id, first_name, last_name, email)
--  VALUES(p_user_type_id, p_first_name, p_last_name, p_email)
--  RETURNING users.id INTO new_id;
--  INSERT INTO clients(fanfit_user_id)
--  VALUES(new_id);
-- END;
-- $$