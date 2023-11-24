package handler

import (
	"context"
	"net/http"

	"forum/internal/cookies"
	"forum/internal/types"
)

type contextKey string

const ctxKey contextKey = "user"

func (h *Handler) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := cookies.GetCookie(r)

		user := &types.User{}
		switch err {
		case http.ErrNoCookie:

		case nil:
			user, err = h.service.UserService.GetUserByToken(cookie.Value)
			if err != nil {
				cookies.DeleteCookie(w)
				http.Redirect(w, r, "/signin", http.StatusSeeOther)
				return
			}
		default:
			ErrorPage(w, r, http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), ctxKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := h.getUserFromContext(r)
		if len(user.Username) == 0 {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
