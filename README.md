# nucifera_predictor
simple coconut price predictor (Final Year project)

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

## Database Schema
![alt text](https://github.com/kasunJKD/nucifera_predictor/blob/main/docs/dbSchema.png "Db schema")

--- 

TODO:
    - [x] python_ model for lstm
    - [x] python_ store predicted data in db
    - [ ] python_ model for GRU
    - [ ] python_ model for 1D
    - [x] store batch by batch with plots
    - [ ] python_ number of feature select option (notImportant)
    - [ ] python_ feature select option (notImportant)
    - [x] python_ upload csv file and read csv data from database 

    - [ ] go_google auth for register and login users
    - [ ] subscribe for sms service for weekly updates

    - [ ] angular_basic ui with form
    - [ ] display plots for batches 
    - [ ] show predicted data per each model
    - [ ] show original data used 

    - [x] db schema added
