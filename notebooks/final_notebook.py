# -*- coding: utf-8 -*-
"""FinalWithTests.ipynb

Automatically generated by Colaboratory.

Original file is located at
    https://colab.research.google.com/drive/1W9ftoxK6R3GKs6ks0Ryj2Hy4AV5gHGWl
"""

import numpy as np
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import LSTM, LeakyReLU
from tensorflow.keras.layers import Dense, Dropout
import pandas as pd
from matplotlib import pyplot as plt
from sklearn.preprocessing import StandardScaler
from sklearn.model_selection import train_test_split
import seaborn as sns
import psycopg2
from sklearn.metrics import mean_squared_error, accuracy_score, mean_absolute_error

df = pd.read_csv('data_2.0.csv',  thousands=',')
df

# conn = psycopg2.connect(database="nuciferaDB", user="postgres", password="9221", host="localhost", port="55432")
# cursor = conn.cursor()
# query = "SELECT * FROM batch1.original ORDER BY Average_Price DESC;"
# df = pd.read_sql_query(query, conn)
# df
# td = pd.to_datetime(df['date'], dayfirst=True, unit='s')
# train_dates = td.sort_values(ascending=False)
# train_dates

td = pd.to_datetime(df['Date'], dayfirst=True)
train_dates = td.sort_values(ascending=False)
train_dates[0].timestamp()

cols = list(df)[1:]
print(cols)
df_for_training = df[cols].astype(float)

scaler = StandardScaler()
scaler = scaler.fit(df_for_training)
df_for_training_scaled = scaler.transform(df_for_training)
scaled = df_for_training_scaled

trainX = []
trainY = []

n_future = 1   # Number of weeks we want to look into the future based on the past weeks.
n_past = 4  # Number of past weeks we want to use to predict the future.

#Reformat input data into a shape: (n_samples x timesteps x n_features)
#In my example, my df_for_training_scaled has a shape (12823, 5)
#12823 refers to the number of data points and 5 refers to the columns (multi-variables).
for i in range(n_past, len(df_for_training_scaled) - n_future +1):
    trainX.append(df_for_training_scaled[i - n_past:i, 0:df_for_training.shape[1]])
    trainY.append(df_for_training_scaled[i + n_future - 1:i + n_future, 0])

# test_size = 0.2 #Test set proportion
# random_state = 42 #seed for reproducibility

# trainX, testX, trainY, testY = train_test_split(trainX, trainY, test_size=test_size, random_state=random_state)

# trainX, trainY, testX, testY = np.array(trainX), np.array(trainY), np.array(testX), np.array(testY)

trainX, trainY= np.array(trainX), np.array(trainY)

print('trainX shape == {}.'.format(trainX.shape))
print('trainY shape == {}.'.format(trainY.shape))
# print('trainX shape == {}.'.format(testX.shape))
# print('trainY shape == {}.'.format(testY.shape))

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

from pandas.tseries.holiday import USFederalHolidayCalendar
from pandas.tseries.offsets import CustomBusinessDay
us_bd = CustomBusinessDay(calendar=USFederalHolidayCalendar())

n_past = 200
n_weeks_for_prediction=100
#prediction period should be from 2023 5 1 - 7- 14 - 21 - 27
predict_period_dates = pd.date_range(list(train_dates)[n_past], periods=n_weeks_for_prediction, freq='W-SUN').tolist()
predict_period_dates

prediction = model.predict(trainX[:n_weeks_for_prediction])
 #shape = (n, 1) where n is the n_days_for_prediction

prediction_copies = np.repeat(prediction, df_for_training.shape[1], axis=-1)
y_pred_future = scaler.inverse_transform(prediction_copies)[:,0]
y_pred_future

# Convert timestamp to date
forecast_dates = []
for time_i in predict_period_dates:
    forecast_dates.append(time_i.date())

df_forecast = pd.DataFrame({'Date':np.array(forecast_dates), 'Average_Price':y_pred_future})
df_forecast['Date']=pd.to_datetime(df_forecast['Date'], dayfirst=True)

forecast_dates

original = df[['Date', 'Average_Price']]
original.head
original['Date']=pd.to_datetime(original['Date'], dayfirst=True)
original = original.loc[original['Date'] >= '07/05/2020']

sns.lineplot(x=original['Date'], y=original['Average_Price'])
sns.lineplot(x=df_forecast['Date'], y=df_forecast['Average_Price'])

"""

---




# ***Inference***






---






"""

testx = []
testy = []

n_future = 1   # Number of weeks we want to look into the future based on the past weeks.
n_past = 4

for i in range(n_past, len(scaled) - n_future +1):
    testx.append(scaled[i - n_past:i, 0:df_for_training.shape[1]])
    testy.append(scaled[i + n_future - 1:i + n_future, 0])

testx= np.array(testx)
print('testx shape == {}.'.format(testx.shape))

prediction = model.predict(testx)

prediction_copies = np.repeat(prediction, scaled.shape[1], axis=-1)
y_pred_future = scaler.inverse_transform(prediction_copies)[:,0]

actual_values = df.iloc[n_past:,1]
dates = df.iloc[n_past:,0]

original_datetime_values = df.iloc[n_past:, 0]

# Plot the actual prices using a black line
plt.plot(dates,actual_values, color='black', label="Actual price")

# Plot the predicted prices using a green line
plt.plot(original_datetime_values,y_pred_future, color='green', label="Predicted price")

# Set the title of the plot using the company name
plt.title("share price")

# Set the x-axis label as 'time'
plt.xlabel("time")
# Set the y-axis label using the company name
plt.ylabel("coconut price")

# Display a legend to differentiate the actual and predicted prices
plt.legend()

# Show the plot on the screen
plt.show()

# Calculate Mean Squared Error (MSE) and Root Mean Squared Error (RMSE)
mse = mean_squared_error(actual_values, y_pred_future)
rmse = np.sqrt(mse)
mae = mean_absolute_error(actual_values, y_pred_future)
# Calculate Mean Squared Error (MSE) manually
Mape = np.mean(np.abs((actual_values - y_pred_future)/actual_values)) * 100

print("Mean Squared Error (MSE):", mse)
print("Root Mean Squared Error (RMSE):", rmse)
print("Mean Absolute Error (MAE):", mae)
print("Mape: ", Mape)
