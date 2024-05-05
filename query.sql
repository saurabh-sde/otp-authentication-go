CREATE DATABASE auth;

CREATE TABLE IF NOT EXISTS app_user ( 
    id SERIAL NOT NULL, name varchar(128) DEFAULT NULL, mobile varchar(60) DEFAULT NULL, status varchar(60) DEFAULT NULL, created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS log_event ( 
    id SERIAL NOT NULL, event_name VARCHAR(45) DEFAULT NULL, user_id varchar(30) NOT NULL, created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (id)
);