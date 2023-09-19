CREATE SCHEMA IF NOT EXISTS batch1;

CREATE TABLE IF NOT EXISTS batch1.original (
    Date integer NOT NULL,
    Average_Price real,
    Rainfall_Kurunegala real,
    Rainfall_Puttalam real,
    Rainfall_Colombo real,
    Min_Temp_Kurunegala real,
    Min_Temp_Puttalam real,
    Min_Temp_Colombo real,
    Max_Temp_Kurunegala real,
    Max_Temp_Puttalam real,
    Max_Temp_Colombo real,
    PRIMARY KEY (Date)
);

CREATE TABLE IF NOT EXISTS batch1.models (
    Model_Id integer NOT NULL,
    Model_Name varchar(50),
    Plot_Fit bytea,
    Plot_Validation bytea,
    Actual_Precited_Graph bytea,
    no_features integer,
    feature_list TEXT [],
    mse real,
    mape real,
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





