package main

import (
	"fmt"
	"golang-coding-challenge/router"
	"log"
	"net/http"
)

// var db *gorm.DB

func initRouter() {
	router := router.New()
	log.Fatal(http.ListenAndServe(":8001", router))
}

func main() {
	fmt.Println("Golang Coding Challenge")

	// connect Databse
	// db = models.ConnectDB()

	// init route
	initRouter()
}
