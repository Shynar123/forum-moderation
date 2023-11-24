package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/delete" {
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
	// postMod := r.Form.Get("post_mod")
	// fmt.Printf("postMod: %v\n", postMod)
	postId, errA := strconv.Atoi(r.Form.Get("post_id"))
	// fmt.Printf("postId: %v\n", postId)
	if err != nil {
		fmt.Println("errA: ", errA)
		return
	}
	h.service.PostService.DeletePost(postId)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
