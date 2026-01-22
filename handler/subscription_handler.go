package handler

import (
	"ecommerce/domain"
	"ecommerce/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func AddSubscription(c *gin.Context) {
	var s domain.Subscription
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := service.CreateSubscription(s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return 
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func GetUserSubscriptions(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	subs, err := service.GetSubscriptionsByUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subs)
}