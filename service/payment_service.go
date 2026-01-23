package service

import (
	"ecommerce/domain"
	"ecommerce/repository"
)

func ProcessPayment(p domain.Payment) (int, error) {
	return repository.ProcessPayment(p)
}

func GetPaymentById(id int) (domain.Payment, error) {
	return repository.GetPaymentById(id)
}

func GetUserPayments(user_id int) ([]domain.Payment, error) {
	return repository.GetUserPayments(user_id)
}