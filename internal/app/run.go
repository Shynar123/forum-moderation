package app

import (
	"database/sql"
	"fmt"
	"forum/internal/repository"
	"forum/internal/service"
	"forum/internal/types"
	"log"

	handler "forum/internal/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func Run() {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	err = types.CreateTables(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	repo := repository.NewRepository(db)

	service := service.NewService(repo)

	handler := handler.NewHandler(service)

	log.Fatal(Server(handler.Routes()))
}
