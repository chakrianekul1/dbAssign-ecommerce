package service

import (
	"ecommerce/db"
	"ecommerce/domain"
	"ecommerce/repository"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func GetProductById(id int) (domain.Product, error) {
	start := time.Now()
	var p domain.Product
	cacheKey := fmt.Sprintf("product:%d", id)
	val, err := db.RDB.Get(db.Ctx, cacheKey).Result()
	if err == nil {
		json.Unmarshal([]byte(val), &p)
		log.Printf("[PERF] GetProductById (CACHE HIT) took: %v", time.Since(start))
		return p, nil
	}
	p, err = repository.GetProductById(id)
	if err != nil {
		return p, err
	}
	data, _ := json.Marshal(p)
	db.RDB.Set(db.Ctx, cacheKey, data, 10*time.Minute)
	log.Printf("[PERF] GetProductById (CACHE MISS) took: %v", time.Since(start))
	return p, nil
}

func GetProducts() ([]domain.Product, error) {
	start := time.Now()
	var products []domain.Product
	cacheKey := "products:all"
	val, err := db.RDB.Get(db.Ctx, cacheKey).Result()
	if err == nil {
		json.Unmarshal([]byte(val), &products)
		log.Printf("[PERF] GetProducts (CACHE HIT) took: %v", time.Since(start))
		return products, nil
	}
	products, err = repository.GetProducts()
	if err != nil {
		return nil, err
	}
	data, _ := json.Marshal(products)
	db.RDB.Set(db.Ctx, cacheKey, data, 10*time.Minute)
	log.Printf("[PERF] GetProducts (CACHE MISS) took: %v", time.Since(start))
	return products, nil
}

func CreateProduct(p domain.Product) (int, error) {
	id, err := repository.CreateProduct(p)
	if err == nil {
		db.RDB.Del(db.Ctx, "products:all")
	}
	return id, err
}

func GetProductsByStoreId(store_id int) ([]domain.Product, error) {
	return repository.GetProductsByStoreId(store_id)
}