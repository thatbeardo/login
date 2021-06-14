CREATE TABLE user_types (
    id              INT PRIMARY KEY,
    description     TEXT NOT NULL
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
CREATE TABLE consumers (
    fanfit_user_id      INT     NOT NULL,
    temp_field          TEXT       NULL,
    PRIMARY KEY (fanfit_user_id),
    FOREIGN KEY (fanfit_user_id) REFERENCES users(id)
);


INSERT INTO user_types(id, disc) VALUES(0, 'Creator');
INSERT INTO user_types(id, disc) VALUES(1, 'Consumer');

-- INSERT INTO users(user_type_id, first_name, last_name, email, username, phone_no, gender, profile_picture, bio)
-- VALUES(0, 'Scott', 'Mathison', 'scott@gmail.com', 'scottMath', '792-985-1998',
-- 'Male', 'test_profile_picture', 'Love working with bodyweight to build strenght and power!');
-- INSERT INTO creators(fanfit_user_id, payment_info, logo_picture, background_picture) 
-- VALUES(1, 'rwbrtb', 'test_image', 'test_image');


-- Inserts a new creator
-- DO $$ 
-- DECLARE
--     last_id INTEGER;
-- BEGIN
-- 	INSERT INTO users(user_type_id, first_name, last_name, email, username, phone_no, gender, profile_picture, bio)
-- 	VALUES(0, 'Scott', 'Mathison', 'scotdvt@gmail.com', 'scotdfvtMath', '791-985-1998',
-- 		   'Male', 'test_profile_picture', 'Love working with bodyweight to build strenght and power!')
--       RETURNING users.id INTO last_id;
	  
-- 	INSERT INTO creators(fanfit_user_id, payment_info, logo_picture, background_picture) 
-- 	VALUES(last_id, 'rwbrtb', 'test_image', 'test_image');
-- END $$;


-- Inserts a new consumer
-- DO $$
-- DECLARE
--     last_id INTEGER;
-- BEGIN
-- 	INSERT INTO users(user_type_id, first_name, last_name, email)
-- 	VALUES(1, 'Gabe', 'Dalessandro', 'gabe@gmail.com')
--       RETURNING users.id INTO last_id;
	  
-- 	INSERT INTO consumers(fanfit_user_id) 
-- 	VALUES(last_id);
-- END $$;