package post

import (
	"fmt"

	"forum/internal/types"
)

func (p *PostDB) GetPostsByUserId(userId int) []*types.Post {
	rows, err := p.db.Query("SELECT snippet_id FROM reactions WHERE user_id=$1", userId)
	if err != nil {
		fmt.Println("GetRowsByUserId1: ", err)
		return nil
	}

	defer rows.Close()
	var posts []*types.Post
	for rows.Next() {
		post := types.Post{}
		err := rows.Scan(&post.Id)
		if err != nil {
			fmt.Println("GetRowsByUserId2:", err)
			return nil
		}
		rows2, err := p.db.Query("SELECT * FROM snippets WHERE id=$1", post.Id)
		if err != nil {
			fmt.Println("GetRowsByUserId3:", err)
			return nil
		}
		for rows2.Next() {

			defer rows2.Close()
			err = rows2.Scan(&post.Id, &post.AuthorId, &post.AuthorName, &post.Title, &post.Content, &post.Created, &post.Status)
			if err != nil {
				fmt.Println("GetRowsByUserId4:", err)
				return nil
			}
			post.Time = post.Created.Format("15:04 January 02, 2006")
			posts = append(posts, &post)
		}

	}

	return posts
}
