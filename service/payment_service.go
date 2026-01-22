package service

import (
	"ecommerce/db"
	"ecommerce/domain"
)

func ProcessPayment(p domain.Payment) (int, error) {
	var id int
	status := p.Status
	if status == "" {
		status = "pending"
	}

	err := db.DB.QueryRow("INSERT INTO payments (order_id, amount, status) VALUES ($1, $2, $3) RETURNING id", p.OrderID, p.Amount, status).Scan(&id)
	return id, err
}