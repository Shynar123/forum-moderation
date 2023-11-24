package post

import (
	"fmt"

	"forum/internal/types"
)

func (p *PostDB) CreateLikeCommentDB(postInfo *types.Reaction) {
	_, err := p.db.Exec("INSERT INTO reactions (comment_id, user_id, type) VALUES ($1, $2, $3)",

		postInfo.CommentId,
		postInfo.UserId,
		postInfo.ReactionType,
	)
	if err != nil {
		fmt.Println("repository post err:", err)
		return
	}
}

func (p *PostDB) CheckLikeCommentDB(postInfo *types.Reaction) bool {
	reaction := &types.Reaction{}
	err := p.db.QueryRow("SELECT type FROM reactions WHERE comment_id=$1 AND user_id=$2", postInfo.CommentId, postInfo.UserId).Scan(

		&reaction.ReactionType,
	)
	if err != nil {
		// fmt.Println("err means no matches", err)
		return true
	}

	if reaction.ReactionType != postInfo.ReactionType {
		// fmt.Println("Need to remove like and add different")
		p.UpdateLikeCommentDB(postInfo)
		return false
	}

	p.RemoveLikeCommentDB(postInfo)
	return false
}

func (p *PostDB) UpdateLikeCommentDB(postInfo *types.Reaction) {
	_, err := p.db.Exec("UPDATE reactions SET type = $1 WHERE user_id = $2 AND comment_id=$4", postInfo.ReactionType, postInfo.UserId, postInfo.CommentId)
	if err != nil {
		fmt.Println("Update:", err)
		return
	}
}

func (p *PostDB) RemoveLikeCommentDB(postInfo *types.Reaction) {
	_, err := p.db.Exec("DELETE FROM reactions WHERE user_id = $1 AND comment_id=$2", postInfo.UserId, postInfo.CommentId)
	if err != nil {
		fmt.Println("RemoveLikeDB:", err)
	}
}

func (p *PostDB) CountLikesComment(reactionType string, commentId int) int {
	rows, err := p.db.Query("SELECT count(id) FROM reactions WHERE type=$1 AND comment_id=$2", reactionType, commentId)

	if err != nil {
		fmt.Println("CountLikes: ", err)
		return 0
	}
	defer rows.Close()
	count := 0

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			fmt.Println(err)
		}
	}
	return count
}
