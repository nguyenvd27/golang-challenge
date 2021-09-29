package repositories

import "gorm.io/gorm"

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

type AccountRepo interface {
	GetAccountIDsByUserIDAccountID(userID, accountID int) ([]int, error)
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
