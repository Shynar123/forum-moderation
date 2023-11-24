package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"forum/internal/types"
	"forum/internal/validity"
)

func (h *Handler) signup(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signup" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}

	var existErr types.ErrText
	if r.Method == http.MethodGet {
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			return
		}

		f := validity.GetForm(r.PostForm)

		var existBool bool

		if f.CheckEmail() && f.CheckName() && f.CheckPassword() {

			user := &types.CreateUserData{
				Username: f.Get("username"),
				Email:    f.Get("email"),
				Password: f.Get("password"),
			}

			existBool, existErr = h.service.UserService.CheckUserExists(user)

			if !existBool {
				err := h.service.UserService.CreateUser(user)
				
				if err != nil {
					log.Fatalln(err)
				}
				http.Redirect(w, r, "/signin", http.StatusSeeOther)
			}
		}
		if !f.CheckPassword() {
			existErr.Pass2 = "Passwords should be the same"
		}

	}
	templ, err := template.ParseFiles("ui/html/signup.html", "ui/html/layout.html")
	if err != nil {
		fmt.Printf("Template not found: %v\n", err)
		ErrorPage(w, r, http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, existErr)
	if err != nil {
		fmt.Printf("Execute error: %v\n", err)
		ErrorPage(w, r, http.StatusInternalServerError)

		return
	}
}
