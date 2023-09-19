from flask import Flask, request, render_template
import csv
import psycopg2
from model import predictLSTM, predictGRU, predict1D
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
            Min_Temp_Kurunegala real,
            Min_Temp_Puttalam real,
            Min_Temp_Colombo real,
            Max_Temp_Kurunegala real,
            Max_Temp_Puttalam real,
            Max_Temp_Colombo real,
            PRIMARY KEY (Date)
        );

        CREATE TABLE IF NOT EXISTS batch{schema_number}.models (
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
            mtk = float(row[8])
            mtp = float(row[9])
            mtc = float(row[10])

            insert_statment = f'INSERT INTO batch{schema_number}.original(Date, Average_Price, Rainfall_Kurunegala,Rainfall_Puttalam, Rainfall_Colombo, Min_Temp_Kurunegala, Min_Temp_Puttalam,Min_Temp_Colombo,Max_Temp_Kurunegala, Max_Temp_Puttalam,Max_Temp_Colombo) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)'
            cursor.execute(insert_statment, (dd, ap, rk, rp, rc, tk, tp, tc, mtk, mtp, mtc))
            
        conn.commit()
        cursor.close()
        conn.close()

        return 'file uploaded successfully'
    
    return 'no file selected'

@app.route('/predict_lstm', methods=['GET'])
def predictLstm():
    return predictLSTM()

@app.route('/predict_gru', methods=['GET'])
def predictGru():
    return predictGRU()

@app.route('/predict_1d', methods=['GET'])
def predict1Dtst():
    return predict1D()

if __name__ == "__main__":
    app.run(host='0.0.0.0',debug=True,port='5000')
