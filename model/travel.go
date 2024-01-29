package model

import "time"

type Travel struct {
	Id        int       `json:"id"`
	DriverId  int       `json:"driverId"`
	StartArea string    `json:"departureArea"`
	EndArea   string    `json:"arrivalArea"`
	Date      string    `json:"date"`
	StartTime time.Time `json:"departureTime"`
	EndTime   time.Time `json:"arrivalTime"`
	Price     int       `json:"price"`
	Places    int       `json:"places"`
	Phone     string    `json:"phone"`
	Comment   string    `json:"comment"`
}

type TravelPost struct {
	DriverId  int       `json:"driverId"`
	StartArea string    `json:"departureArea"`
	EndArea   string    `json:"arrivalArea"`
	Date      string    `json:"date"`
	StartTime time.Time `json:"departureTime"`
	EndTime   time.Time `json:"arrivalTime"`
	Price     int       `json:"price"`
	Places    int       `json:"places"`
	Phone     string    `json:"phone"`
	Comment   string    `json:"comment"`
}
