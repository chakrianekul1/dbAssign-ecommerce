package handler

import (
	"ecommerce/domain"
	"ecommerce/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
    var p domain.Product
    storeID, _ := strconv.Atoi(c.Param("store_id"))
    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    p.StoreID = storeID
    
    id, err := service.CreateProduct(p)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"id": id})
}

func GetOneProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p, err := service.GetProductById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}