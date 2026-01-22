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
    if err != nil {return o, err}
	rows, err := db.DB.Query("SELECT product_id, quantity FROM order_items WHERE order_id = $1", id)
	if err != nil {return o, err}
	var order_items []domain.OrderItem
	for rows.Next() {
		var order_item domain.OrderItem
		err := rows.Scan(&order_item.ProductID, &order_item.Quantity)
		if err != nil {return o, err}
		order_items = append(order_items, order_item)
	}
	defer rows.Close()
	o.Items = order_items
	return o, nil
}

func GetOrdersByUserId(id int) ([]domain.Order, error) {
	var orders []domain.Order
	rows, err := db.DB.Query("SELECT id, user_id, status, created_at FROM orders WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order domain.Order
		err = rows.Scan(&order.ID, &order.UserID, &order.Status, &order.CreatedAt)
		if err != nil {
			return nil, err
		}
		itemRows, err := db.DB.Query("SELECT product_id, quantity FROM order_items WHERE order_id = $1", order.ID)
		if err != nil {
			return nil, err
		}
		var items []domain.OrderItem
		for itemRows.Next() {
			var item domain.OrderItem
			if err := itemRows.Scan(&item.ProductID, &item.Quantity); err != nil {
				itemRows.Close()
				return nil, err
			}
			items = append(items, item)
		}
		itemRows.Close() 
		order.Items = items
		orders = append(orders, order)
	}
	return orders, nil
}