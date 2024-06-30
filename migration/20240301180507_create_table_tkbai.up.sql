CREATE TABLE IF NOT EXISTS
    tkbai_data
(
    id             INT PRIMARY KEY,
    test_id        VARCHAR(100),
    name           VARCHAR(50),
    student_number VARCHAR(50),
    major          VARCHAR(30),
    date_of_test   TIMESTAMP,
    toefl_score    VARCHAR(5),
    insert_date    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);