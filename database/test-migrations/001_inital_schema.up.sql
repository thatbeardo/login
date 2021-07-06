BEGIN;

CREATE TYPE user_type AS ENUM ('client', 'creator');

CREATE TABLE IF NOT EXISTS users (
    id              SERIAL     PRIMARY KEY,
    user_type    user_type     NOT NULL,
    first_name      TEXT    NOT NULL,
    last_name       TEXT    NOT NULL,
    email           TEXT    NOT NULL UNIQUE,
    created_date    TIMESTAMP NOT NULL DEFAULT Now(),
    username        TEXT    NULL UNIQUE,
    phone_no        TEXT    NULL UNIQUE,
    gender          TEXT    NULL,
    profile_picture VARCHAR(1000) NULL,
    bio             VARCHAR(1000) NULL,
    background_picture TEXT         NULL
);

CREATE TABLE IF NOT EXISTS clients (
    fanfit_user_id      INT     NOT NULL,
    temp_field          TEXT       NULL,
    PRIMARY KEY (fanfit_user_id),
    FOREIGN KEY (fanfit_user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS creators (
    fanfit_user_id      INT        NOT NULL,
    payment_info        TEXT       NULL,
    logo_picture        TEXT       NULL,
    PRIMARY KEY (fanfit_user_id),
    FOREIGN KEY (fanfit_user_id) REFERENCES users(id)
);
-- Inserts a new creator
DO $$
DECLARE
    last_id INTEGER;
BEGIN
    INSERT INTO users(user_type, first_name, last_name, email, username, phone_no, gender, profile_picture, bio, background_picture)
    VALUES('creator', 'Scott', 'Mathison', 'scott@gmail.com', 'scottMath', '791-985-1998',
           'Male', 'scottMath/scott_mathison_profile_picture.png', 'Love working with bodyweight to build strenght and power!', 'scottMath/scott_mathison_background_image.png')
    RETURNING users.id INTO last_id;
    INSERT INTO creators(fanfit_user_id, payment_info, logo_picture)
    VALUES(last_id, 'rwbrtb', 'test_image');
END $$;
DO $$
DECLARE
    last_id INTEGER;
BEGIN
    INSERT INTO users(user_type, first_name, last_name, email, username, phone_no, gender, profile_picture, bio, background_picture)
    VALUES('creator', 'Stefana', 'Hall', 'stefana@gmail.com', 'definingFitness', '443-890-1952',
           'Female', 'test_profile_picture', 'Love female body positivity!', 'test_image')
    RETURNING users.id INTO last_id;
    INSERT INTO creators(fanfit_user_id, payment_info, logo_picture)
    VALUES(last_id, 'rwbrtb', 'test_image');
END $$;
DO $$
DECLARE
    last_id INTEGER;
BEGIN
    INSERT INTO users(user_type, first_name, last_name, email, username, phone_no, gender, profile_picture, bio, background_picture)
    VALUES('creator', 'Gabe', 'Dalessandro', 'gdalessa@usc.edu', 'italianstallion', '570-332-5722',
           'Male', 'test_profile_picture', 'Sports and living a healthy lifestyle!', 'test_image')
    RETURNING users.id INTO last_id;
    INSERT INTO creators(fanfit_user_id, payment_info, logo_picture)
    VALUES(last_id, 'rwbrtb', 'test_image');
END $$;



COMMIT;













