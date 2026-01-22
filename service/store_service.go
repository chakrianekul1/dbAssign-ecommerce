package service

import (
	"ecommerce/db"
	"ecommerce/domain"
)

func CreateStore(s domain.Store) (int, error) {
	var id int
	err := db.DB.QueryRow("INSERT INTO stores (name, status) VALUES ($1, $2) RETURNING id", s.Name, "active").Scan(&id)
	return id, err
}

func GetAllStores() ([]domain.Store, error) {
	rows, err := db.DB.Query("SELECT id, name, status FROM stores")
	if err != nil { return nil, err }
	defer rows.Close()
	var stores []domain.Store
	for rows.Next() {
		var s domain.Store
		rows.Scan(&s.ID, &s.Name, &s.Status)
		stores = append(stores, s)
	}
	return stores, nil
}