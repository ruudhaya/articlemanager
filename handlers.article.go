package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func showIndexPage(context *gin.Context) {
	articles := getAllArticles()

	context.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Home Page",
		"payload": articles,
	})
}

func getArticle(context *gin.Context) {
	if id, err := strconv.Atoi(context.Param("article_id")); err == nil {
		if article, err := getArticleByID(id); err == nil {
			context.HTML(http.StatusOK, "article.html", gin.H{
				"title":   article.Title,
				"payload": article,
			})
		} else {
			context.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		context.AbortWithStatus(http.StatusNotFound)
	}
}
