package controller

import (
	"service"

	"github.com/gin-gonic/gin"
)

// Index ...
func Index(ctx *gin.Context) {
	articles := service.GetEditorials()
	ctx.HTML(200, "index.html", gin.H{"articles": articles})
}

// Show ...
func Show(ctx *gin.Context) {
	id := ctx.Param("id")
	article := service.GetOneEditorial(id)
	prevArticle, nextArticle := service.GetPrevNextEditorial(id)
	ctx.HTML(200, "article.html", gin.H{"article": article, "prevArticle": prevArticle, "nextArticle": nextArticle})
}
