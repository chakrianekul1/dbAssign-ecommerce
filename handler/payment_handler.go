package handler

import (
	"ecommerce/domain"
	"ecommerce/service"
	"net/http"
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