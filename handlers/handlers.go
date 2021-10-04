package handlers

import (
	"encoding/json"
	"fmt"
	"golang-coding-challenge/models"
	"golang-coding-challenge/repositories"
	"golang-coding-challenge/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	db                 *gorm.DB
	transactionUsecase usecases.TransactionUseCase
)

func checkError(w http.ResponseWriter, err error, message string) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": message,
		})
		return
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage worked")
}

func GetTransactionsOfAnUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	queries := r.URL.Query()
	var (
		user_id, account_id int
		err                 error
	)

	user_id, err = strconv.Atoi(params["user_id"])
	checkError(w, err, "Invalid User Id")

	if len(queries["account_id"]) > 0 {
		account_id, err = strconv.Atoi(queries["account_id"][0])
		checkError(w, err, "Invalid Account Id")
	}

	transactionJsonList, err := transactionUsecase.GetTransactions(user_id, account_id)
	checkError(w, err, "Not Found Account")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactionJsonList)
}

func PostTransactionsOfAnUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var (
		transactionRequest usecases.CreateTransactionRequest
		user_id            int
		err                error
	)

	user_id, err = strconv.Atoi(params["user_id"])
	checkError(w, err, "Invalid User Id")

	json.NewDecoder(r.Body).Decode(&transactionRequest)

	newTransactionJson, err := transactionUsecase.CreateTransaction(transactionRequest, user_id)
	if err != nil {
		if err.Error() == fmt.Errorf("not enough balance").Error() {
			checkError(w, err, "not enough balance")
		} else if err.Error() == fmt.Errorf("can not create transaction").Error() {
			checkError(w, err, "can not create transaction")
		} else {
			checkError(w, err, "Not Found Account")
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTransactionJson)
}

func init() {
	db = models.ConnectDB()
	accountRepo := repositories.NewAccountRepo(db)
	transactionRepo := repositories.NewTransactionRepo(db)
	transactionUsecase = usecases.NewTransactionUsecase(accountRepo, transactionRepo)
}

func GetTransactionsOfAnUserV2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	queries := r.URL.Query()
	var (
		user_id, account_id int
		err                 error
		page                int = 1
		size                int = 5
	)

	user_id, err = strconv.Atoi(params["user_id"])
	checkError(w, err, "Invalid User Id")

	if len(queries["account_id"]) > 0 {
		account_id, err = strconv.Atoi(queries["account_id"][0])
		checkError(w, err, "Invalid Account Id")
	}

	if len(queries["account_id"]) > 0 {
		account_id, err = strconv.Atoi(queries["account_id"][0])
		checkError(w, err, "Invalid Account Id")
	}

	if len(queries["page"]) > 0 {
		page, err = strconv.Atoi(queries["page"][0])
		checkError(w, err, "Error Argument")
	}

	if len(queries["size"]) > 0 {
		size, err = strconv.Atoi(queries["size"][0])
		checkError(w, err, "Error Argument")
	}

	transactionJsonList, err := transactionUsecase.GetTransactionsV2(user_id, account_id, page, size)
	checkError(w, err, "Not Found Account")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactionJsonList)
}
