package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/majharul-web/go-crud/initializers"
	"github.com/majharul-web/go-crud/models"
)

func PostCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

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

func GetAllPost(c *gin.Context) {

	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"message": "posts get successfully!",
		"post":    posts,
	})

}

func GetSinglePost(c *gin.Context) {

	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"message": "posts get successfully!",
		"post":    post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"message": "posts updated successfully!",
		"post":    post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"message": "posts deleted successfully!",
		"post":    post,
	})

}
