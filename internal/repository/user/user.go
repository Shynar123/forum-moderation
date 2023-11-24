package user

import (
	"database/sql"
	"fmt"

	t "forum/internal/types"

	"golang.org/x/crypto/bcrypt"
)

type UserDB struct {
	DB *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{DB: db}
}

func (db *UserDB) CreateUserDB(user *t.User) {
	_, err := db.DB.Exec("INSERT INTO users (email, username, password) VALUES ($1, $2, $3)",
		user.Email,
		user.Username,
		user.PasswordHash)
	if err != nil {

		fmt.Println("repository user err:", err)
		return
	}
	userRole := &t.Roles{
		Username: user.Username,
	}
	db.CreateUserRoleDB(userRole)
}

func (db *UserDB) CheckLoginDB(user *t.GetUserData) (int, error) {
	var userMatch t.CreateUserData
	err := db.DB.QueryRow("SELECT id, password FROM users WHERE username= $1", user.Username).Scan(
		&userMatch.Id,

		&userMatch.Password,
	)
	if err != nil {
		// fmt.Println(err)
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userMatch.Password), []byte(user.Password))

	if err != nil {
		// fmt.Println(err)
		return 0, err
	}

	return userMatch.Id, nil
}

func (db *UserDB) GetUserEmailDB(userEmail string) error {
	var userMatch t.CreateUserData
	err := db.DB.QueryRow("SELECT id FROM users WHERE email= $1", userEmail).Scan(
		&userMatch.Id)
	if err != nil {
		fmt.Println("email:", err)
		return err
	}

	return nil
}

func (db *UserDB) GetUserNameDB(userName string) error {
	var nameMatch t.CreateUserData
	err := db.DB.QueryRow("SELECT id  FROM users WHERE  username= $1", userName).Scan(
		&nameMatch.Id)
	if err != nil {
		fmt.Println("username:", err)
		return err
	}

	return nil
}

func (db *UserDB) GetUserIdDB(user *t.CreateUserData) int {
	var nameMatch t.CreateUserData
	err := db.DB.QueryRow("SELECT id  FROM users WHERE  username= $1 OR email=$2", user.Username, user.Email).Scan(
		&nameMatch.Id)
	if err != nil {
		fmt.Println("username id:", err)
		return 0
	}

	return nameMatch.Id
}

func (db *UserDB) AddTokenDB(userid int, cookieToken string) error {
	query := `UPDATE users
	SET token = ?, expires = DATETIME('now', '+7 hours')
	WHERE ? = id` // expiration datetime = now + 1 hours
	if _, err := db.DB.Exec(query, cookieToken, userid); err != nil {
		return err
	}
	return nil
}

func (db *UserDB) RemoveTokenDB(token string) error {
	query := `UPDATE users
	SET token = NULL, expires = NULL
	WHERE token = ?`
	_, err := db.DB.Exec(query, token)
	return err
}

func (db *UserDB) GetUserByToken(token string) (*t.User, error) {
	user := &t.User{}
	err := db.DB.QueryRow("SELECT id, username, email FROM users WHERE token= $1", token).Scan(
		&user.Id,
		&user.Username,
		&user.Email)
	if err != nil {
		fmt.Println("GetUserByToken:   ", err)
		return nil, err
	}
	return user, nil
}
