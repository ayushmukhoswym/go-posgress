package controllers

import (
	"Users/ayushmukhopadhyay/Desktop/go/PLATFORMS/go-posgress/initializers"
	"Users/ayushmukhopadhyay/Desktop/go/PLATFORMS/go-posgress/models"

	"github.com/gin-gonic/gin"
)

// PostsCreate    godoc
// @Summary      Create Post
// @Description  Creates a Post
// @Tags         Post
// @Accept json
// @Param        post  body  models.Post  true  "Post Data"
// @Produce      application/json
// @Success      200  {array} models.Post
// @Router       /posts [post]
func PostsCreate(c *gin.Context) {
	var requestBody struct {
		Body  string
		Title string
	}
	c.Bind(&requestBody)

	post := models.Post{Title: requestBody.Title, Body: requestBody.Title}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}
