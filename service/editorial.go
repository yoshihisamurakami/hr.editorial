package service

import (
	"db"
	"fmt"
	"model"
	"strings"

	"github.com/jinzhu/gorm"
)

var dbInstance *gorm.DB

func init() {
	fmt.Println("### service editorial Init!")
}

func GetEditorials() []model.EditorialView {
	// db.Init()
	// dbInstance := db.GetDb()
	dbInit()
	defer db.Close()

	var article []model.Editorial
	var articleView []model.EditorialView
	dbInstance.Limit(30).Order("date DESC, media_id DESC, created_at DESC").Find(&article)

	for _, v := range article {
		av := model.EditorialView{}
		av.Init(v)
		articleView = append(articleView, av)
	}
	return articleView
}

func GetOneEditorial(id string) model.EditorialView {
	// db.Init()
	// dbInstance := db.GetDb()
	// defer db.Close()
	dbInit()

	var article model.Editorial
	var articleView model.EditorialView
	dbInstance.Find(&article, id)
	articleView.Init(article)
	articleView.FormattedBody = getArticleBody(article.Body)
	return articleView
}

func GetPrevNextEditorial(id string) (prevArticle model.Editorial, nextArticle model.Editorial) {
	// db.Init()
	// dbInstance := db.GetDb()
	defer db.Close()

	var article []model.Editorial
	var count int
	dbInstance.Order("date DESC, media_id DESC, created_at DESC").Find(&article).Count(&count)
	max := count

	for i, _ := range article {
		if fmt.Sprint(article[i].ID) == id {
			if i == 0 {
				nextArticle = article[i+1]
			} else if i+1 == max {
				prevArticle = article[i-1]
			} else {
				prevArticle = article[i-1]
				nextArticle = article[i+1]
			}
			return
		}
	}
	return
}

func dbInit() {
	db.Init()
	dbInstance = db.GetDb()
}

func getArticleBody(body string) []string {
	return strings.Split(body, "\n")
}
