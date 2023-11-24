package handler

import (
	"fmt"
	"net/http"

	"forum/internal/cookies"
	"forum/internal/render"
	"forum/internal/types"
)

var Categories []string

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}

	cookie, errC := cookies.GetCookie(r)
	var data bool
	var status string
	if  errC == nil {
		data = true
		user, err := h.service.UserService.GetUserByToken(cookie.Value)
		if err != nil {
			fmt.Println("Get User by token err:", err)
			return
		}
	
		status=h.service.UserService.GetRequestStatus(user.Id)
	}


	switch r.Method {
	case http.MethodGet:
		//   //need to put somewhere

		posts := []*types.Post{}
		var err error

		query := r.URL.Query()
		Categories = append(Categories, query["category"]...)

		if len(Categories) == 0 {
			posts, err = h.service.PostService.GetAllPosts()
		} else {
			posts, err = h.service.PostService.Filter(Categories)
		}

		if err != nil {
			fmt.Printf("err: %v\n", err)
			ErrorPage(w, r, http.StatusInternalServerError)
			return
		}

		render.Render(w, "ui/html/index.html", render.WebPage{
			IsLoggedin: data,
			Posts:      posts,
			Role:       h.GetRole(w, r),
			Status:     status,
		})

	default:
		ErrorPage(w, r, http.StatusMethodNotAllowed)
	}
}
