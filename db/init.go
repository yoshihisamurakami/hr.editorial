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
	// 接続用
	url := os.Getenv("DATABASE_URL")
	connection, err := pq.ParseURL(url)
	if err != nil {
		panic(err.Error())
	}
	//connection += " sslmode=require"
	connection += " sslmode=disable"
	gorm.Open("postgres", connection)
	db, err = gorm.Open("postgres", connection)
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
