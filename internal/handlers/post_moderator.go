package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) postModerator(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post_moderator" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Println("post_moderator err: ", err)
		return
	}
	postId, err := strconv.Atoi(r.Form.Get("post_id"))
	if err != nil {
		fmt.Println("Post ID is not a number", err)
		return
	}
	postDecision := r.FormValue("button")

	if postDecision == "Accept" {
		h.service.PostService.UpdatePostStatus(postId)
	} else if postDecision == "Delete" {
		h.service.PostService.DeletePost(postId)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
