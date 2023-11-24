package post

import (
	"fmt"

	"forum/internal/types"
)

func (p *PostDB) Filter(categories []string) ([]*types.Post, error) {
	newpost := []*types.Post{}

	for _, category := range categories {
		rows, err := p.db.Query("SELECT snippet_id FROM categories WHERE category=$1 ORDER BY id DESC", category)
		if err != nil {
			fmt.Println("Filter: ", err)
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			post := &types.Post{}
			if err := rows.Scan(&post.Id); err != nil {
				fmt.Println(err)
			}
			err = p.db.QueryRow("SELECT * FROM snippets WHERE id=$1", post.Id).Scan(

				&post.Id,
				&post.AuthorId,
				&post.AuthorName,
				&post.Title,
				&post.Content,
				&post.Created,
				&post.Status,
			)
			if err != nil {
				fmt.Println("filterErr:", err)
			}

			post.Time = post.Created.Format("15:04 January 02, 2006")
			rows1, err := p.db.Query("SELECT category FROM categories WHERE snippet_id= $1", post.Id)
			if err != nil {
				fmt.Println("GetAllPostsERR: ", err)
			}

			for rows1.Next() {

				var category string
				err := rows1.Scan(&category)
				if err != nil {
					return nil, err
				}
				post.Categories = append(post.Categories, category)
			}

			if err != nil {
				fmt.Println("Filter err:", err)
				return nil, err
			}
			newpost = append(newpost, post)
		}
		for _, post := range newpost {
			post.Likes = p.CountLikes(post.Id, "like")
			post.Dislikes = p.CountLikes(post.Id, "dislike")
			post.Comments = p.GetAllComments(post.Id)
		}

	}

	return newpost, nil
}
