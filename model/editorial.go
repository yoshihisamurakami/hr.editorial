package model

import (
	"db"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Editorial struct {
	gorm.Model
	Url     string
	Date    string
	MediaId int
	Title   string
	Body    string
}

func (e Editorial) Count(ei EditorialInfo) int {
	return db.EditorialsCount(ei.Url)
}

func (e Editorial) Insert(ei EditorialInfo) {
	db.Init()
	db := db.GetDb()
	fmt.Printf("[MDLins] address = %p\n", db)
	defer db.Close()

	editorial := Editorial{}
	editorial.Date = ei.Date
	editorial.MediaId = ei.MediaId
	editorial.Title = ei.Title
	editorial.Body = ei.Body
	editorial.Url = ei.Url
	db.Create(&editorial)
}
