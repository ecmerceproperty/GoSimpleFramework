package main

import (
	"blog/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		infrastructure.LoadEnv()
		infrastructure.NewDatabase()

		ctx.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
	})
	router.Run(":8000")
}
