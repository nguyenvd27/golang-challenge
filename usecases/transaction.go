package usecases

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"fmt"
	"golang-coding-challenge/models/entities"
	"golang-coding-challenge/repositories"
	"golang-coding-challenge/transfers"

	"gorm.io/gorm"
)

type TransactionUseCase interface {
	GetTransactions(userID, accountID int) ([]transfers.TransactionJson, error)
	CreateTransaction(reqBody CreateTransactionRequest, userID int) (*transfers.TransactionJson, error)
}

type transactionUsecase struct {
	accountRepo     repositories.AccountRepo
	transactionRepo repositories.TransactionRepo
}

type CreateTransactionRequest struct {
	AccountID       int     `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func NewTransactionUsecase(db *gorm.DB) TransactionUseCase {
	return &transactionUsecase{
		accountRepo:     repositories.NewAccountRepo(db),
		transactionRepo: repositories.NewTransactionRepo(db),
	}
}

func (uc *transactionUsecase) GetTransactions(userID, accountID int) ([]transfers.TransactionJson, error) {
	account_ids, err := uc.accountRepo.GetAccountIDsByUserIDAccountID(userID, accountID)
	if err != nil {
		return nil, err
	}

	transactions, err := uc.transactionRepo.GetTransactionsByAccountIDs(account_ids)
	if err != nil {
		return nil, err
	}

	if len(transactions) > 0 {
		transactionJsonList := transfers.GetTransactionsJsonList(transactions)
		return transactionJsonList, nil
	}

	return []transfers.TransactionJson{}, nil
}

func (uc *transactionUsecase) CreateTransaction(reqBody CreateTransactionRequest, userID int) (*transfers.TransactionJson, error) {
	account_ids, err := uc.accountRepo.GetAccountIDsByUserIDAccountID(userID, reqBody.AccountID)
	if err != nil {
		return nil, err
	}

	if len(account_ids) > 0 {
		new_transaction, err := uc.transactionRepo.CreateTransaction(entities.Transaction{
			AccountID:       uint(reqBody.AccountID),
			Amount:          reqBody.Amount,
			TransactionType: reqBody.TransactionType,
		})
		if err != nil {
			return nil, err
		}
		transactionJson := transfers.GetTransactionsJson(new_transaction)
		return &transactionJson, nil
	}

	return nil, fmt.Errorf("not found user account")
}
