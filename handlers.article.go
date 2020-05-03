package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func showIndexPage(context *gin.Context) {
	articles := getAllArticles()

	context.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home Page",
		"payload": articles,
	})

}