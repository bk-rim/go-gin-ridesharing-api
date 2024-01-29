package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"

	"api/allo-dakar/model"

	"api/allo-dakar/database"
)

func GetUsers(c *gin.Context) {

	rows, err := database.DB.Query(`SELECT * FROM "users"`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var userTemp model.User
		err := rows.Scan(&userTemp.Id, &userTemp.Username, &userTemp.Email, &userTemp.Password)
		if err != nil {
			panic(err)
		}
		users = append(users, userTemp)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, users)
}
