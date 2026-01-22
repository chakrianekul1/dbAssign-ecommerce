package handler

import (
	"ecommerce/domain"
	"ecommerce/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddStore(c *gin.Context) {
	var s domain.Store
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := service.CreateStore(s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "Store created successfully"})
}

func ListStores(c *gin.Context) {
	stores, err := service.GetAllStores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stores)
}