package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"

	"api/allo-dakar/model"
)

func GetUsers(c *gin.Context) {

	var user model.User

	users, err := user.FindAll()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}
