package routes

import (
	"example/data-acces/internal/app/handlers"
	"github.com/gin-gonic/gin"
)

func SetRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/users", handlers.GetUserList)
	return router
}
