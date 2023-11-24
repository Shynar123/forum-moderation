package post

import (
	"fmt"

	"forum/internal/types"
)

func (p *PostDB) CreateComment(comment *types.Comment) {
	_, err := p.db.Exec("INSERT INTO comments (snippet_id, user_id, author_name, content, created) VALUES ($1, $2, $3, $4,DATETIME('now', '+6 hours'))",
		comment.PostId,
		comment.UserId,
		comment.Username,
		comment.Content,
	)
	if err != nil {
		fmt.Println("repository post err:", err)
		return
	}
}

func (p *PostDB) GetAllComments(postId int) []types.Comment {
	comments := []types.Comment{}
	query := "SELECT * FROM comments WHERE snippet_id= $1 ORDER BY id DESC"

	rows, err := p.db.Query(query, postId)
	if err != nil {
		fmt.Println("GetAllCommentsErr:", err)
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		comment := types.Comment{}

		err := rows.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Username, &comment.Content, &comment.Created)
		if err != nil {
			fmt.Println("GetAllCommentsErr Query:", err)
			return nil
		}
		comment.Time = comment.Created.Format("15:04 January 02, 2006")
		comment.Likes = p.CountLikesComment("like", comment.Id)
		comment.Dislikes = p.CountLikesComment("dislike", comment.Id)
		comments = append(comments, comment)
	}

	return comments
}
func (p *PostDB)DeleteComment(commentId int){
	_, err := p.db.Exec("DELETE FROM comments WHERE id", commentId)
	if err != nil {
		fmt.Println("Delete comment:", err)
	}
}