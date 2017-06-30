package main

import (
	"log"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)

var db *sql.DB

func InitStorage() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", conf.Bot.Resources.Database)
	checkErr(err)

	return db
}

func UpsertUser(user User) {

}

func GetUser(telegramName string) *User {

	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
