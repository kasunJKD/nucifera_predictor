CREATE SCHEMA IF NOT EXISTS batch1;

CREATE TABLE IF NOT EXISTS batch1.original (
    Date integer NOT NULL,
    Average_Price real,
    Rainfall_Kurunegala real,
    Rainfall_Puttalam real,
    Rainfall_Colombo real,
    Temp_Kurunegala real,
    Temp_Puttalam real,
    Temp_Colombo real,
    PRIMARY KEY (Date)
);

CREATE TABLE IF NOT EXISTS batch1.models (
    Model_Id integer NOT NULL,
    Model_Name varchar(50),
    Plot_Fit bytea,
    Plot_Validation bytea,
    no_features integer,
    feature_list TEXT [],
    PRIMARY KEY (Model_Id)
);

CREATE TABLE IF NOT EXISTS batch1.predictions (
    Model_Id integer NOT NULL,
    Date integer,
    Price real,
    CONSTRAINT fk_model_Id
        FOREIGN KEY(Model_Id)
        REFERENCES batch1.models(Model_Id)
);





