from flask import Flask, request, render_template
import csv
import psycopg2
from .model import predictLSTM

app = Flask(__name__)

@app.route('/upload', methods=['POST'])
def upload():
    file = request.files['csv_file']

    if file:
        conn = psycopg2.connect(database="nuciferaDB", user="postgres", password="9221", host="localhost", port="8080")
        cursor = conn.cursor()

        #create schema
        schema_number = "1"
        create_schema_query = f"CREATE SCHEMA IF NOT EXISTS batch{schema_number};"
        cursor.execute(create_schema_query)
        conn.commit()

        #create tables
        create_tables_query = f'''
        CREATE TABLE IF NOT EXISTS batch{schema_number}.original (
            Date varchar(30) NOT NULL,
            Average_Price real,
            Rainfall_Kurunegala real,
            Rainfall_Puttalam real,
            Rainfall_Colombo real,
            Temp_Kurunegala real,
            Temp_Puttalam real,
            Temp_Colombo real,
            PRIMARY KEY (Date)
        );

        CREATE TABLE IF NOT EXISTS batch{schema_number}.models (
            Model_Id integer NOT NULL,
            Model_Name varchar(50),
            Plot_Fit bytea,
            Plot_Validation bytea,
            no_features integer,
            feature_list TEXT [],
            PRIMARY KEY (Model_Id)
        );

        CREATE TABLE IF NOT EXISTS batch{schema_number}.predictions (
            Model_Id integer NOT NULL,
            Date integer,
            Price real,
            CONSTRAINT fk_model_Id
                FOREIGN KEY(Model_Id)
                REFERENCES models(Model_Id)
        );

        CREATE INDEX index_predictions ON batch{schema_number}.predictions (Model_Id);       
        '''
        cursor.execute(create_tables_query)
        conn.commit()

        csv_data = csv.reader(file)
        next(csv_data)
        for row in csv_data:
            cursor.execute('''INSERT INTO batch%s.original(Date, Average_Price, Rainfall_Kurunegala,
              Rainfall_Puttalam, Rainfall_Colombo, Temp_Kurunegala, Temp_Puttalam, Temp_Colombo) VALUES (%s, %s, %s, %s, %s, %s, %s, %s)''', '1', row)
            
        conn.commit()
        cursor.close()
        conn.close()

        return 'file uploaded successfully'
    
    return 'no file selected'

@app.route('/predict', methods=['GET'])
def predict():
    return predictLSTM()