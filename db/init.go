package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	fmt.Println("DB_USERNAME = " + os.Getenv("EDITORIAL_DB_USERNAME"))
	db, err = gorm.Open("postgres",
		"user="+os.Getenv("EDITORIAL_DB_USERNAME")+" password="+os.Getenv("EDITORIAL_DB_PASSWORD")+" dbname=editorial sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetDb() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
