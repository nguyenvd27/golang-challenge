package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// const USER_NAME = "root"
// const PASSWORD = ""
// const DATABASE_PATH = "tcp(127.0.0.1:3306)/golang_coding_challenge"

func ConnectDB() *gorm.DB {
	// dsn := USER_NAME + ":" + PASSWORD + "@" + DATABASE_PATH + "?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := os.Getenv("USER_NAME") + ":" + os.Getenv("PASSWORD") + "@" + DATABASE_PATH + "?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := "tester:secret@tcp(db:3306)/golang_coding_challenge" + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
