package models

import (
    "database/sql"
    "os"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(){
    var err error
    DB, err = sql.Open("sqlite3", "./database/database.db")
    if err != nil{
        log.Fatal(err)
    }

    schema, err := os.ReadFile("./database/migrations/001_create_customers.sql")
    if err != nil{
        log.Fatal(err)
    }

    _, err = DB.Exec(string(schema))
    if err != nil{
        log.Fatal(err)
    }
}
