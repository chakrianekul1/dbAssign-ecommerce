package service

import (
	"ecommerce/db"
	"ecommerce/domain"
)

func CreateSubscription(s domain.Subscription) (int, error) {
	var id int
	err := db.DB.QueryRow("INSERT INTO subscriptions (user_id, product_id, status) VALUES ($1, $2, $3) RETURNING id", s.UserID, s.ProductID, "active").Scan(&id)
	return id, err
}

func GetSubscriptionsByUserID(userID int) ([]domain.Subscription, error) {
	rows, err := db.DB.Query("SELECT id, user_id, product_id, start_date, end_date, status FROM subscriptions WHERE user_id = $1", userID)
	if err != nil { return nil, err }
	defer rows.Close()
	var subs []domain.Subscription
	for rows.Next() {
		var s domain.Subscription
		err = rows.Scan(&s.ID, &s.UserID, &s.ProductID, &s.StartDate, &s.EndDate, &s.Status)
		if err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}
	return subs, nil
}