package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"api/allo-dakar/controller"
	"api/allo-dakar/database"
	"api/allo-dakar/utils"
)

func TestGetUsers(t *testing.T) {

	utils.LoadEnv("../.env")
	database.Connect()
	router := gin.Default()
	router.GET("/users", controller.GetUsers)

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}
