package repository

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

func GetOrderOnly(id int) (domain.Order, error) {
	var o domain.Order
	query := `SELECT id, user_id, status, created_at FROM orders WHERE id = $1`
	err := db.DB.QueryRow(query, id).Scan(&o.ID, &o.UserID, &o.Status, &o.CreatedAt)
	return o, err
}

func GetOrderItems(orderID int) ([]domain.OrderItem, error) {
	rows, err := db.DB.Query("SELECT product_id, quantity FROM order_items WHERE order_id = $1", orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.OrderItem
	for rows.Next() {
		var item domain.OrderItem
		err := rows.Scan(&item.ProductID, &item.Quantity)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func GetOrdersByUserId(user_id int) ([]domain.Order, error) {
	rows, err := db.DB.Query("SELECT id, user_id, status, created_at FROM orders WHERE user_id = $1", user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []domain.Order
	for rows.Next() {
		var o domain.Order
		err = rows.Scan(&o.ID, &o.UserID, &o.Status, &o.CreatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}