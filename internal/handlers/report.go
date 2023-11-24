package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/internal/types"
)

func (h *Handler) report(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	postId, err := strconv.Atoi(r.Form.Get("post_id"))
	if err != nil {
		fmt.Println(err)
		return
	}
	author := h.getUserFromContext(r)
	data := &types.Report{
		ReportType:  r.Form.Get("report_type"),
		PostId:      postId,
		ModeratorId: author.Id,
		PostTitle:   r.Form.Get("post_title"),
	}
	h.service.PostService.ReportPost(data)
	http.Redirect(w, r, fmt.Sprintf("/post?id=%d", postId), http.StatusSeeOther)
}
