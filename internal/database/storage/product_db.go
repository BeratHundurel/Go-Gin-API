package storage

import (
	"example/data-acces/internal/app/models"
	"example/data-acces/internal/database"
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
)

type Cache struct {
	products []models.Product
	mutex    sync.RWMutex
}

var productCache Cache

func init() {
	productCache = Cache{
		products: []models.Product{},
	}
}

func fetchProducts(rows *sqlx.Rows, productCh chan models.Product, errCh chan error) {
	for rows.Next() {
		var product models.Product
		if err := rows.StructScan(&product); err != nil {
			errCh <- err
			return
		}
		productCh <- product

		productCache.mutex.Lock()
		productCache.products = append(productCache.products, product)
		productCache.mutex.Unlock()
	}
}

func monitorErrors(errCh chan error) {
	for err := range errCh {
		log.Println("Error:", err)
	}
}

func GetProducts() ([]models.Product, error) {
	// Check the cache first
	productCache.mutex.RLock()
	if len(productCache.products) > 0 {
		cachedProducts := productCache.products
		productCache.mutex.RUnlock()
		return cachedProducts, nil
	}
	productCache.mutex.RUnlock()

	rows, err := database.DB.Queryx(`SELECT * FROM public."product"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Properly close the rows

	productCh := make(chan models.Product)
	errCh := make(chan error)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fetchProducts(rows, productCh, errCh)
		defer close(productCh)
	}()

	var products []models.Product
	for product := range productCh {
		products = append(products, product)
	}

	close(errCh)
	go monitorErrors(errCh)
	wg.Wait()

	// Cache the fetched products
	productCache.mutex.Lock()
	productCache.products = products
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
