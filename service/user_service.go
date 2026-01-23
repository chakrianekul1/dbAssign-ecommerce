package service

import (
	"ecommerce/domain"
	"ecommerce/repository"
)

func CreateUser(u domain.User) (int, error) {
	return repository.CreateUser(u)
}

func GetUserByID(id int) (domain.User, error) {
	return repository.GetUserByID(id)
}

func GetUsers() ([]domain.User, error) {
	return repository.GetUsers()
}