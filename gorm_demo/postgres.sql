CREATE DATABASE recordings;

CREATE SCHEMA IF NOT EXISTS recordings;

DROP TABLE IF EXISTS recordings.album;

CREATE TABLE recordings.album (
    id SERIAL  PRIMARY KEY,
    title VARCHAR(128) NOT NULL ,
    artist VARCHAR(255) NOT NULL ,
    price DECIMAL(5,2) NOT NULL
);

INSERT INTO recordings.album
(title, artist, price)
VALUES
    ('Blue Train', 'John Coltrane', 56.99),
    ('Giant Steps', 'John Coltrane', 63.99),
    ('Jeru', 'Gerry Mulligan', 17.99),
    ('Sarah Vaughan', 'Sarah Vaughan', 34.98);