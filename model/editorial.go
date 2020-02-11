package model

import (
	"db"
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

func (e Editorial) Init() {
	db.Init()
	db := db.GetDb()
	defer db.Close()

	if !db.HasTable("editorials") {
		db.AutoMigrate(&Editorial{})
	}
}

func (e Editorial) Count(ei EditorialInfo) (count int) {
	db.Init()
	db := db.GetDb()
	defer db.Close()

	db.Table("editorials").Where("url = ?", ei.Url).Count(&count)
	return
}

func (e Editorial) Insert(ei EditorialInfo) {
	db.Init()
	db := db.GetDb()
	defer db.Close()

	editorial := Editorial{}
	editorial.Date = ei.Date
	editorial.MediaId = ei.MediaId
	editorial.Title = ei.Title
	editorial.Body = ei.Body
	editorial.Url = ei.Url
	db.Create(&editorial)
}