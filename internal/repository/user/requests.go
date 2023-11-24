package user

import (
	"fmt"

	"forum/internal/types"
)

func (db *UserDB) CreateRequest(request *types.Request) {
	_, err := db.DB.Exec("INSERT INTO requests (username, user_id, status) VALUES ($1, $2, $3)", request.Username, request.UserId, request.Status)
	if err != nil {
		fmt.Println("CreateRequest:", err)
		return
	}
}

func (db *UserDB) DeleteRequest(username string) {
	_, err := db.DB.Exec("DELETE FROM requests WHERE snippet_id=$1", username)
	if err != nil {
		fmt.Println("Delete post:", err)
	}
}

func (db *UserDB) UpdateRequestStatus(status string, userId string) {
	_, err := db.DB.Exec("UPDATE requests SET status = $1 WHERE username = $2", status, userId)
	if err != nil {
		fmt.Println("ReportResponse:", err)
		return
	}
}

func (db *UserDB) GetRequestStatus(userId int) string {
	var status string
	err := db.DB.QueryRow("SELECT status FROM requests WHERE user_id= $1", userId).Scan(
		&status)
	if err != nil {
		return ""
	}
	return status
}

func (db *UserDB) GetAllRequests() []*types.Request {
	query := "SELECT * FROM requests ORDER BY id DESC"

	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println("GetAllrequests1:", err)
		return nil
	}

	defer rows.Close()
	var requests []*types.Request

	for rows.Next() {
		request := types.Request{}

		err := rows.Scan(&request.Id, &request.UserId, &request.Username, &request.Status)
		if err != nil {
			fmt.Println("GetAllrequests2:", err)
			return nil
		}

		requests = append(requests, &request)

	}

	if err := rows.Err(); err != nil {
		fmt.Println("GetAllrequests3:", err)
		return nil
	}
	return requests
}
