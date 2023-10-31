package routes

import (
	"example/data-acces/internal/app/handlers"
	"github.com/gin-gonic/gin"
)

func SetRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/users", handlers.GetUserList)
	router.GET("/users/:id", handlers.GetUser)
	router.GET("/products", handlers.GetProductsList)
	router.GET("/products/:id", handlers.GetProduct)
	return router
}
