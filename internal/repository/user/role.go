package user

import (
	"fmt"

	"forum/internal/types"
)

func (db *UserDB) CreateUserRoleDB(user *types.Roles) {
	_, err := db.DB.Exec("INSERT INTO roles (username, role) VALUES ($1, $2)",
		user.Username,
		"User")
	if err != nil {
		fmt.Println("repository userRole err:", err)
		return
	}
}

func (db *UserDB) ChangeUserRoleDB(user *types.Roles) {
	_, err := db.DB.Exec("UPDATE roles SET role = $1 WHERE username = $2", user.Role, user.Username)
	if err != nil {
		fmt.Println("Update role:", err)
		return
	}
}

func (db *UserDB) GetAllUsers() []*types.Roles {
	rows, err := db.DB.Query("SELECT * FROM roles")
	if err != nil {
		fmt.Println("Get All Users error:")
		return nil
	}
	defer rows.Close()
	var users []*types.Roles
	for rows.Next() {
		user := types.Roles{}
		err := rows.Scan(&user.Id, &user.Username, &user.Role)
		if err != nil {
			fmt.Println("Get All users rows err:", err)
			return nil
		}
		users = append(users, &user)
	}

	return users
}

func (db *UserDB) GetUserRole(username string) string {
	user := &types.Roles{}
	err := db.DB.QueryRow("SELECT role FROM roles WHERE username= $1", username).Scan(
		&user.Role)
	if err != nil {
		fmt.Println("GetUserRole:   ", err)
		return "Guest"
	}
	return user.Role
}

func (db *UserDB) UpdateRole(user *types.Roles) {
	_, err := db.DB.Exec("UPDATE roles SET role = $1 WHERE username = $2", user.Role, user.Username)
	if err != nil {
		fmt.Println("Update role:", err)
		return
	}
}
