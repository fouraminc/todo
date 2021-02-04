package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
type simple struct {
	Message string
}
func GetTodos(c *gin.Context)  {
	fmt.Println("In GetTodos")
	c.JSON(http.StatusOK,simple{"Hello World"})
}