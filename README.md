# nucifera_predictor
Final year research

--> postgres as db

--> Python model for admin purposes and running new data sets manually
    - read data from db (original data)
    - run the prediction model and save new data to db
    - add new features and update them to the template

--> Golang Backend
    - Get db data
    - send sms and email with data to the date
    - Oauth

--> angular frontend
    - simple frontend to show dashboard
    - form to add email or sms for subscription 

## Architecture
![alt text](https://github.com/kasunJKD/nucifera_predictor/blob/main/docs/architecture.png "architecture")

## Database Schema
![alt text](https://github.com/kasunJKD/nucifera_predictor/blob/main/docs/dbSchema.png "Db schema")

--- 

### TODO

    - [x] python_ model for lstm
    - [x] python_ store predicted data in db
    - [x] store batch by batch with plots
    - [x] ----> testing check point-----> DONE
    - [x] ----> running model----> DONE
    - [x] python_ model for GRU
    - [x] python_ model for 1D
    - [x] python_ upload csv file and read csv data from database
    - [x] add mean squred error per batch per model 
    - [ ] python_ number of feature select option (notImportant)
    - [ ] python_ feature select option (notImportant)
    - [ ] download csv format if user wanted to add a new batch
    - [x] everytime new upload happening new schema is created
    - [x] buttons added to fire models in admin mode

    - [x] go_google auth for register and login users
    - [x] create db connection for model database
    - [x] select all data from batch data and prediction retrival
    - [ ] subscribe for sms service for weekly updates

    - [x] angular_basic ui with form
    - [x] added tailwind
    - [ ] display plots for batches 
    - [ ] show predicted data per each model
    - [ ] show original data used 

    - [x] db schema added

### Bug fixes
    - [x] flask_ db inserting data type errors fixed
    - [x] flask_ db Date data type changed to unix time fixed
    - [x] flask_ strings convert to floats fixed
    - [x] lstm model function issues fixed
    - [x] login issue in golang fixed
    - [x] 1D is not working after getting data from database
    - [x] golang endpoint issues fixed
### Bugs

    
