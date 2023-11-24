package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/internal/types"
)

func (h *Handler) commentLike(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/comment/like" {
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
	commentId, err := strconv.Atoi(r.Form.Get("comment_id"))
	if err != nil {
		fmt.Println("comment ID is not a number:", err)
	}
	reaction := &types.Reaction{
		CommentId:    commentId,
		PostId:       postId,
		UserId:       author.Id,
		ReactionType: r.Form.Get("status"),
	}

	if h.service.PostService.CheckLikeComment(reaction) {
		h.service.PostService.CreateLikeComment(reaction)
	}

	http.Redirect(w, r, fmt.Sprintf("/post?id=%d", reaction.PostId), http.StatusSeeOther)
}
