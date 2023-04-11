package main

import (
	"blog/api/controller"
	"blog/api/repository"
	"blog/api/routes"
	"blog/api/service"
	"blog/infrastructure"
	"blog/models"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {
	// router := gin.Default()
	// router.GET("/", func(ctx *gin.Context) {
	// 	infrastructure.LoadEnv()
	// 	infrastructure.NewDatabase()

	// 	ctx.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
	// })
	// router.Run(":8000")

	router := infrastructure.NewGinRouter()
	db := infrastructure.NewDatabase()
	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)
	postRouter := routes.NewPostRoute(postController, router)
	postRouter.Setup()

	db.DB.AutoMigrate(&models.Post{})
	router.Gin.Run(":8000")
}
