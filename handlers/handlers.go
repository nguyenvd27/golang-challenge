package handlers

import (
	"encoding/json"
	"fmt"
	"golang-coding-challenge/models"
	"golang-coding-challenge/models/entities"
	"golang-coding-challenge/transfers"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage worked")
}

func GetTransactionsOfAnUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	queries := r.URL.Query()

	var (
		transactions     []entities.Transaction
		query_account_id []int
	)

	if queries["account_id"] == nil {
		db.Table("accounts").Select("id").Where("user_id = ?", params["user_id"]).Scan(&query_account_id)
		db.Preload(clause.Associations).Where("account_id IN ?", query_account_id).Find(&transactions)
	} else {
		db.Table("accounts").Select("id").Where("id = ? AND user_id = ?", queries["account_id"], params["user_id"]).Scan(&query_account_id)
		db.Preload(clause.Associations).Where("account_id IN ?", query_account_id).Find(&transactions)
	}

	w.WriteHeader(http.StatusOK)
	if len(transactions) > 0 {
		transactionJsonList := transfers.GetTransactionsJsonList(transactions)
		json.NewEncoder(w).Encode(transactionJsonList)
	} else {
		json.NewEncoder(w).Encode(transactions)
	}
}

func PostTransactionsOfAnUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var (
		transaction      entities.Transaction
		query_account_id []int
	)
	json.NewDecoder(r.Body).Decode(&transaction)

	db.Table("accounts").Select("id").Where("id = ? AND user_id = ?", transaction.AccountID, params["user_id"]).Scan(&query_account_id)
	if len(query_account_id) > 0 {
		db.Create(&transaction)
		db.Preload(clause.Associations).Find(&transaction)
		transactionJson := transfers.GetTransactionsJson(transaction)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(transactionJson)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Not Found Account",
		})
	}
}

func init() {
	db = models.ConnectDB()
}
