CREATE DATABASE IF NOT EXISTS tkbai;

USE tkbai;

DROP TABLE IF EXISTS tkbai_data;
DROP TABLE IF EXISTS tkbai_user;

CREATE TABLE
    IF NOT EXISTS tkbai_data (
                                 id INTEGER PRIMARY KEY AUTO_INCREMENT,
                                 student_address VARCHAR(100),
                                 name VARCHAR(50),
                                 student_number VARCHAR(50),
                                 major VARCHAR(30),
                                 insert_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE
    IF NOT EXISTS tkbai_user (
                                 id INTEGER PRIMARY KEY AUTO_INCREMENT,
                                 email VARCHAR(25),
                                 password VARCHAR(50),
                                 insert_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO
    tkbai_user (email, password)
VALUES
    (
        'admin1@mail.com',
        'pWncrRkBIuehkDbdd0Ck4jhUHCao8j/N8BUXnD2c5wo='
    );