package model

import (
	"api/allo-dakar/database"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) Save() (*User, error) {
	_, err := database.DB.Exec(`INSERT INTO "users" (username, email, password) VALUES ($1, $2, $3)`, user.Username, user.Email, user.Password)
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) FindAll() ([]User, error) {
	rows, err := database.DB.Query(`SELECT * FROM "users"`)
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var userTemp User
		err := rows.Scan(&userTemp.Id, &userTemp.Username, &userTemp.Email, &userTemp.Password)
		if err != nil {
			return []User{}, err
		}
		users = append(users, userTemp)
	}

	if err := rows.Err(); err != nil {
		return []User{}, err
	}

	return users, nil
}

func (user *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (user *UserLogin) FindByEmail() (*User, error) {
	row := database.DB.QueryRow(`SELECT * FROM "users" WHERE email = $1`, user.Email)
	var userGetted User

	err := row.Scan(&userGetted.Id, &userGetted.Username, &userGetted.Email, &userGetted.Password)
	if err != nil {
		return &User{}, err
	}
	return &userGetted, nil
}
