package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"

	"api/allo-dakar/model"

	"api/allo-dakar/database"
)

func CreateTravel(c *gin.Context) {

	var travel model.TravelPost
	if err := c.BindJSON(&travel); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row := database.DB.QueryRow(`INSERT INTO "travels" (id_driven, start, end_area, price, date_travel, start_time, end_time, places, phone, comment)
	 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`, travel.DriverId, travel.StartArea, travel.EndArea, travel.Price,
		travel.Date, travel.StartTime, travel.EndTime, travel.Places, travel.Phone, travel.Comment)

	var id int64
	err := row.Scan(&id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	travelCreated := model.Travel{
		Id:        int(id),
		DriverId:  travel.DriverId,
		StartArea: travel.StartArea,
		EndArea:   travel.EndArea,
		Price:     travel.Price,
		Date:      travel.Date,
		StartTime: travel.StartTime,
		EndTime:   travel.EndTime,
		Places:    travel.Places,
		Phone:     travel.Phone,
		Comment:   travel.Comment,
	}

	c.IndentedJSON(http.StatusCreated, travelCreated)
}

func GetTravels(c *gin.Context) {

	rows, err := database.DB.Query(`SELECT * FROM "travels"`)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer rows.Close()

	var travels []model.Travel
	for rows.Next() {
		var travelTemp model.Travel
		err := rows.Scan(&travelTemp.Id, &travelTemp.DriverId, &travelTemp.StartArea, &travelTemp.EndArea, &travelTemp.Price, &travelTemp.Date,
			&travelTemp.StartTime, &travelTemp.EndTime, &travelTemp.Places, &travelTemp.Phone, &travelTemp.Comment)
		if err != nil {
			panic(err)
		}
		travels = append(travels, travelTemp)
	}

	if err := rows.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, travels)
}

func GetTravelsByStartAndEnd(c *gin.Context) {

	start := c.Param("start")
	end := c.Param("end")
	rows, err := database.DB.Query(`SELECT * FROM "travels" WHERE start = $1 AND end_area = $2`, start, end)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var travels []model.Travel
	for rows.Next() {
		var travelTemp model.Travel
		err := rows.Scan(&travelTemp.Id, &travelTemp.DriverId, &travelTemp.StartArea, &travelTemp.EndArea, &travelTemp.Price, &travelTemp.Date,
			&travelTemp.StartTime, &travelTemp.EndTime, &travelTemp.Places, &travelTemp.Phone, &travelTemp.Comment)
		if err != nil {
			panic(err)
		}
		travels = append(travels, travelTemp)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, travels)
}

func GetTravelsByStartAndEndAndDate(c *gin.Context) {

	start := c.Param("start")
	end := c.Param("end")
	date := c.Param("date")
	rows, err := database.DB.Query(`SELECT * FROM "travels" WHERE start = $1 AND end_area = $2 AND date_travel = $3`, start, end, date)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var travels []model.Travel
	for rows.Next() {
		var travelTemp model.Travel
		err := rows.Scan(&travelTemp.Id, &travelTemp.DriverId, &travelTemp.StartArea, &travelTemp.EndArea, &travelTemp.Price, &travelTemp.Date,
			&travelTemp.StartTime, &travelTemp.EndTime, &travelTemp.Places, &travelTemp.Phone, &travelTemp.Comment)
		if err != nil {
			panic(err)
		}
		travels = append(travels, travelTemp)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, travels)
}

func GetTravelsByIdDriver(c *gin.Context) {

	idDriver := c.Param("idDriver")
	rows, err := database.DB.Query(`SELECT * FROM "travels" WHERE id_driven = $1`, idDriver)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var travels []model.Travel
	for rows.Next() {
		var travelTemp model.Travel
		err := rows.Scan(&travelTemp.Id, &travelTemp.DriverId, &travelTemp.StartArea, &travelTemp.EndArea, &travelTemp.Price, &travelTemp.Date,
			&travelTemp.StartTime, &travelTemp.EndTime, &travelTemp.Places, &travelTemp.Phone, &travelTemp.Comment)
		if err != nil {
			panic(err)
		}
		travels = append(travels, travelTemp)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, travels)
}

func DeleteTravel(c *gin.Context) {

	idTravel := c.Param("idTravel")

	rows, err := database.DB.Query(`SELECT * FROM "travels" WHERE id = $1`, idTravel)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !rows.Next() {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Travel not found"})
		return
	}

	defer rows.Close()

	_, err = database.DB.Query(`DELETE FROM "travels" WHERE id = $1`, idTravel)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Travel deleted"})
}
