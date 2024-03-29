from flask import Flask, request, render_template, redirect, url_for
import csv
import psycopg2
from model import predictLSTM, predictGRU, predict1D
import sys
import codecs
import datetime


app = Flask(__name__)
# Define a secret key for the session
app.secret_key = 'your_secret_key'

@app.route("/")
def redirect_to_upload():
    return redirect(url_for("upload_file"))

@app.route("/functions")
def go_to_functions():
    return render_template('func.html')

# Define a route for the file upload page
@app.route('/upload', methods=['GET', 'POST'])
def upload_file():
    if request.method == 'POST':
        file = request.files['csv_file']
        if file:
            conn = psycopg2.connect(database="nuciferaDB", user="postgres", password="9221", host="nucifera-db", port="5432")
            cursor = conn.cursor()

            # Initialize schema_number
            schema_number = 1

            while True:
                # Define the schema name with the current schema_number
                schema_name = f"batch{schema_number}"
                
                # Check if the schema exists
                check_schema_query = f"SELECT schema_name FROM information_schema.schemata WHERE schema_name = '{schema_name}';"
                cursor.execute(check_schema_query)
                
                # Fetch the result (should be an empty list if schema doesn't exist)
                existing_schemas = cursor.fetchall()
                
                if not existing_schemas:
                    # Schema doesn't exist, create it
                    create_schema_query = f"CREATE SCHEMA {schema_name};"
                    cursor.execute(create_schema_query)
                    conn.commit()
                    break  # Exit the loop if the schema is created successfully
                else:
                    # Schema already exists, increment schema_number and try again
                    schema_number += 1

            # create tables
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

                insert_statement = f'INSERT INTO batch{schema_number}.original(Date, Average_Price, Rainfall_Kurunegala,Rainfall_Puttalam, Rainfall_Colombo, Min_Temp_Kurunegala, Min_Temp_Puttalam,Min_Temp_Colombo,Max_Temp_Kurunegala, Max_Temp_Puttalam,Max_Temp_Colombo) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)'
                cursor.execute(insert_statement, (dd, ap, rk, rp, rc, tk, tp, tc, mtk, mtp, mtc))

            conn.commit()
            cursor.close()
            conn.close()

            return redirect(url_for("go_to_functions"))

    return render_template('upload.html')


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
