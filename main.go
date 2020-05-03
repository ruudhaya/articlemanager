// main.go

package main

import (
	"github.com/gin-gonic/gin"
	_ "net/http"
)

var router *gin.Engine

func main() {

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	initializeRoutes()

	router.Run()
}
