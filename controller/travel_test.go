package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"api/allo-dakar/controller"
	"api/allo-dakar/database"
	"api/allo-dakar/model"
	"api/allo-dakar/utils"
)

var idTravel int64

func TestCreateTravel(t *testing.T) {
	utils.LoadEnv("../.env")
	database.Connect()
	router := gin.Default()
	router.POST("/travels", controller.CreateTravel)

	startTime, err := time.Parse("15:04", "10:00")
	if err != nil {
		t.Fatal(err)
	}

	endTime, err := time.Parse("15:04", "12:00")
	if err != nil {
		t.Fatal(err)
	}

	travel := model.TravelPost{
		DriverId:  1,
		StartArea: "StartArea",
		EndArea:   "EndArea",
		Price:     10,
		Date:      "2022-01-01",
		StartTime: startTime,
		EndTime:   endTime,
		Places:    4,
		Phone:     "1234567890",
		Comment:   "Test travel",
	}
	payload, _ := json.Marshal(travel)

	req, err := http.NewRequest("POST", "/travels", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response model.Travel
	responseMap := map[string]interface{}{}
	err = json.Unmarshal(w.Body.Bytes(), &responseMap)
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	idTravelFloat := responseMap["id"].(float64)
	idTravel = int64(idTravelFloat)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateTravelWithBadRequestBody(t *testing.T) {
	utils.LoadEnv("../.env")
	database.Connect()
	router := gin.Default()
	router.POST("/travels", controller.CreateTravel)

	travel := "testtravelbadRequest"
	payload, _ := json.Marshal(travel)

	req, err := http.NewRequest("POST", "/travels", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetTravels(t *testing.T) {
	utils.LoadEnv("../.env")
	database.Connect()
	router := gin.Default()
	router.GET("/travels", controller.GetTravels)

	req, err := http.NewRequest("GET", "/travels", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTravelsByStartAndEnd(t *testing.T) {
	utils.LoadEnv("../.env")
	database.Connect()
	router := gin.Default()
	router.GET("/travels/:start/:end", controller.GetTravelsByStartAndEnd)

	req, err := http.NewRequest("GET", "/travels/StartArea/EndArea", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTravelsByStartAndEndAndDate(t *testing.T) {
	utils.LoadEnv("../.env")
	database.Connect()
	router := gin.Default()
	router.GET("/travels/:start/:end/:date", controller.GetTravelsByStartAndEndAndDate)

	req, err := http.NewRequest("GET", "/travels/Start Area/End Area/2022-01-01", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTravelsByIdDriver(t *testing.T) {
	utils.LoadEnv("../.env")
	database.Connect()
	router := gin.Default()
	router.GET("/travels/:idDriver", controller.GetTravelsByIdDriver)

	req, err := http.NewRequest("GET", "/travels/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTravel(t *testing.T) {
	utils.LoadEnv("../.env")
	database.Connect()
	router := gin.Default()
	router.DELETE("/travels/:idTravel", controller.DeleteTravel)

	req, err := http.NewRequest("DELETE", "/travels/"+strconv.Itoa(int(idTravel)), nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	var message map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &message)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, "Travel deleted", message["message"])
}

func TestDeleteTravelWithWrongId(t *testing.T) {
	utils.LoadEnv("../.env")
	database.Connect()
	router := gin.Default()
	router.DELETE("/travels/:idTravel", controller.DeleteTravel)

	req, err := http.NewRequest("DELETE", "/travels/11111111", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	req, err = http.NewRequest("DELETE", "/travels/abc", nil)

	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

}
