CREATE DATABASE "go-bankmate" WITH OWNER = postgres ENCODING = 'UTF8' LC_COLLATE = 'English_Indonesia.1252' LC_CTYPE = 'English_Indonesia.1252' TABLESPACE = pg_default CONNECTION
LIMIT
    = -1 IS_TEMPLATE = False;

CREATE TABLE m_customer (
    id_customer SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL
);

CREATE TABLE m_merchant (
    id_merchant SERIAL PRIMARY KEY,
    merchant_name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE t_wallet (
    id_wallet SERIAL PRIMARY KEY,
    id_customer INT NOT NULL UNIQUE REFERENCES m_customer(id_customer),
    wallet_amount NUMERIC(18, 2) NOT NULL
);

CREATE TABLE t_token (
    id_token SERIAL PRIMARY KEY,
    id_customer INT NOT NULL REFERENCES m_customer(id_customer),
    token VARCHAR(255) NOT NULL,
    revoked BOOLEAN DEFAULT false
);

CREATE TABLE t_deposit (
    id_deposit SERIAL PRIMARY KEY,
    id_customer INT NOT NULL REFERENCES m_customer(id_customer), 
    deposit_amount NUMERIC(18, 2) NOT NULL,
    deposit_description VARCHAR(255) NOT NULL,
    date_time TIMESTAMP WITHOUT TIME ZONE DEFAULT 'now()'
);

CREATE TABLE t_payment (
    id_payment SERIAL PRIMARY KEY,
    id_customer INT NOT NULL REFERENCES m_customer(id_customer),
    payment_code VARCHAR(255) NOT NULL,
    id_merchant INT NOT NULL REFERENCES m_merchant(id_merchant),
    payment_amount NUMERIC(18, 2) NOT NULL,
    payment_description VARCHAR(255) NOT NULL,
    date_time TIMESTAMP WITHOUT TIME ZONE DEFAULT 'now()'
);

CREATE TABLE t_log(
    id_log SERIAL PRIMARY KEY,
    id_customer INT NOT NULL REFERENCES m_customer(id_customer),
    activity VARCHAR(255),
    date_time TIMESTAMP WITHOUT TIME ZONE DEFAULT 'now()'
);

INSERT INTO m_merchant (merchant_name) VALUES ('Tokopaedi');