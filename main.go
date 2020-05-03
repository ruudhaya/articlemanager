// main.go

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http"
)

var router *gin.Engine

func main() {

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	initializeRoutes()

	router.Run()
}

func initializeRoutes() {
	router.GET("/", showIndexPage)
}

func showIndexPage(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
		},
	)
}