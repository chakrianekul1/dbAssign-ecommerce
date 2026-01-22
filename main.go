package main

import (
	"ecommerce/db"
	"ecommerce/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()

	r.POST("/users", handler.AddUser) 
	r.GET("/users/:id", handler.GetOneUser) 
	r.POST("/stores", handler.AddStore) 
	r.GET("/stores", handler.ListStores) 
	r.POST("/stores/:store_id/products", handler.AddProduct)
	r.GET("/products/:id", handler.GetOneProduct) 
	r.POST("/orders", handler.PlaceOrder) 
	r.GET("/orders/:id", handler.GetOneOrder) 
	r.POST("/subscriptions", handler.AddSubscription) 
	r.GET("/users/:id/subscriptions", handler.GetUserSubscriptions) 
	
	r.POST("/payments", handler.CreatePayment)

	r.Run(":8080")
}