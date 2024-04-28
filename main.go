package main

import (
	"github.com/gin-gonic/gin"
	"github.com/majharul-web/go-crud/controllers"
	"github.com/majharul-web/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// L4pbrR4rIRn8mFim
func main() {
	r := gin.Default()

	r.POST("/post", controllers.PostCreate)

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong pong",
		})

	})
	r.Run()

}

//CompileDaemon -command="./go-crud"
