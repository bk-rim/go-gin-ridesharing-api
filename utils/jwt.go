package utils

import (
	"api/allo-dakar/model"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var privateKey = []byte("secret")

func GenerateJwt(user model.UserLogin) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * 120).Unix(),
	})

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ExtractTokenFromReqHeader(context *gin.Context) string {
	var tokenStr string

	bearerToken := context.Request.Header.Get("Authorization")
	bearerTokenTab := strings.Split(bearerToken, " ")

	if len(bearerTokenTab) == 2 {
		tokenStr = bearerTokenTab[1]
	}
	return tokenStr
}

func getToken(context *gin.Context) (*jwt.Token, error) {

	tokenStr := ExtractTokenFromReqHeader(context)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return privateKey, nil
	})
	return token, err

}

func VerifyJwt(context *gin.Context) (jwt.MapClaims, error) {

	token, err := getToken(context)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
