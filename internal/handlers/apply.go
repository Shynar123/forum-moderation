package handler

import (
	"fmt"
	"net/http"

	"forum/internal/cookies"
	"forum/internal/types"
)

func (h *Handler) apply(w http.ResponseWriter, r *http.Request) {
	cookie, errC := cookies.GetCookie(r)
	if errC != nil {
		fmt.Println("cookie err:", errC)
		return
	}
	user, errU := h.service.UserService.GetUserByToken(cookie.Value)
	if errU != nil {
		fmt.Println("Get User by token err:", errU)
		return
	}

	status := h.service.UserService.GetRequestStatus(user.Id)
	if status == "" {
		request := &types.Request{
			Username: user.Username,
			UserId:   user.Id,
			Status:   "applied",
		}
		h.service.UserService.CreateRequest(request)
	}

	// h.service.UserService.CreateRequest()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
