package domain

import "time"

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Phone   string `json:"phone,omitempty"`
	Address string `json:"address,omitempty"`
}

type Store struct {
	ID     int    `json:"id"`
	Name   string `json:"name" binding:"required"`
	Status string `json:"status"`
}

type Product struct {
	ID           int       `json:"id"`
	StoreID      int       `json:"store_id"`
	Name         string    `json:"name" binding:"required"`
	Price        float64   `json:"price" binding:"required"`
	Availability bool      `json:"availability"`
	CreatedAt    time.Time `json:"created_at"`
}

type OrderItem struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

type Order struct {
	ID        int         `json:"id"`
	UserID    int         `json:"user_id" binding:"required"`
	Items     []OrderItem `json:"items" binding:"required"` 
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}

type Payment struct {
	ID          int       `json:"id"`
	OrderID     int       `json:"order_id" binding:"required"`
	UserID 		int 	  `json:"user_id" binding:"required"`
	Amount      float64   `json:"amount" binding:"required"`
	Status      string    `json:"status"`
	PaymentDate time.Time `json:"payment_date"`
}

type Subscription struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id" binding:"required"`
	ProductID int        `json:"product_id" binding:"required"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date,omitempty"`
	Status    string     `json:"status"`
}