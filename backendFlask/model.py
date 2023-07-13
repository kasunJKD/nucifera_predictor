import numpy as np
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import LSTM, LeakyReLU
from tensorflow.keras.layers import Dense, Dropout
import pandas as pd
from matplotlib import pyplot as plt
from sklearn.preprocessing import StandardScaler
import seaborn as sns
import psycopg2
import base64
import io

def predictLSTM ():
    conn = psycopg2.connect(database="nuciferaDB", user="postgres", password="9221", host="localhost", port="8080")
    cursor = conn.cursor()
    query = "SELECT * FROM batch1.original;"
    df = pd.read_sql_query(query, conn)

    train_dates = pd.to_datetime(df['Date'], dayfirst=True)

    cols = list(df)[1:5]
    df_for_training = df[cols].astype(float)

    scaler = StandardScaler()
    scaler = scaler.fit(df_for_training)
    df_for_training_scaled = scaler.transform(df_for_training)

    trainX = []
    trainY = []

    n_future = 1   # Number of weeks we want to look into the future based on the past weeks.
    n_past = 8  # Number of past weeks we want to use to predict the future.

    for i in range(n_past, len(df_for_training_scaled) - n_future +1):
        trainX.append(df_for_training_scaled[i - n_past:i, 0:df_for_training.shape[1]])
        trainY.append(df_for_training_scaled[i + n_future - 1:i + n_future, 0])

    trainX, trainY = np.array(trainX), np.array(trainY)

    # define the Autoencoder model
    model = Sequential()
    model.add(LSTM(64, activation='relu', input_shape=(trainX.shape[1], trainX.shape[2]), return_sequences=True))
    model.add(LSTM(32, activation='relu', return_sequences=False))
    model.add(Dropout(0.3))
    model.add(Dense(trainY.shape[1]))

    model.compile(optimizer='adam', loss='mse')
    model.summary()

    # fit the model
    history = model.fit(trainX, trainY, epochs=50, batch_size=16, validation_split=0.1, verbose=1)
    plt.plot(history.history['loss'], label='Training loss')
    plt.plot(history.history['val_loss'], label='Validation loss')
    plt.legend()

    buffer = io.BytesIO()
    plt.savefig(buffer, format='png', bbox_inches='tight')
    buffer.seek(0)
    image_base64_validation = base64.b64encode(buffer.getvalue()).decode('utf-8')

    n_past = 200
    n_weeks_for_prediction=100
    #prediction period should be from 2023 5 1 - 7- 14 - 21 - 27
    predict_period_dates = pd.date_range(list(train_dates)[1], periods=n_weeks_for_prediction, freq='W-SUN').tolist()

    prediction = model.predict(trainX[:n_weeks_for_prediction])

    prediction_copies = np.repeat(prediction, df_for_training.shape[1], axis=-1)
    y_pred_future = scaler.inverse_transform(prediction_copies)[:,0]

    forecast_dates = []
    for time_i in predict_period_dates:
        forecast_dates.append(time_i.date())

    df_forecast = pd.DataFrame({'Date':np.array(forecast_dates), 'Average_Price':y_pred_future})
    df_forecast['Date']=pd.to_datetime(df_forecast['Date'], dayfirst=True)

    original = df[['Date', 'Average_Price']]
    original.head
    original['Date']=pd.to_datetime(original['Date'], dayfirst=True)

    sns.lineplot(x=original['Date'], y=original['Average_Price'])
    sns.lineplot(x=df_forecast['Date'], y=df_forecast['Average_Price'])

    buffer2 = io.BytesIO()
    plt.savefig(buffer2, format='png', bbox_inches='tight')
    buffer2.seek(0)
    image_base64_fit = base64.b64encode(buffer2.getvalue()).decode('utf-8')

    cursor.execute('''
        INSERT INTO batch1.models (Model_Id, Model_Name, Plot_Fit, Plot_Validation, no_features, feature_list)
        VALUES (%d, %s, %s, %s, %d, %s)
    ''', 1, "LSTM", image_base64_fit,image_base64_validation, 4, ['test', 'test'])

    conn.commit()

    for index, row in df_forecast.iterrows():
        values = tuple(row)
        query = f"INSERT INTO batch1.predictions VALUES 1, {values}"
        cursor.execute(query)
    
    conn.commit()

    cursor.close()
    conn.close()

    return 'LSTM fired'





