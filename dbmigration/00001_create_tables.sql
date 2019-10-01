-- +goose Up
CREATE TABLE users
(
    id       int NOT NULL PRIMARY KEY,
    username text,
    name     text,
    surname  text
);

INSERT INTO users
VALUES (0, 'test', '', ''),
       (1, 'gil', 'Gilbert', 'Ryumugabe');