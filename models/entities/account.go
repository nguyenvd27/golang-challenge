package entities

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name         string `gorm:"column:name"`
	Bank         string `gorm:"column:bank"`
	UserID       uint   `gorm:"column:user_id"`
	Transactions []Transaction
}
