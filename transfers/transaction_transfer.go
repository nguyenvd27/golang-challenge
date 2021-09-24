package transfers

import "golang-coding-challenge/models/entities"

type TransactionJson struct {
	Id              int     `json:"id"`
	AccountID       uint    `json:"account_id"`
	Amount          float64 `json:"amount"`
	Bank            string  `json:"bank"`
	TransactionType string  `json:"transaction_type"`
	CreatedAt       string  `json:"created_at"`
}

func GetTransactionsJsonList(transactions []entities.Transaction) []TransactionJson {
	var (
		transactionsJsonList []TransactionJson
		transactionJson      TransactionJson
	)
	for _, v := range transactions {
		transactionJson.Id = int(v.ID)
		transactionJson.AccountID = v.AccountID
		transactionJson.Amount = v.Amount
		transactionJson.Bank = v.Account.Bank
		transactionJson.TransactionType = v.TransactionType
		transactionJson.CreatedAt = v.CreatedAt.String()

		transactionsJsonList = append(transactionsJsonList, transactionJson)
	}
	return transactionsJsonList
}

func GetTransactionsJson(transaction entities.Transaction) TransactionJson {
	var transactionJson TransactionJson

	transactionJson.Id = int(transaction.ID)
	transactionJson.AccountID = transaction.AccountID
	transactionJson.Amount = transaction.Amount
	transactionJson.Bank = transaction.Account.Bank
	transactionJson.TransactionType = transaction.TransactionType
	transactionJson.CreatedAt = transaction.CreatedAt.String()

	return transactionJson
}
