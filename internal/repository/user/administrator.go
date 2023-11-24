package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (db *UserDB) CreateAdministrator() {
	hashedPW, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Admin pass hash err:", err)
		return
	}
	_, err = db.DB.Exec("INSERT INTO users (email, username, password) VALUES ($1, $2, $3)",
		"admin@gmail.com",
		"Administrator",
		hashedPW)
	if err != nil {

		fmt.Println("repository admin err:", err)
		return
	}
	_, err = db.DB.Exec("INSERT INTO roles (username, role) VALUES ($1, $2)",
		"Administrator",
		"Administrator")

	if err != nil {
		fmt.Println("repository adminRole err:", err)
		return
	}
}
