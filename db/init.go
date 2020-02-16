package db

import (
	"fmt"
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
	url := os.Getenv("DATABASE_URL")
	connection, err := pq.ParseURL(url)
	if err != nil {
		panic(err.Error())
	}
	connection += " sslmode=disable"
	gorm.Open("postgres", connection)
	db, err = gorm.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[DBinit] address = %p\n", db)
}

// GetDb ..
func GetDb() *gorm.DB {
	return db
}

func EditorialsCount(url string) (count int) {
	Init()
	defer Close()

	db.Table("editorials").Where("url = ?", url).Count(&count)
	return
}

// Close ..
func Close() {
	fmt.Println("db Close..")
	if err := db.Close(); err != nil {
		panic(err)
	}
}
