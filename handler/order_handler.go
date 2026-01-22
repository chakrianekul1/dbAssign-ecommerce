package handler

import (
	"ecommerce/domain"
	"ecommerce/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
    var o domain.Order
    if err := c.ShouldBindJSON(&o); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    id, err := service.CreateOrder(o)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"id": id})
}

func GetOneOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := service.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func GetUserOrders(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Param("id"))
	orders, err := service.GetOrdersByUserId(user_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "User orders not found"})
		return
	}
	c.JSON(http.StatusOK, orders)
}