package handler

import (
	"net/http"

	"forum/internal/render"
)

func (h *Handler) likedPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/mylikes" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}
	author := h.getUserFromContext(r)
	posts := h.service.PostService.GetPosts(author.Id)

	render.Render(w, "ui/html/liked.html", render.WebPage{
		Posts:      posts,
	})
}
