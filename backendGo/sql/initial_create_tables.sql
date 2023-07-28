CREATE EXTENSION "pgcrypto";

--Creation of user table
CREATE TABLE IF NOT EXISTS users (
    localId UUID NOT NULL,
    idToken text,
    emailVerified BOOLEAN,
    createdAt TIMESTAMP,
    updatedAt TIMESTAMP,
    passwordHash varchar (150) DEFAULT ''::character varying,
    PRIMARY KEY (localId)
);

--userInfo table
CREATE TABLE IF NOT EXISTS userInfo (
    localId UUID NOT NULL,
    displayName varchar(100),
    firstName varchar(100),
    lastName varchar(100),
    photoUrl varchar(500),
    email varchar(150),
     CONSTRAINT fk_localId
        FOREIGN KEY(localId)
        REFERENCES users(localId)
);
