package service

import (
	"ecommerce/db"
	"ecommerce/domain"
	"ecommerce/repository"
	"encoding/json"
	"log"
	"time"
)

func CreateStore(s domain.Store) (int, error) {
	id, err := repository.CreateStore(s)
	if err == nil {
		db.RDB.Del(db.Ctx, "stores:all")
	}
	return id, err
}

func GetAllStores() ([]domain.Store, error) {
	start := time.Now()
	var stores []domain.Store
	cacheKey := "stores:all"
	val, err := db.RDB.Get(db.Ctx, cacheKey).Result()
	if err == nil {
		json.Unmarshal([]byte(val), &stores)
		log.Printf("[PERF] GetAllStores (CACHE HIT) took: %v", time.Since(start))
		return stores, nil
	}
	stores, err = repository.GetAllStores()
	if err != nil {
		return nil, err
	}
	data, _ := json.Marshal(stores)
	db.RDB.Set(db.Ctx, cacheKey, data, 30*time.Minute)
	log.Printf("[PERF] GetAllStores (CACHE MISS) took: %v", time.Since(start))
	return stores, nil
}

func GetStoreById(id int) (domain.Store, error) {
	return repository.GetStoreById(id)
}