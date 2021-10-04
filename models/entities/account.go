package entities

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name         string  `gorm:"column:name"`
	Bank         string  `gorm:"column:bank"`
	Balance      float64 `gorm:"column:balance"`
	UserID       uint    `gorm:"column:user_id"`
	Transactions []Transaction
}
