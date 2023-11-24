package post

import "forum/internal/types"

func (p *PostService) CreateComment(comment *types.Comment) {
	p.repo.CreateComment(comment)
}
func (p *PostService)DeleteComment(commentId int){
	p.repo.DeleteComment(commentId)
}