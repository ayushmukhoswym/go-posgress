package main

import (
	"Users/ayushmukhopadhyay/Desktop/go/PLATFORMS/go-posgress/controllers"
	_ "Users/ayushmukhopadhyay/Desktop/go/PLATFORMS/go-posgress/docs"
	"Users/ayushmukhopadhyay/Desktop/go/PLATFORMS/go-posgress/initializers"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

// @title Service API
// @version 1.0
// @description Demo service API

// @host localhost:3000
func main() {
	r := gin.Default()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/posts", controllers.PostsCreate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
