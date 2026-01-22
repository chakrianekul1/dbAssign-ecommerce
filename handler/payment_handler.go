package handler

import (
	"ecommerce/domain"
	"ecommerce/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	var p domain.Payment
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := service.ProcessPayment(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "Payment recorded"})
}

func GetPayment(c *gin.Context) {
	payment_id, _ := strconv.Atoi(c.Param("payment_id"))
	payment, err := service.GetPaymentById(payment_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func GetPayments(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Param("user_id"))
	payments, err := service.GetUserPayments(user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}