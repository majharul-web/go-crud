package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/majharul-web/go-crud/initializers"
	"github.com/majharul-web/go-crud/models"
)

func PostCreate(c *gin.Context) {

	post := models.Post{Title: "hello world", Body: "Programming"}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": "post created successfully!",
		"post":    post,
	})
}
