package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

var (
	_db *gorm.DB
	err error
)

// Init ..
func Init() {
	url := os.Getenv("DATABASE_URL")
	connection, err := pq.ParseURL(url)
	fmt.Println("### connection URL = " + url)
	if err != nil {
		panic(err.Error())
	}
	connection += " sslmode=disable"
	_db, err = gorm.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
}

// GetDb ..
func GetDb() *gorm.DB {
	return _db
}

func EditorialsCount(url string) (count int) {
	Init()
	defer Close()

	_db.Table("editorials").Where("url = ?", url).Count(&count)
	return
}

// Close ..
func Close() {
	fmt.Println("### connection Close..")
	if err := _db.Close(); err != nil {
		panic(err)
	}
}
