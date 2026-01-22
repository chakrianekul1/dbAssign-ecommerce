package service

import (
	"ecommerce/db"
	"ecommerce/domain"
)

func CreateUser(u domain.User) (int, error) {
	var id int
	err := db.DB.QueryRow("INSERT INTO users (name, email, phone, address) VALUES ($1, $2, $3, $4) RETURNING id", u.Name, u.Email, u.Phone, u.Address).Scan(&id)
	return id, err
}

func GetUserByID(id int) (domain.User, error) {
	var u domain.User
	err := db.DB.QueryRow("SELECT id, name, email, phone, address FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.Address)
	return u, err
}