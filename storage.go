package main

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"log"
	"strconv"
)

var db *sql.DB

func InitStorage() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", conf.Bot.Resources.Database)
	checkErr(err)

	createSqlQuery := "CREATE TABLE IF NOT EXISTS users (" +
					"	`id` INTEGER NOT NULL PRIMARY KEY," +
					"	`login` TEXT NOT NULL," +
					"	`name` TEXT NOT NULL," +
					"	`status` VARCHAR(128) NOT NULL," +
					"	`selected_game_id` INTEGER NOT NULL," +
					"	`cur_game_id` INTEGER NOT NULL," +
					"	`cur_question_id` INTEGER NOT NULL," +
					"	`cur_attempt` INTEGER NOT NULL," +
					"	`cur_hint_attempt` INTEGER NOT NULL," +
					"	`cur_points` INTEGER NOT NULL" +
					")"

	db.Exec(createSqlQuery, nil)

	return db
}

func InsertUser(user *User) {
	log.Println("Insert new user with id " + strconv.Itoa(user.ID))
	stmt, err := db.Prepare("INSERT INTO users (id,login,name,status,selected_game_id,cur_game_id,cur_question_id, cur_attempt, cur_hint_attempt, cur_points) VALUES (?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(
		user.ID,
		user.Login,
		user.Name,
		user.Status,
		user.CurrentGameId,
		user.SelectedGameId,
		user.CurrentQuestionId,
		user.CurrentAttempt,
		user.CurrentHintAttempt,
		user.CurrentPoints,
	)

	checkErr(err)
}

func UpdateUser(user *User) {
	stmt, err := db.Prepare("UPDATE users SET login=?,name=?,status=?,selected_game_id=?,cur_game_id=?,cur_question_id=?, cur_attempt=?,cur_hint_attempt=?,cur_points=? WHERE id = ?")
	checkErr(err)

	_, err = stmt.Exec(
		user.Login,
		user.Name,
		user.Status,
		user.SelectedGameId,
		user.CurrentGameId,
		user.CurrentQuestionId,
		user.CurrentAttempt,
		user.CurrentHintAttempt,
		user.CurrentPoints,
		user.ID,
	)

	checkErr(err)
}

func GetUser(id int) *User {
	log.Printf("Trying to get user by id=%d\n", id)
	rows, err := db.Query("SELECT id,login,name, status, selected_game_id, cur_game_id, cur_question_id, cur_attempt, cur_hint_attempt, cur_points FROM users WHERE id = ?", id)
	checkErr(err)

	var user User = User{}
	if rows.Next() {
		user = User{}
		err = rows.Scan(
			&user.ID,
			&user.Login,
			&user.Name,
			&user.Status,
			&user.SelectedGameId,
			&user.CurrentGameId,
			&user.CurrentQuestionId,
			&user.CurrentAttempt,
			&user.CurrentHintAttempt,
			&user.CurrentPoints,
		)
		rows.Close()
		return &user
	}

	return nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
