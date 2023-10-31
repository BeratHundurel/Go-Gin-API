package handlers

import (
	"example/data-acces/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserList(c *gin.Context) {
	// Use the DB connection to retrieve user data
	users, err := database.GetUsers()
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the user data as a response
	c.JSON(http.StatusOK, users)
}
