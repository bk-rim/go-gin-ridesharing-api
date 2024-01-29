package database

import (
	_ "github.com/lib/pq"

	"database/sql"
	"os"
)

var DB *sql.DB

func Connect() {
	var db *sql.DB
	var err error

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSslmode := os.Getenv("DB_SSL_MODE")

	connStr := "postgres://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=" + dbSslmode

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	DB = db
}

func MigrateDb() {

	migrateUsersTable()
	migrateTravelsTable()

}

func migrateUsersTable() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS "users" (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL,
			email VARCHAR(50) NOT NULL,
			password text NOT NULL
		);
	`)

	if err != nil {
		panic(err)
	}
}

func migrateTravelsTable() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS "travels" (
			id SERIAL PRIMARY KEY,
			id_driver INTEGER NOT NULL,
			start_area VARCHAR(50) NOT NULL,
			end_area VARCHAR(50) NOT NULL,
			price INTEGER NOT NULL,
			date_travel DATE NOT NULL,
			start_time TIME,
			end_time TIME,
			places INTEGER NOT NULL,
			phone VARCHAR(50) NOT NULL,
			comment text 
		);
	`)
	if err != nil {
		panic(err)
	}
}
