package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func showIndexPage(context *gin.Context) {
	articles := getAllArticles()

	render(context,
			gin.H{
				"title":   "Home Page",
				"payload": articles,
			},
			"index.html")
}

func getArticle(context *gin.Context) {
	if id, err := strconv.Atoi(context.Param("article_id")); err == nil {
		if article, err := getArticleByID(id); err == nil {
			render(context, gin.H{
				"title":   article.Title,
				"payload": article,
			}, "article.html")
			//context.HTML(http.StatusOK, "article.html", gin.H{
			//	"title":   article.Title,
			//	"payload": article,
			//})
		} else {
			context.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		context.AbortWithStatus(http.StatusNotFound)
	}
}
