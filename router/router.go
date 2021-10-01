package router

import (
	"golang-coding-challenge/handlers"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handlers.HomePage).Methods("GET")
	router.HandleFunc("/api/users/{user_id}/transactions", handlers.GetTransactionsOfAnUser).Methods("GET")
	router.HandleFunc("/api/users/{user_id}/transactions", handlers.PostTransactionsOfAnUser).Methods("POST")

	router.HandleFunc("/api/v2/users/{user_id}/transactions", handlers.GetTransactionsOfAnUserV2).Methods("GET")

	return router
}
