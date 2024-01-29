package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"

	"api/allo-dakar/model"
	"api/allo-dakar/utils"
)

func CreateUser(c *gin.Context) {

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := user.BeforeSave()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode password"})
		return
	}

	savedUser, err := user.Save()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}

	c.IndentedJSON(http.StatusCreated, savedUser)
}

func AuthenticateUser(c *gin.Context) {

	var user model.UserLogin
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userGetted, errQuery := user.FindByEmail()
	if errQuery != nil {
		if errQuery == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		}
		return
	}

	errCompare := userGetted.ComparePassword(user.Password)
	if errCompare != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	jwt, err := utils.GenerateJwt(user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate jwt"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Login successful", "jwt": jwt, "username": userGetted.Username, "email": userGetted.Email})
	return
}
