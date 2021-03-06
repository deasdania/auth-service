package config

import (
	// "auth-service/api/models"
	"auth-service/seed"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"os"

	"github.com/jinzhu/gorm"
)

func InitDb() *gorm.DB {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("ACCOUNT_DATABASE_USER"),
		os.Getenv("ACCOUNT_DATABASE_PASSWORD"),
		os.Getenv("ACCOUNT_DATABASE_HOST"),
		os.Getenv("ACCOUNT_DATABASE_PORT"),
		os.Getenv("ACCOUNT_DATABASE_NAME"),
	)
	db, err := gorm.Open("mysql", dataSource)

	if err != nil {
		panic(err)
		return nil
	}

	seed.Load(db)

	return db
}
