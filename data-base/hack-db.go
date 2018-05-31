package data_base

import "database/sql"

var db, _ = sql.Open("sqllite3", "cache/web.db")

func CreateDataBase() {
	db.Exec("CREATE TABLE if NOT EXISTS users (user_id INTEGER, user_name TEXT, password TEXT)")
	db.Exec("CREATE TABLE if NOT EXISTS bank_cards (user_id INTEGER, currency TEXT, balance REAL)")
	db.Exec("CREATE TABLE if NOT EXISTS shop_account (user_id INTEGER, user_name TEXT, password TEXT, balance REAL)")
	db.Exec("CREATE TABLE if NOT EXISTS blog (user_id INTEGER, blog_id INTEGER, balance REAL)")
	db.Exec("CREATE TABLE if NOT EXISTS questions (user_id INTEGER, blog_id INTEGER, question TEXT, question_id INTEGER )")
	db.Exec("CREATE TABLE if NOT EXISTS answers (user_id INTEGER, question_id INTEGER , answer TEXT)")
}

func InsertDB( into string, what string, value string) error {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO " + into + " (" + what + ") VALUE (" + value + ")")
	_, err := stmt.Exec()
	tx.Commit()

	return err
}

func SelectDB(who string, from string, where string) string {
	result := ""
	q, err := db.Query("SELECT " + who + " FROM " + from + " WHERE " + where)

	if err != nil {
		return err.Error()
	}

	for q.Next() {
		q.Scan(&result)
	}

	return result
}
