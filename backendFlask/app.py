from flask import Flask, request, render_template
import csv
import psycopg2
from model import predictLSTM
import sys
import codecs
import datetime

app = Flask(__name__)

@app.route('/upload', methods=['POST'])
def upload():
    file = request.files['csv_file']
    if file:
        conn = psycopg2.connect(database="nuciferaDB", user="postgres", password="9221", host="nucifera-db", port="5432")
        cursor = conn.cursor()
        #create schema
        schema_number = "1"
        create_schema_query = f"CREATE SCHEMA IF NOT EXISTS batch{schema_number};"
        cursor.execute(create_schema_query)
        conn.commit()

        #create tables
        create_tables_query = f'''
        CREATE TABLE IF NOT EXISTS batch{schema_number}.original (
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
                REFERENCES batch{schema_number}.models(Model_Id)
        );     
        '''
        cursor.execute(create_tables_query)
        conn.commit()
        stream = codecs.iterdecode(file.stream, 'utf-8')
        # Read the CSV file
        csv_data = csv.reader(stream, dialect=csv.excel)
        # Skip the first row (header)
        next(csv_data)
        for row in csv_data:  
            # Convert the string to a datetime object
            date_string = row[0].replace(" ", "")  # Remove white spaces from the date_string
            datetime_obj = datetime.datetime.strptime(date_string,"%d/%m/%Y")
            # Convert the datetime object to Unix timestamp
            unix_timestamp = datetime_obj.timestamp()

            dd = int(unix_timestamp)
            ap = float(row[1])
            rk = float(row[2])
            rp = float(row[3])
            rc = float(row[4])
            tk = float(row[5])
            tp = float(row[6])
            tc = float(row[7])

            insert_statment = f'INSERT INTO batch{schema_number}.original(Date, Average_Price, Rainfall_Kurunegala,Rainfall_Puttalam, Rainfall_Colombo, Temp_Kurunegala, Temp_Puttalam,Temp_Colombo) VALUES (%s, %s, %s, %s, %s, %s, %s, %s)'
            cursor.execute(insert_statment, (dd, ap, rk, rp, rc, tk, tp, tc))
            
        conn.commit()
        cursor.close()
        conn.close()

        return 'file uploaded successfully'
    
    return 'no file selected'

@app.route('/predict', methods=['GET'])
def predict():
    return predictLSTM()

if __name__ == "__main__":
    app.run(host='0.0.0.0',debug=True,port='5000')
