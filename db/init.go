package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

var (
	db  *gorm.DB
	err error
)

// Init ..
func Init() {
	//fmt.Println("DB_USERNAME = " + os.Getenv("EDITORIAL_DB_USERNAME"))
	// db, err = gorm.Open("postgres",
	// 	"user="+os.Getenv("EDITORIAL_DB_USERNAME")+" password="+os.Getenv("EDITORIAL_DB_PASSWORD")+" dbname=editorial sslmode=disable")

	// heroku接続用
	url := os.Getenv("DATABASE_URL")
	connection, err := pq.ParseURL(url)
	if err != nil {
		panic(err.Error())
	}
	connection += " sslmode=require"
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
}

// GetDb ..
func GetDb() *gorm.DB {
	return db
}

// Close ..
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
