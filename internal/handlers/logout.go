package handler

import (
	"fmt"
	"net/http"

	"forum/internal/cookies"
)

func (h *Handler) logout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}
	cookie, err := cookies.GetCookie(r)
	if cookie == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if err != nil {
		fmt.Println(err)
		ErrorPage(w, r, http.StatusInternalServerError)
	}

	h.service.UserService.RemoveToken(cookie.Value)
	cookies.DeleteCookie(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
