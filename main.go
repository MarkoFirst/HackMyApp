package main

import (
	. "./routing"
	"fmt"
	"net/http"
	"text/template"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/julienschmidt/httprouter"
)

var (
	db, _ = sql.Open("sqllite3", "cache/web.db")
)

//_______________________________________База________________________________________________________________________

func createDataBase() { // Создание БД
	db.Exec("CREATE TABLE if NOT EXISTS users (user_id INTEGER, user_name TEXT, password TEXT)")
	db.Exec("CREATE TABLE if NOT EXISTS bank_cards (user_id INTEGER, currency TEXT, balance REAL)")
	db.Exec("CREATE TABLE if NOT EXISTS shop_account (user_id INTEGER, user_name TEXT, password TEXT, balance REAL)")
	db.Exec("CREATE TABLE if NOT EXISTS blog (user_id INTEGER, blog_id INTEGER, balance REAL)")
	db.Exec("CREATE TABLE if NOT EXISTS questions (user_id INTEGER, blog_id INTEGER, question TEXT, question_id INTEGER )")
	db.Exec("CREATE TABLE if NOT EXISTS answers (user_id INTEGER, question_id INTEGER , answer TEXT)")
}

func insertDB( into string, what string, value string) error { // Запрос на добавление в базу
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO " + into + " (" + what + ") VALUE (" + value + ")")
	_, err := stmt.Exec()
	tx.Commit()

	return err
}

func selectDB(who string, from string, where string) string { // Запрос поиска в базе
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

// __________________________________________________________________________________________________________________

func searchInDb(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	search := r.FormValue("search")

	fmt.Println(search)

	t, err := template.ParseFiles("templates/search.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Println(err.Error())
	}

	t.ExecuteTemplate(w, "search", search)
}

func main() {
	fmt.Println("Server run")

	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	router.GET("/", IndexHandler)
	router.GET("/shop", ShopHandler)
	router.GET("/blog", BlogHandler)
	router.GET("/bank", BankHandler)
	router.GET("/kript", KriptHandler)
	router.GET("/error", ErrorHandler)
	router.GET("/search", searchInDb)

	log.Fatal(http.ListenAndServe(":3000", router))
}
