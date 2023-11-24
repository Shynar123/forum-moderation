package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"forum/internal/types"
)

type Message struct {
	Requests []*types.Request
	Reports  []*types.Report
	Role     string
}

func (h *Handler) messages(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/messages" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}

	message := &Message{
		Requests: h.service.UserService.GetAllRequests(),
		Reports:  h.service.PostService.GetAllReports(),
		Role:     "Administrator",
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			fmt.Println("messages err: ", err)
			return
		}
		decision := r.FormValue("button")
		response := r.Form.Get("text")


		if decision == "Delete Post" {
			postId, err := strconv.Atoi(r.Form.Get("post_id"))
			if err != nil {
				fmt.Println("Post ID is not a number", err)
				return
			}
			if len(response) > 0 {
				h.service.PostService.ReportResponse(response, postId)
			}

			h.service.PostService.DeletePost(postId)
			h.service.PostService.ReportComplete(postId)

		} else if decision == "Decline report" {
			postId, err := strconv.Atoi(r.Form.Get("post_id"))
			if err != nil {
				fmt.Println("Post ID is not a number", err)
				return
			}
			if len(response) > 0 {
				h.service.PostService.ReportResponse(response, postId)
			}
			h.service.PostService.ReportComplete(postId)
		}
		user := &types.Roles{
			Username: r.Form.Get("username"),
			Role:     "Moderator",
		}
		if decision == "Accept application" {

			h.service.UserService.UpdateRole(user)
			h.service.UserService.UpdateRequestStatus("accepted", user.Username)
		}
		if decision == "Decline application" {
			h.service.UserService.UpdateRequestStatus("declined", user.Username)

		}
	}


	templ, err := template.ParseFiles("ui/html/messages.html", "ui/html/layout.html")
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
