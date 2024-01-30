package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"

	"api/allo-dakar/model"
)

func CreateTravel(c *gin.Context) {

	var travel model.Travel
	if err := c.BindJSON(&travel); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	travelCreated, err := travel.Save()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, travelCreated)
}

func GetTravels(c *gin.Context) {

	var travel model.Travel

	travels, err := travel.FindAll()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, travels)
}

func GetTravelsByStartAndEnd(c *gin.Context) {

	start := c.Param("start")
	end := c.Param("end")
	var travel model.Travel

	travels, err := travel.FindByStartAndEnd(start, end)

	if err != nil {
		if err.Error() == "No travels found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, travels)
}

func GetTravelsByStartAndEndAndDate(c *gin.Context) {

	start := c.Param("start")
	end := c.Param("end")
	date := c.Param("date")

	var travel model.Travel
	travels, err := travel.FindByStartAndEndAndDate(start, end, date)

	if err != nil {
		if err.Error() == "No travels found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, travels)
}

func GetTravelsByIdDriver(c *gin.Context) {

	idDriver := c.Param("idDriver")

	driverID, err := strconv.ParseInt(idDriver, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid driver ID"})
		return
	}

	var travel model.Travel
	travels, err := travel.FindByIdDriver(driverID)

	if err != nil {
		if err.Error() == "No travels found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, travels)
}

func DeleteTravel(c *gin.Context) {

	idTravel := c.Param("idTravel")

	travelID, err := strconv.ParseInt(idTravel, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid travel ID"})
		return
	}

	var travel model.Travel
	err = travel.Delete(travelID)

	if err != nil {
		if err.Error() == "Travel not found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Travel deleted"})
}
