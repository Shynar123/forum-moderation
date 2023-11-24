package handler

import (
	"forum/internal/render"
	"net/http"
)

func (h *Handler) createdPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/created" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorPage(w, r, http.StatusMethodNotAllowed)
		return
	}
	user := h.getUserFromContext(r)
	posts, err := h.service.PostService.GetPostsByUserID(user.Id)
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError)
		return
	}
	render.Render(w, "ui/html/index.html", render.WebPage{
		IsLoggedin: true,
		Posts:      posts,
	})
}
