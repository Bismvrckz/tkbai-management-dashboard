CREATE TABLE IF NOT EXISTS
    tkbai_user
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT ,
    email       VARCHAR(25),
    password    VARCHAR(50),
    insert_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO tkbai_user (email, password) VALUES ('admin1@mail.com','Asd123!');