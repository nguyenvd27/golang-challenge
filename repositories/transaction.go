package repositories

import (
	"golang-coding-challenge/models/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepo interface {
	GetTransactionsByAccountIDs(accountIDs []int) ([]entities.Transaction, error)
	CreateTransaction(trans entities.Transaction) (entities.Transaction, error)
}

type transactionDB struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepo {
	return &transactionDB{
		db: db,
	}
}

func (transaction *transactionDB) GetTransactionsByAccountIDs(accountIDs []int) ([]entities.Transaction, error) {
	var (
		transactions []entities.Transaction
		err          error
	)

	err = transaction.db.Preload(clause.Associations).Where("account_id IN ?", accountIDs).Find(&transactions).Error

	return transactions, err
}

func (transaction *transactionDB) CreateTransaction(trans entities.Transaction) (entities.Transaction, error) {
	err := transaction.db.Create(&trans).Error
	transaction.db.Preload(clause.Associations).Find(&trans)

	return trans, err
}
