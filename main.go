/*
Create an CRUD APP First time using Golang,GIN,GORM,Postgres
*/

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


func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.GetAllPost)
	r.GET("/posts/:id", controllers.GetSinglePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong:server is running",
		})
	})
	r.Run()

}

//CompileDaemon -command="./go-crud"
