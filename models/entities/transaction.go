package entities

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Amount          float64 `gorm:"column:amount" json:"amount"`
	TransactionType string  `gorm:"column:transaction_type" json:"transaction_type"`
	AccountID       uint    `gorm:"column:account_id" json:"account_id"`
	Account         Account `gorm:"foreignKey:AccountID"`
}
