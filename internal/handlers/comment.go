package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/internal/types"
)

func (h *Handler) commentCreate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/comment/create" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		ErrorPage(w, r, http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	author := h.getUserFromContext(r)
	postId, err := strconv.Atoi(r.Form.Get("post_id"))
	if err != nil {
		fmt.Println(err)
		return
	}
	comment := &types.Comment{
		PostId:   postId,
		UserId:   author.Id,
		Content:  r.Form.Get("text"),
		Username: author.Username,
	}

	h.service.PostService.CreateComment(comment)

	http.Redirect(w, r, fmt.Sprintf("/post?id=%d", comment.PostId), http.StatusSeeOther)
}
