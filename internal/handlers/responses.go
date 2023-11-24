package handler

import (
	"fmt"
	"net/http"
	"text/template"

	"forum/internal/cookies"
)

func (h *Handler) responses(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/responses" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}
	cookie, _ := cookies.GetCookie(r)

	user, errU := h.service.UserService.GetUserByToken(cookie.Value)
	if errU != nil {
		fmt.Println("Get User by token err:", errU)
		return
	}
	message := &Message{
		Reports: h.service.PostService.GetReportsByID(user.Id),
		Role:    "Moderator",
	}
	templ, err := template.ParseFiles("ui/html/responses.html", "ui/html/layout.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, message)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
