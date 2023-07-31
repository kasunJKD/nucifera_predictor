CREATE EXTENSION "pgcrypto";

--Creation of user table
CREATE TABLE IF NOT EXISTS users (
    userId UUID NOT NULL,
    email varchar(150),
    emailVerified BOOLEAN,
    createdAt TIMESTAMP,
    updatedAt TIMESTAMP,
    passwordHash varchar(150) DEFAULT ''::character varying,
    PRIMARY KEY (userId)
);

--userInfo table
CREATE TABLE IF NOT EXISTS userInfo (
    userId UUID NOT NULL,
    displayName varchar(100),
    firstName varchar(100),
    lastName varchar(100),
    gender varchar(150),
    address varchar(150),
     CONSTRAINT fk_userId
        FOREIGN KEY(userId)
        REFERENCES users(userId)
);