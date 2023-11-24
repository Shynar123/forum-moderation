package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) deleteComment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/comment/delete" {
		ErrorPage(w, r, http.StatusNotFound)
	}
	if r.Method != http.MethodPost {
		ErrorPage(w, r, http.StatusMethodNotAllowed)
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	postId, err := strconv.Atoi(r.Form.Get("post_id"))
	if err != nil {
		fmt.Println("post ID is not a number:", err)
	}
	commentId,errA:=strconv.Atoi(r.Form.Get("Id")) 
	if err != nil {
		fmt.Println("errA: ",errA)
		return
	}
	h.service.PostService.DeleteComment(commentId)
	http.Redirect(w, r,fmt.Sprintf("/post?id=%d", postId), http.StatusSeeOther)
}
