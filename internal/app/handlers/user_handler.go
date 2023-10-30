package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/models/internal/app/models"
	"net/http"
)

func GetUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, User)
}
