package post

import (
	"forum/internal/types"
)

func (p *PostService) GetPosts(userId int) []*types.Post {
	posts := p.repo.GetPostsByUserId(userId)

	return posts
}
