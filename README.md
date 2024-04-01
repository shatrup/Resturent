# Resturent
A RESTful API example for simple Resturents application with Go

#How to start

#structure
Resturent
├--- app
│   |---app.go               // Define the Route and expose the url
│--- domain
|   |--- domain.go           // json fields of Resturent
|--- dto
|   |--- resturentDetails.go  // Database filed structure
├--- config
│   └-- dbConnection.go        // Configuration of Database
|--- main.go
|--- go.mod
|

Restaurants details
##The endpoints and the corresponding REST operations are defined as follows:-

###RESTAURANTS
127.0.0.1:10000/resturents
GET : This method on above URL returns all the restaurants available in the database in json format
127.0.0.1:10000/resturents
POST : This method posts a new restaurant and accept application/JSON format for the operation
127.0.0.1:10000/resturents/{id}
Delete : This method deletes the given restaurant if the restaurant_id exists.
127.0.0.1:10000/resturents/{id}
Put : This method update entry the given restaurant if the restaurant_id exists.

#Mysql
MySQL is an open source SQL relational database management system that’s developed and supported by Oracle.
It is easy to use. This is based on RDBMS.
