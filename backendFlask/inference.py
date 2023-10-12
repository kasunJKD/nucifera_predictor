import numpy as np
from matplotlib import pyplot as plt
import base64
import io

def inference (df, scaled, df_for_training, model, scaler, train_dates):
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

    plt.figure(figsize=(10, 6))
    # Plot the actual prices using a black line
    plt.plot(train_dates[n_past:],actual_values, color='black', label="Actual price")

    # Plot the predicted prices using a green line
    plt.plot(train_dates[n_past:],y_pred_future, color='green', label="Predicted price")

    # Set the title of the plot using the company name
    plt.title("share price")

    # Set the x-axis label as 'time'
    plt.xlabel("time")
    # Set the y-axis label using the company name
    plt.ylabel("coconut price")

    buffer = io.BytesIO()
    plt.savefig(buffer, format='png')
    buffer.seek(0)
    acual_predicted_graph_lstm = base64.b64encode(buffer.read()).decode()
    plt.close()
    buffer.close()

    # Calculate Mean Squared Error (MSE) manually
    Mape = np.mean(np.abs((actual_values - y_pred_future)/actual_values)) * 100

    return Mape, acual_predicted_graph_lstm
