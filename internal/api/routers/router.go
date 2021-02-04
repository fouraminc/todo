package routers

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"todo/internal/api/controllers"
)

func Setup() *gin.Engine {

	app := gin.Default()

	f,_ := os.Create("log/api.log")
	gin.DefaultWriter = io.MultiWriter(f)
	app.Use(gin.Logger())

	app.GET("/todo", controllers.GetTodos)

	return app

}