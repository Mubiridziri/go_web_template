package controllers

import (
	"database/sql"
	"gin-template/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const (
	UserKey = "AUTH"
)

func LoginAction(c *gin.Context) {
	session := sessions.Default(c)

	loginForm := models.UserLogin{}

	if err := c.BindJSON(&loginForm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	db := models.DB

	var user models.User

	db.Where("username = @username", sql.Named("username", loginForm.Username)).First(&user)

	if pErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginForm.Password)); pErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	session.Set(UserKey, loginForm.Username)
	err := session.Save()
	if err != nil {
		return
	}

	c.JSON(http.StatusAccepted, loginForm)

}

func LogoutAction(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(UserKey)
	err := session.Save()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Error save session!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
