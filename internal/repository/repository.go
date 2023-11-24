package repository

import (
	"database/sql"

	"forum/internal/repository/post"
	"forum/internal/repository/user"
	"forum/internal/types"
)

type Repository struct {
	UserRepo types.UserRepo
	PostRepo types.PostRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepo: user.NewUserDB(db),
		PostRepo: post.NewPostDB(db),
	}
}
