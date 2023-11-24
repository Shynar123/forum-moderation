package service

import (
	"forum/internal/repository"
	"forum/internal/service/post"
	"forum/internal/service/user"
	"forum/internal/types"

)

type Service struct {
	UserService types.UserService
	PostService types.PostService 
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService: user.NewUserService(repo.UserRepo),
		PostService: post.NewPostService(repo.PostRepo),
	}
}
