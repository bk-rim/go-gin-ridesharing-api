package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"api/allo-dakar/utils"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := utils.VerifyJwt(c)
		if err != nil {
			fmt.Println(err)
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
