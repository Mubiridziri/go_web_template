package controllers

import (
	"gin-template/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserAction(c *gin.Context) {

	user := models.User{}

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.DB.Save(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func ProfileAction(c *gin.Context) {
	user := c.MustGet("user")

	c.JSON(http.StatusOK, user)
}
