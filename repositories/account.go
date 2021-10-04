package repositories

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"fmt"
	"golang-coding-challenge/models/entities"

	"gorm.io/gorm"
)

type AccountRepo interface {
	GetAccountIDsByUserIDAccountID(userID, accountID int) ([]int, error)
	GetAccountUser(userID, accountID int) (entities.Account, error)
	UpdateAccountUser(accountUser entities.Account, newBalance float64) (entities.Account, error)
}

type accountDB struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) AccountRepo {
	return &accountDB{
		db: db,
	}
}

func (account *accountDB) GetAccountIDsByUserIDAccountID(userID, accountID int) ([]int, error) {
	var (
		account_ids []int
		err         error
	)

	if accountID > 0 {
		err = account.db.Table("accounts").Select("id").Where("id = ? AND user_id = ?", accountID, userID).Scan(&account_ids).Error
	} else {
		err = account.db.Table("accounts").Select("id").Where("user_id = ?", userID).Scan(&account_ids).Error
	}

	return account_ids, err
}

func (account *accountDB) GetAccountUser(userID, accountID int) (entities.Account, error) {
	var (
		accountUser entities.Account
		err         error
	)

	if accountID > 0 {
		err = account.db.Table("accounts").Where("id = ? AND user_id = ?", accountID, userID).First(&accountUser).Error
	} else {
		err = fmt.Errorf("error argument")
	}

	return accountUser, err
}

func (account *accountDB) UpdateAccountUser(accountUser entities.Account, newBalance float64) (entities.Account, error) {
	var err error

	accountUser.Balance = newBalance

	err = account.db.Save(&accountUser).Error

	return accountUser, err
}
