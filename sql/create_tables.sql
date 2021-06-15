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





INSERT INTO user_types(id, description) VALUES(0, 'Creator');
INSERT INTO user_types(id, description) VALUES(1, 'Consumer');

INSERT INTO users(user_type_id, first_name, last_name, email, username, phone_no, gender, profile_picture, bio)
VALUES(0, 'Scott', 'Mathison', 'scott@gmail.com', 'scottMath', '792-985-1998',
'Male', 'test_profile_picture', 'Love working with bodyweight to build strenght and power!');
INSERT INTO creators(fanfit_user_id, payment_info, logo_picture, background_picture) 
VALUES(1, 'rwbrtb', 'test_image', 'test_image');

-- MAKING THE FUNCTION TO CALL
-- DO InsertCreator 
-- DECLARE
--     last_id INTEGER;
-- BEGIN
-- 	INSERT INTO users(user_type_id, first_name, last_name, email, username, phone_no, gender, profile_picture, bio)
-- 	VALUES(0, 'Scott', 'Mathison', 'scotdvt@gmail.com', 'scotdfvtMath', '791-985-1998',
-- 		   'Male', 'test_profile_picture', 'Love working with bodyweight to build strenght and power!')
--     RETURNING users.id INTO last_id;
	  
-- 	INSERT INTO creators(fanfit_user_id, payment_info, logo_picture, background_picture) 
-- 	VALUES(last_id, 'rwbrtb', 'test_image', 'test_image');
-- END InsertCreator; 


-- Inserts a new creator
-- DO $$ 
-- DECLARE
--     last_id INTEGER;
-- BEGIN
-- 	INSERT INTO users(user_type_id, first_name, last_name, email, username, phone_no, gender, profile_picture, bio)
-- 	VALUES(0, 'Scott', 'Mathison', 'scotdvt@gmail.com', 'scotdfvtMath', '791-985-1998',
-- 		   'Male', 'test_profile_picture', 'Love working with bodyweight to build strenght and power!')
--     RETURNING users.id INTO last_id;
	  
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


-- CREATE OR REPLACE FUNCTION P_SHOPIFY_FULFILLMENT_LOAD (
--     p_id                                    text,
--     p_invoice_id                            text,
--     p_tracking_number                       text)
--  RETURNS void
--     LANGUAGE plpgsql
--     AS $$
-- DECLARE
--     l_invoice_id                            numeric;
-- 	l_sales_channel_shipment_id             text;
-- BEGIN
-- 	BEGIN
-- 	    SELECT
-- 	        invoice_id
-- 	    INTO
-- 	        l_invoice_id
-- 	    FROM
-- 	        invoice
-- 	    WHERE
-- 	        sales_channel_order_id = p_invoice_id AND
-- 	        sales_channel_code = 'S';
-- 	EXCEPTION
-- 	    When NO_DATA_FOUND THEN l_invoice_id := NULL;
-- 	END;
--     IF l_invoice_id IS NULL THEN
--         NULL;
--     ELSE
	
--         BEGIN
--             SELECT
--                 sales_channel_shipment_id
--             INTO
--                 l_sales_channel_shipment_id
--             FROM
--                 shipment
--             WHERE
--                 invoice_id = l_invoice_id AND
--                 sales_channel_shipment_id = p_id ;
--         EXCEPTION
--             WHEN NO_DATA_FO