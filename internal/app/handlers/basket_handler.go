package handlers

import (
	"example/data-acces/internal/database/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddToBasket(c *gin.Context) {
    productID, err := strconv.Atoi(c.Param("productID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid productID"})
        return
    }

    quantity, err := strconv.Atoi(c.Param("quantity"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
        return
    }
	basket_item, err := storage.CreateBasketItem(c, productID, quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, basket_item)
	return
}
