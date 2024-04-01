package config

import (
  "fmt"
  "time"
  //"database/sql"
  	_ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
  )


const dbAddr = "127.0.0.1";
const dbName = "";
const  dbUser = "";
const dbPasswd = "";
const dbPort = "3306"

func DbConnection() *sqlx.DB {
  dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
