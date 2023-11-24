package handler

import (
	"fmt"
	"net/http"
	"strings"

	"forum/internal/render"
	"forum/internal/types"
)

func (h *Handler) postCreate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/create" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}
	if r.Method == http.MethodGet {
	} else if r.Method == http.MethodPost { //
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			return
		}

		// user:=&types.GetUserData{
		// 	Username: r.Form.Get("username"),
		// 	Password: r.Form.Get("password"),
		// }

		author := h.getUserFromContext(r)

		categoriesForm := r.PostForm["categories"]

		post := &types.CreatePost{
			AuthorId:   author.Id,
			AuthorName: author.Username,
			Title:      strings.TrimSpace(r.Form.Get("title")),
			Content:    strings.TrimSpace(r.Form.Get("text")),
			Categories: categoriesForm,
		}

		if len(post.Title) == 0 || len(post.Content) == 0 {
			var ErrText string = "Please fill out all fields"
			w.WriteHeader(http.StatusBadRequest)
			render.Render(w, "ui/html/createpost.html", render.WebPage{
				Errtext: ErrText,
			})
			return
		}

		_, err = h.service.PostService.CreateNewPost(post)
		if err != nil {
			fmt.Println(err)
			ErrorPage(w, r, http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		ErrorPage(w, r, http.StatusMethodNotAllowed)
		return
	}
	render.Render(w, "ui/html/createpost.html", render.WebPage{})
}
