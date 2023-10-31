package handlers

import (
	"example/data-acces/internal/database/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProductsList(c *gin.Context) {
	products, err := storage.GetProducts()
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// Return the user data as a response
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	// Get the "id" parameter from the URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	product, err := storage.GetProduct(id, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}
