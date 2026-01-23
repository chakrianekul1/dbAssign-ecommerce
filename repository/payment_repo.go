package repository

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
	err := db.DB.QueryRow("INSERT INTO payments (order_id, user_id, amount, status) VALUES ($1, $2, $3, $4) RETURNING id", p.OrderID, p.UserID, p.Amount, status).Scan(&id)
	return id, err
}

func GetPaymentById(id int) (domain.Payment, error) {
	var payment domain.Payment
	err := db.DB.QueryRow("SELECT id, order_id, user_id, amount, status FROM payments WHERE id = $1", id).Scan(&payment.ID, &payment.OrderID, &payment.UserID, &payment.Amount, &payment.Status)
	return payment, err
}

func GetUserPayments(user_id int) ([]domain.Payment, error) {
	var payments []domain.Payment
	rows, err := db.DB.Query("SELECT id, order_id, user_id, amount, status FROM payments WHERE user_id = $1", user_id)
	if err != nil {
		return payments, err
	}
	defer rows.Close()
	for rows.Next() {
		var payment domain.Payment
		err = rows.Scan(&payment.ID, &payment.OrderID, &payment.UserID, &payment.Amount, &payment.Status)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}