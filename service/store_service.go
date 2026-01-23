package service

import (
	"ecommerce/db"
	"ecommerce/domain"
	"encoding/json"
	"log"
	"time"
)

func CreateStore(s domain.Store) (int, error) {
	var id int
	err := db.DB.QueryRow("INSERT INTO stores (name, status) VALUES ($1, $2) RETURNING id", s.Name, "active").Scan(&id)
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

	rows, err := db.DB.Query("SELECT id, name, status FROM stores")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s domain.Store
		rows.Scan(&s.ID, &s.Name, &s.Status)
		stores = append(stores, s)
	}

	data, _ := json.Marshal(stores)
	db.RDB.Set(db.Ctx, cacheKey, data, 30 * time.Minute)

	log.Printf("[PERF] GetAllStores (CACHE MISS) took: %v", time.Since(start))
	return stores, nil
}

func GetStoreById(id int) (domain.Store, error) {
	var store domain.Store
	err := db.DB.QueryRow("SELECT id, name, status FROM stores WHERE id = $1", id).Scan(&store.ID, &store.Name, &store.Status)
	return store, err
}