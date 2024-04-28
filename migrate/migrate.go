package main

import (
	"github.com/majharul-web/go-crud/initializers"
	"github.com/majharul-web/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
