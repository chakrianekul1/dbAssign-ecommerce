package service

import (
	"ecommerce/db"
	"ecommerce/domain"
)

func CreateOrder(o domain.Order) (int, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return 0, err
	}

	var orderID int
	err = tx.QueryRow("INSERT INTO orders (user_id, status) VALUES ($1, $2) RETURNING id", o.UserID, "created").Scan(&orderID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, item := range o.Items {
		_, err = tx.Exec("INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)", orderID, item.ProductID, item.Quantity)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	err = tx.Commit()
	return orderID, err
}

func GetOrderByID(id int) (domain.Order, error) {
    var o domain.Order
    query := `SELECT id, user_id, status, created_at FROM orders WHERE id = $1`
    err := db.DB.QueryRow(query, id).Scan(&o.ID, &o.UserID, &o.Status, &o.CreatedAt)
    return o, err
}