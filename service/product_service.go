package service

import (
	"ecommerce/db"
	"ecommerce/domain"
)

func CreateProduct(p domain.Product) (int, error) {
	var id int
	err := db.DB.QueryRow("INSERT INTO products (store_id, name, price, availability) VALUES ($1, $2, $3, $4) RETURNING id", p.StoreID, p.Name, p.Price, p.Availability).Scan(&id)
	return id, err
}

func GetProductById(id int) (domain.Product, error) {
	var p domain.Product
	err := db.DB.QueryRow("SELECT id, store_id, name, price, availability, created_at FROM products WHERE id = $1", id).Scan(&p.ID, &p.StoreID, &p.Name, &p.Price, &p.Availability, &p.CreatedAt)
	return p, err
}