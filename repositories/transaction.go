package repositories

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"golang-coding-challenge/models/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepo interface {
	GetTransactionsByAccountIDs(accountIDs []int) ([]entities.Transaction, error)
	CreateTransaction(trans entities.Transaction) (entities.Transaction, error)

	GetTransactionsWithPaginate(accountIDs []int, page, size int) ([]entities.Transaction, int64, error)
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

func (transaction *transactionDB) GetTransactionsWithPaginate(accountIDs []int, page, size int) ([]entities.Transaction, int64, error) {
	var (
		transactions []entities.Transaction
		err          error
		total        int64
	)

	err = transaction.db.Preload(clause.Associations).
		Where("account_id IN ?", accountIDs).
		Limit(size).Offset((page - 1) * size).
		Order("id asc").Find(&transactions).Error

	transaction.db.Table("transactions").Where("account_id IN ?", accountIDs).Count(&total)

	return transactions, total, err
}
