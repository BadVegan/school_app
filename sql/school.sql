# CREATE DATABASE IF NOT EXISTS school;

DROP DATABASE school;
-- Create a new UTF-8 `schoolboard` database.
CREATE DATABASE school CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE school;

DROP TABLE IF EXISTS teachers;
CREATE TABLE teachers
(
    id      INTEGER      NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name    VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    email   VARCHAR(100),
    phone   VARCHAR(20)
) ENGINE = INNODB;

DROP TABLE IF EXISTS class;
CREATE TABLE class
(
    id         INTEGER      NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name       VARCHAR(100) NOT NULL,
    teacher_id int default null,
    FOREIGN KEY (teacher_id) REFERENCES teachers (id)
) ENGINE = INNODB;

DROP TABLE IF EXISTS students;
CREATE TABLE students
(
    id       INTEGER      NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name     VARCHAR(100) NOT NULL,
    surname  VARCHAR(100) NOT NULL,
    email    VARCHAR(100),
    phone    VARCHAR(20),
    class_id int default null,
    FOREIGN KEY (class_id) REFERENCES class (id)
) ENGINE = INNODB;

DROP TABLE IF EXISTS summary_lesson;
CREATE TABLE summary_lesson
(
    id       INTEGER  NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created  DATETIME NOT NULL,
    summary  JSON     NOT NULL,
    class_id int default null,
    FOREIGN KEY (class_id) REFERENCES class (id)
) ENGINE = INNODB;

INSERT INTO summary_lesson(created, summary)
VALUES (UTC_TIMESTAMP(),
        '{
          "name": "Safari",
          "os": "Mac",
          "resolution": {
            "x": 1920,
            "y": 1080
          }
        }');

CREATE TABLE presence
(
    id      INTEGER  NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created DATETIME NOT NULL,
    present BOOLEAN
) ENGINE = INNODB;

# DROP TABLE IF EXISTS class_summary_lesson;
# CREATE TABLE class_summary_lesson
# (
#     class_id  INTEGER NOT NULL,
#     lesson_id INTEGER NOT NULL
# ) ENGINE = INNODB;


INSERT INTO teachers(name, surname, email, phone)
VALUES ('Nauczycielka', 'Dorota', 'dorota@test.pl', '111-111-111');

INSERT INTO teachers(name, surname, email, phone)
VALUES ('Nauczycielka', 'Kazia', 'kazia@test.pl', '111-111-222');

INSERT INTO teachers(name, surname, email, phone)
VALUES ('Nauczyciel', 'Borys', 'borys@test.pl', '111-111-333');



INSERT INTO class(name, teacher_id)
VALUES ('grupa 1', 1);
INSERT INTO class(name, teacher_id)
VALUES ('grupa 2', 2);
INSERT INTO class(name, teacher_id)
VALUES ('grupa 3', 3);

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