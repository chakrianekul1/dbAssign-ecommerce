package service

import (
	"ecommerce/db"
	"ecommerce/domain"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func GetProductById(id int) (domain.Product, error) {
	start := time.Now() // Start high-precision timer
	var p domain.Product
	cacheKey := fmt.Sprintf("product:%d", id)

	val, err := db.RDB.Get(db.Ctx, cacheKey).Result()
	if err == nil {
		json.Unmarshal([]byte(val), &p)
		log.Printf("[PERF] GetProductById (CACHE HIT) took: %v", time.Since(start))
		return p, nil
	}

	err = db.DB.QueryRow("SELECT id, store_id, name, price, availability, created_at FROM products WHERE id = $1", id).
		Scan(&p.ID, &p.StoreID, &p.Name, &p.Price, &p.Availability, &p.CreatedAt)
	
	if err != nil {
		return p, err
	}

	data, _ := json.Marshal(p)
	db.RDB.Set(db.Ctx, cacheKey, data, 10 * time.Minute)

	log.Printf("[PERF] GetProductById (CACHE MISS) took: %v", time.Since(start))
	return p, nil
}

func GetProductsByStoreId(store_id int) ([]domain.Product, error) {
	var products []domain.Product
	rows, err := db.DB.Query("SELECT id, store_id, name, price, availability, created_at FROM products WHERE store_id = $1", store_id)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.Product
		err = rows.Scan(&product.ID, &product.StoreID, &product.Name, &product.Price, &product.Availability, &product.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, err
}

func CreateProduct(p domain.Product) (int, error) {
	var id int
	err := db.DB.QueryRow("INSERT INTO products (store_id, name, price, availability) VALUES ($1, $2, $3, $4) RETURNING id", p.StoreID, p.Name, p.Price, p.Availability).Scan(&id)
	if err == nil {
		db.RDB.Del(db.Ctx, "products:all")
	}
	return id, err
}

func GetProducts() ([]domain.Product, error) {
	start := time.Now() // Start high-precision timer
	var products []domain.Product
	cacheKey := "products:all"

	val, err := db.RDB.Get(db.Ctx, cacheKey).Result()
	if err == nil {
		json.Unmarshal([]byte(val), &products)
		log.Printf("[PERF] GetProducts (CACHE HIT) took: %v", time.Since(start))
		return products, nil
	}

	rows, err := db.DB.Query("SELECT id, store_id, name, price, availability, created_at FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p domain.Product
		rows.Scan(&p.ID, &p.StoreID, &p.Name, &p.Price, &p.Availability, &p.CreatedAt)
		products = append(products, p)
	}

	data, _ := json.Marshal(products)
	db.RDB.Set(db.Ctx, cacheKey, data, 10 * time.Minute)
	
	log.Printf("[PERF] GetProducts (CACHE MISS) took: %v", time.Since(start))
	return products, nil
}