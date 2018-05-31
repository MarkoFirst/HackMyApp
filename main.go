package main

import (
	. "./routing"
	"fmt"
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
)

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
	router.GET("/search", SearchInDb)

	log.Fatal(http.ListenAndServe(":3000", router))
}
