package storage

import (
	"database/sql"
	"example/data-acces/internal/app/models"
	"example/data-acces/internal/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetBasket retrieves the user's basket from the database.
// If the basket doesn't exist, it creates one and then retrieves it.
func GetBasket(userID int, c *gin.Context) (models.Basket, error) {
	var basket models.Basket

	rows, err := database.DB.Queryx(`SELECT id, user_id FROM public.basket WHERE user_id = $1`, userID)
	if err != nil {
		return models.Basket{}, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&basket.Id, &basket.UserId); err != nil {
			return models.Basket{}, err
		}
	}
	if err == sql.ErrNoRows {
		basket, err = createBasket(c)
		if err != nil {
			return models.Basket{}, err
		}
	}
	return basket, nil
}

func GetBasketItem(basketID int, productID int, c *gin.Context) ([]models.BasketItem, error) {
	var basketItems []models.BasketItem
	rows, err := database.DB.Queryx(`SELECT id, basket_id, product_id, quantity FROM public.basket_item WHERE basket_id = $1 AND product_id = $2`, basketID, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var basketItem models.BasketItem
		if err := rows.Scan(&basketItem.Id, &basketItem.BasketId, &basketItem.ProductId, &basketItem.Quantity); err != nil {
			return nil, err
		}

		basket, product, err := getUserBasketAndProduct(c, productID)
		if err != nil {
			return nil, err
		}

		basketItem.Basket = basket
		basketItem.Product = product
		basketItems = append(basketItems, basketItem)
	}

	return basketItems, nil
}

// createBasket creates a new basket for the given user and returns it.
func createBasket(c *gin.Context) (models.Basket, error) {
	// Create a cookie for the user and set it in the response.
	// Create a new cookie.
	userID, err := c.Cookie("user_id")
	if userID == "" || err != nil {
		userID = "1"
		c.SetCookie("user_id", userID, 3600*24*30, "/", "", false, true)
	}

	user_id, err := strconv.Atoi(userID)
	if err != nil {
		return models.Basket{}, err
	}
	basket := models.Basket{
		UserId: user_id,
	}

	_, err = database.DB.Exec(`INSERT INTO public.basket (user_id) VALUES ($1)`, basket.UserId)
	if err != nil {
		return models.Basket{}, err
	}

	return basket, nil
}

func CreateBasketItem(c *gin.Context, productID int, quantity int) ([]models.BasketItem, error) {
	// Get the user ID from the cookie.
	basket, product, err := getUserBasketAndProduct(c, productID)
	if err != nil {
		return nil, err
	}
	basket_item, err := GetBasketItem(basket.Id, product.Id, c)
	if len(basket_item) > 0 {
		// The product is already in the basket, so update the quantity.
		newQuantity := quantity + basket_item[0].Quantity
		_, err = database.DB.Exec(`UPDATE public.basket_item SET quantity = $1 WHERE basket_id = $2 AND product_id = $3`, newQuantity, basket.Id, product.Id)
		if err != nil {
			fmt.Printf("error in updating the basket item")
			return nil, err
		}
		return basket_item, nil
	}
	if err != nil {
		return nil, err
	}
	newItem := models.BasketItem{
		ProductId: productID,
		Quantity:  quantity,
		Product:   product,
		BasketId:  basket.Id,
		Basket:    basket,
	}
	_, err = database.DB.Exec(`INSERT INTO public.basket_item (product_id, quantity, basket_id) VALUES ($1, $2, $3)`, newItem.ProductId, newItem.Quantity, newItem.BasketId)
	if err != nil {
		fmt.Printf("error in db exc")
		return []models.BasketItem{}, err
	}
	return []models.BasketItem{newItem}, nil
}

func getUserBasketAndProduct(c *gin.Context, productID int) (models.Basket, models.Product, error) {
	// Get the user ID from the cookie.
	cookie, err := c.Request.Cookie("user_id")
	if err != nil {
		createBasket(c)
	}
	userID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		return models.Basket{}, models.Product{}, err
	}

	basket, err := GetBasket(userID, c)
	if err != nil {
		return models.Basket{}, models.Product{}, err
	}

	product, err := GetProduct(productID, err)
	if err != nil {
		return models.Basket{}, models.Product{}, err
	}

	return basket, product, nil
}
