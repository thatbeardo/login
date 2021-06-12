CREATE TABLE user_types (
    id      INT PRIMARY KEY,
    disc    TEXT NOT NULL
);
CREATE TABLE users (
    id              INT     PRIMARY KEY,
    user_type_id    INT     NOT NULL,
    first_name      TEXT    NOT NULL,
    last_name       TEXT    NOT NULL,
    email           TEXT    NOT NULL UNIQUE,
    created_date    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    username        TEXT    NULL UNIQUE,
    phone_no        INT     NULL UNIQUE,
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
    PRIMARY KEY (fanfit_user_id),
    FOREIGN KEY (fanfit_user_id) REFERENCES users(id)
);