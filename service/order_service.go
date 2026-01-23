package service

import (
	"ecommerce/domain"
	"ecommerce/repository"
)

func CreateOrder(o domain.Order) (int, error) {
	return repository.CreateOrder(o)
}

func GetOrderByID(id int) (domain.Order, error) {
	o, err := repository.GetOrderOnly(id)
	if err != nil {
		return o, err
	}
	items, err := repository.GetOrderItems(id)
	if err != nil {
		return o, err
	}
	o.Items = items
	return o, nil
}

func GetOrdersByUserId(id int) ([]domain.Order, error) {
	orders, err := repository.GetOrdersByUserId(id)
	if err != nil {
		return nil, err
	}
	for i := range orders {
		items, err := repository.GetOrderItems(orders[i].ID)
		if err != nil {
			return nil, err
		}
		orders[i].Items = items
	}
	return orders, nil
}