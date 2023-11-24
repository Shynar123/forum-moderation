package post

import "forum/internal/types"

func (p *PostService) CreateLike(postInf *types.Reaction) error{
	return p.repo.CreateLikeDB(postInf)
}

func (p *PostService) CheckLike(postInf *types.Reaction) bool {
	return p.repo.CheckLikeDB(postInf)
}
func (p *PostService)CountLikes(postId int, reactionType string)int{
	return p.repo.CountLikes(postId,reactionType)
}
func (p *PostService) CheckLikeComment(postInf *types.Reaction) bool {
	return p.repo.CheckLikeCommentDB(postInf)
}
func (p *PostService) CreateLikeComment(postInf *types.Reaction) {
	p.repo.CreateLikeCommentDB(postInf)
}