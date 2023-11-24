package post

import (
	"fmt"

	"forum/internal/types"
)

func (p *PostDB) CreateLikeDB(postInfo *types.Reaction) error{
	_, err := p.db.Exec("INSERT INTO reactions (snippet_id, user_id, type) VALUES ($1, $2, $3)",
		
		postInfo.PostId,
		postInfo.UserId,
		postInfo.ReactionType,
	)
	if err != nil {
		fmt.Println("create like err:", err)
		return err
	}
	return nil
}

func (p *PostDB) CheckLikeDB(postInfo *types.Reaction) bool {
	reaction := &types.Reaction{}
	err := p.db.QueryRow("SELECT type FROM reactions WHERE snippet_id=$1 AND user_id=$2", postInfo.PostId, postInfo.UserId).Scan(

		&reaction.ReactionType,
	)
	if err != nil {
		// fmt.Println("err means no matches", err)
		return true
	}

	if reaction.ReactionType != postInfo.ReactionType {
		// fmt.Println("Need to remove like and add different")
		p.UpdateLikeDB(postInfo)
		return false
	}

	p.RemoveLikeDB(postInfo)
	return false
}

func (p *PostDB) UpdateLikeDB(postInfo *types.Reaction) {
	_, err := p.db.Exec("UPDATE reactions SET type = $1 WHERE user_id = $2 AND snippet_id = $3", postInfo.ReactionType, postInfo.UserId, postInfo.PostId)
	if err != nil {
		fmt.Println("Update:", err)
		return
	}
}

func (p *PostDB) RemoveLikeDB(postInfo *types.Reaction) {
	_, err := p.db.Exec("DELETE FROM reactions WHERE user_id = $1 AND snippet_id", postInfo.UserId, postInfo.PostId)
	if err != nil {
		fmt.Println("RemoveLikeDB:", err)
	}
}

func (p *PostDB) CountLikes(postId int, reactionType string) int {
	rows, err := p.db.Query("SELECT count(id) FROM reactions WHERE snippet_id=$1 AND type=$2", postId, reactionType)
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
