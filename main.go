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
    r.GET("/users", handler.GetAllUsers)
    r.GET("/users/:id", handler.GetOneUser) 

    r.POST("/stores", handler.AddStore) 
    r.GET("/stores", handler.ListStores) 
    r.GET("/stores/id/:store_id", handler.GetOneStore) 

    r.POST("/stores/:store_id/products", handler.AddProduct)
    r.GET("/products/id/:id", handler.GetOneProduct) 
    r.GET("/products", handler.GetAllProducts)
    r.GET("/stores/:store_id/products", handler.GetStoreProducts)

    r.POST("/orders", handler.PlaceOrder) 
    r.GET("/orders/id/:id", handler.GetOneOrder) 
    r.GET("/users/:id/orders", handler.GetUserOrders)

    r.POST("/subscriptions", handler.AddSubscription) 
    r.GET("/users/:id/subscriptions", handler.GetUserSubscriptions) 
    
    r.POST("/payments", handler.CreatePayment)
    r.GET("/payments/id/:payment_id", handler.GetPayment) 
    r.GET("/users/:id/payments", handler.GetPayments)

	r.Run(":8000")
}