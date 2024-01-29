package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"api/allo-dakar/controller"
	"api/allo-dakar/database"
	"api/allo-dakar/model"
	"api/allo-dakar/utils"
)

func TestCreateUser(t *testing.T) {

	utils.LoadEnv("../.env")
	database.Connect()
	router := gin.Default()
	router.POST("/users", controller.CreateUser)

	user := model.User{
		Username: "testuser",
		Email:    "test244@example.com",
		Password: "testpassword",
	}
	payload, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	if w.Code == http.StatusCreated {
		_, err := database.DB.Query(`DELETE FROM "users" WHERE email = $1`, user.Email)
		if err != nil {
			t.Fatal(err)
		}
	}

}

func TestCreateUserWithBadRequestBody(t *testing.T) {

	utils.LoadEnv("../.env")
	database.Connect()
	router := gin.Default()
	router.POST("/users", controller.CreateUser)

	user := "testuserbadRequest"
	payload, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAuthenticateUser(t *testing.T) {

	utils.LoadEnv("../.env")
	database.Connect()
	router := gin.Default()
	router.POST("/login", controller.AuthenticateUser)

	userLogin := model.UserLogin{
		Email:    "test3@example.com",
		Password: "testpassword",
	}
	payloadLogin, _ := json.Marshal(userLogin)

	reqLogin, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payloadLogin))
	if err != nil {
		t.Fatal(err)
	}
	reqLogin.Header.Set("Content-Type", "application/json")

	wLogin := httptest.NewRecorder()

	router.ServeHTTP(wLogin, reqLogin)

	assert.Equal(t, http.StatusOK, wLogin.Code)
}

func TestAuthenticateUserWithBadRequestBody(t *testing.T) {

	utils.LoadEnv("../.env")
	database.Connect()
	router := gin.Default()
	router.POST("/login", controller.AuthenticateUser)

	userLogin := "testuserbadRequest"
	payloadLogin, _ := json.Marshal(userLogin)

	reqLogin, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payloadLogin))
	if err != nil {
		t.Fatal(err)
	}
	reqLogin.Header.Set("Content-Type", "application/json")

	wLogin := httptest.NewRecorder()

	router.ServeHTTP(wLogin, reqLogin)

	assert.Equal(t, http.StatusBadRequest, wLogin.Code)
}

func TestAuthenticateUserWithWrongPassword(t *testing.T) {

	utils.LoadEnv("../.env")
	database.Connect()
	router := gin.Default()
	router.POST("/login", controller.AuthenticateUser)

	userLogin := model.UserLogin{
		Email:    "test3@example.com",
		Password: "falsepassword",
	}

	payloadLogin, _ := json.Marshal(userLogin)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payloadLogin))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthenticateUserWithWrongEmail(t *testing.T) {

	utils.LoadEnv("../.env")
	database.Connect()
	router := gin.Default()
	router.POST("/login", controller.AuthenticateUser)

	userLogin := model.UserLogin{
		Email:    "test3@badmail.com",
		Password: "testpassword",
	}

	payloadLogin, _ := json.Marshal(userLogin)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payloadLogin))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
