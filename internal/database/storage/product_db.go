package storage

import (
	"example/data-acces/internal/app/models"
	"example/data-acces/internal/database"
	"sync"
	"time"
)

type Cache struct {
	products        []models.Product
	lastUpdate      time.Time
	cacheExpiration time.Duration
	mutex           sync.RWMutex
}

var productCache Cache

func init() {
	productCache = Cache{
		products:        []models.Product{},
		lastUpdate:      time.Time{},
		cacheExpiration: 72 * time.Hour, // Set your desired cache expiration time
	}
}

func GetProducts() ([]models.Product, error) {
	productCache.mutex.RLock()
	if len(productCache.products) > 0 && time.Since(productCache.lastUpdate) < productCache.cacheExpiration {
		defer productCache.mutex.RUnlock()
		return productCache.products, nil
	}
	productCache.mutex.RUnlock()

	var products []models.Product
	if err := database.DB.Select(&products, `SELECT * FROM public.product`); err != nil {
		return nil, err
	}

	productCache.mutex.Lock()
	productCache.products = products
	productCache.lastUpdate = time.Now()
	productCache.mutex.Unlock()

	return products, nil
}

func GetProduct(id int, err error) (models.Product, error) {
	var product models.Product
	database.DB.Get(&product, `SELECT*FROM public."product" WHERE id=$1`, id)
	if err != nil {
		return product, err
	}
	return product, nil
}
