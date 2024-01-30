package model

import (
	"api/allo-dakar/database"
	"errors"
	"time"
)

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

func (travel *Travel) Save() (*Travel, error) {
	var id int
	err := database.DB.QueryRow(`INSERT INTO "travels" (id_driven, start, end_area, price, date_travel, start_time, end_time, places, phone, comment)
	 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`, travel.DriverId, travel.StartArea, travel.EndArea, travel.Price,
		travel.Date, travel.StartTime, travel.EndTime, travel.Places, travel.Phone, travel.Comment).Scan(&id)

	if err != nil {
		return &Travel{}, err
	}

	travel.Id = int(id)

	return travel, nil
}

func (travel *Travel) FindAll() ([]Travel, error) {
	rows, err := database.DB.Query(`SELECT * FROM "travels"`)
	if err != nil {
		return []Travel{}, err
	}
	defer rows.Close()

	var travels []Travel
	for rows.Next() {
		var travelTemp Travel
		err := rows.Scan(&travelTemp.Id, &travelTemp.DriverId, &travelTemp.StartArea, &travelTemp.EndArea, &travelTemp.Price, &travelTemp.Date,
			&travelTemp.StartTime, &travelTemp.EndTime, &travelTemp.Places, &travelTemp.Phone, &travelTemp.Comment)
		if err != nil {
			return []Travel{}, err
		}
		travels = append(travels, travelTemp)
	}

	if err := rows.Err(); err != nil {
		return []Travel{}, err
	}

	return travels, nil
}

func (travel *Travel) FindByStartAndEnd(start string, end string) ([]Travel, error) {
	rows, err := database.DB.Query(`SELECT * FROM "travels" WHERE start = $1 AND end_area = $2`, start, end)

	if err != nil {
		return []Travel{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return []Travel{}, errors.New("No travels found")
	}
	var travels []Travel
	for rows.Next() {
		var travelTemp Travel
		err := rows.Scan(&travelTemp.Id, &travelTemp.DriverId, &travelTemp.StartArea, &travelTemp.EndArea, &travelTemp.Price, &travelTemp.Date,
			&travelTemp.StartTime, &travelTemp.EndTime, &travelTemp.Places, &travelTemp.Phone, &travelTemp.Comment)
		if err != nil {
			return []Travel{}, err
		}
		travels = append(travels, travelTemp)
	}

	if err := rows.Err(); err != nil {
		return []Travel{}, err
	}

	return travels, nil
}

func (travel *Travel) FindByStartAndEndAndDate(start string, end string, date string) ([]Travel, error) {
	rows, err := database.DB.Query(`SELECT * FROM "travels" WHERE start = $1 AND end_area = $2 AND date_travel = $3`, start, end, date)

	if err != nil {
		return []Travel{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return []Travel{}, errors.New("No travels found")
	}
	var travels []Travel
	for rows.Next() {
		var travelTemp Travel
		err := rows.Scan(&travelTemp.Id, &travelTemp.DriverId, &travelTemp.StartArea, &travelTemp.EndArea, &travelTemp.Price, &travelTemp.Date,
			&travelTemp.StartTime, &travelTemp.EndTime, &travelTemp.Places, &travelTemp.Phone, &travelTemp.Comment)
		if err != nil {
			return []Travel{}, err
		}
		travels = append(travels, travelTemp)
	}

	if err := rows.Err(); err != nil {
		return []Travel{}, err
	}

	return travels, nil
}

func (travel *Travel) FindByIdDriver(idDriver int64) ([]Travel, error) {

	rows, err := database.DB.Query(`SELECT * FROM "travels" WHERE id_driven = $1`, idDriver)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("No travels found")
	}
	var travels []Travel
	for rows.Next() {
		var travelTemp Travel
		err := rows.Scan(&travelTemp.Id, &travelTemp.DriverId, &travelTemp.StartArea, &travelTemp.EndArea, &travelTemp.Price, &travelTemp.Date,
			&travelTemp.StartTime, &travelTemp.EndTime, &travelTemp.Places, &travelTemp.Phone, &travelTemp.Comment)
		if err != nil {
			return nil, err
		}
		travels = append(travels, travelTemp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return travels, nil
}

func (travel *Travel) Delete(idTravel int64) error {

	rows, err := database.DB.Query(`SELECT * FROM "travels" WHERE id = $1`, idTravel)

	if err != nil {
		return err
	}

	if !rows.Next() {
		return errors.New("Travel not found")
	}

	defer rows.Close()

	_, err = database.DB.Exec(`DELETE FROM "travels" WHERE id = $1`, idTravel)

	if err != nil {
		return err
	}

	return nil
}
