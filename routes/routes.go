package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shraddha0602/Go-PostgreSQL/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)

	router.GET("/todos", controllers.GetTodos)
	router.POST("/todo", controllers.CreateTodo)
	router.GET("/todo/:todoId", controllers.GetTodo)
	router.PUT("/todo/:todoId", controllers.EditTodo)
	router.DELETE("/todo/:todoId", controllers.DeleteTodo)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome to API",
	})
	return
}
