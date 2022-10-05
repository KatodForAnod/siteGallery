CREATE TABLE IF NOT EXISTS Users
(
    Email    varchar(40) PRIMARY KEY,
    Id       SERIAL UNIQUE,
    Name     varchar(40),
    Password varchar(80)
);

CREATE TABLE IF NOT EXISTS Img
(
    Id      SERIAL PRIMARY KEY,
    User_Id INTEGER REFERENCES Users (id),
    Img     text,
    Tags    varchar(40)[]
);

CREATE INDEX idx_scores ON Img USING GIN (Tags);
SET enable_seqscan TO off;

