package handler

import (
	"fmt"
	"forum/internal/cookies"
	"forum/internal/render"
	"forum/internal/types"
	"net/http"
	
)

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signin" {
		ErrorPage(w, r, http.StatusNotFound)
		return
		
	}
	admin:=&types.GetUserData{
		Username: "Administrator",
	}
	_,err:=h.service.UserService.CheckLogin(admin)
	if err!=nil{
		h.service.UserService.CreateAdministrator()
	}
	
	if r.Method == http.MethodGet {
	} else if r.Method == http.MethodPost { //
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			return
		}
		user := &types.GetUserData{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}


		userid, err := h.service.UserService.CheckLogin(user)
		if err == nil {
			cookieToken := cookies.SetCookie(w) 
			h.service.UserService.AddToken(userid, cookieToken)
	
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			render.Render(w, "ui/html/signin.html", render.WebPage{
				Errtext: "Username or password is incorrect",
			})

			return
		}

	} else {
		ErrorPage(w, r, http.StatusMethodNotAllowed)
		return
	}
	render.Render(w, "ui/html/signin.html", render.WebPage{})
}
