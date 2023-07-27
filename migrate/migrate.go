package main

import (
	"Users/ayushmukhopadhyay/Desktop/go/PLATFORMS/go-posgress/initializers"
	"Users/ayushmukhopadhyay/Desktop/go/PLATFORMS/go-posgress/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
