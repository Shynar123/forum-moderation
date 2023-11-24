package handler

import (
	"fmt"
	"net/http"
	"text/template"

	"forum/internal/cookies"
	"forum/internal/types"
)

type Data struct {
	Users []*types.Roles
	Role  string
}

func (h *Handler) Roles(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/roles" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			return
		}
		user := &types.Roles{
			Username: r.Form.Get("username"),
			Role:     r.FormValue("roles"),
		}
		h.service.UserService.UpdateRole(user)
	}

	templ, err := template.ParseFiles("ui/html/User_roles.html", "ui/html/layout.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	data := &Data{
		Users: h.service.UserService.GetAllUsers(),
		Role:  h.GetRole(w, r),
	}

	err = templ.Execute(w, data)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetRole(w http.ResponseWriter, r *http.Request) string {
	cookie, errC := cookies.GetCookie(r)
	if errC != nil {
		return "Guest"
	}
	if cookie == nil {
		return "Guest"
	}
	user, errU := h.service.UserService.GetUserByToken(cookie.Value)
	if errU != nil {
		fmt.Println("Get User by token err:", errU)
		return ""
	}

	role := h.service.UserService.GetUserRole(user.Username)
	
	return role
}
