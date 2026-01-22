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

func GetUsers() ([]domain.User, error) {
	users := []domain.User{}
	rows, err := db.DB.Query("SELECT id, name, email, phone, address FROM users")
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Address)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	defer rows.Close()
	return users, err
}