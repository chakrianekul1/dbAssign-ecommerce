package handler

import (
	"ecommerce/domain"
	"ecommerce/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var u domain.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := service.CreateUser(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func GetOneUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
	users, err := service.GetUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error" : "error getting users"})
		return
	}
	c.JSON(http.StatusOK, users)
}