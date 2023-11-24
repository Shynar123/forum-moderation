package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/internal/types"
)

func (h *Handler) postLike(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/like" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		ErrorPage(w, r, http.StatusMethodNotAllowed)
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	author := h.getUserFromContext(r)
	postId, err := strconv.Atoi(r.Form.Get("post_id"))
	if err != nil {
		fmt.Println("post ID is not a number:", err)
	}
	reaction := &types.Reaction{
		PostId:       postId,
		UserId:       author.Id,
		ReactionType: r.Form.Get("status"),
	}

	if h.service.PostService.CheckLike(reaction) {
		err := h.service.PostService.CreateLike(reaction)
		if err != nil {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
		}

	}

	http.Redirect(w, r, fmt.Sprintf("/post?id=%d", reaction.PostId), http.StatusSeeOther)
}
