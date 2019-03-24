# CREATE DATABASE IF NOT EXISTS school;

DROP DATABASE school;
-- Create a new UTF-8 `schoolboard` database.
CREATE DATABASE school CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE school;

-- Create a `class` table.
CREATE TABLE class
(
  id   INTEGER      NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL
)ENGINE=INNODB;

-- Create a `students` table.
CREATE TABLE students
(
  id      INTEGER      NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name    VARCHAR(100) NOT NULL,
  surname VARCHAR(100) NOT NULL,
  email   VARCHAR(100),
  phone   VARCHAR(20),
  class_id int default null,
  FOREIGN KEY (class_id) REFERENCES class (id)
)ENGINE=INNODB;

INSERT INTO class(name) VALUES ('grupa 1');
INSERT INTO class(name) VALUES ('grupa 2');
INSERT INTO class(name) VALUES ('grupa 3');

INSERT INTO students(name, surname, email, phone, class_id)
VALUES ('Jan', 'Szyszka', 'szyszka@test.pl', '111-111-111', 1);

INSERT INTO students(name, surname, email, phone, class_id)
VALUES ('Czesio', 'Ogórek', 'czesio@test.pl', '222-222-222', 1);

INSERT INTO students(name, surname, email, phone, class_id)
VALUES ('Anna', 'Zegarek', 'anna@test.pl', '333-333-333', 2);

INSERT INTO students(name, surname, email, phone, class_id)
VALUES ('Tomasz', 'Pomidor', 'tomasz@test.pl', '444-444-444', 2);

INSERT INTO students(name, surname, email, phone, class_id)
VALUES ('Wacław', 'Goły', 'waclaw@test.pl', '555-555-555', 3);

INSERT INTO students(name, surname, email, phone, class_id)
VALUES ('Kazimierz', 'Sztywny', 'kazik@test.pl', '666-666-666', 3);