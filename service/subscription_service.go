package service

import (
	"ecommerce/domain"
	"ecommerce/repository"
)

func CreateSubscription(s domain.Subscription) (int, error) {
	return repository.CreateSubscription(s)
}

func GetSubscriptionsByUserID(userID int) ([]domain.Subscription, error) {
	return repository.GetSubscriptionsByUserID(userID)
}