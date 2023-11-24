package handler

import (
	"net/http"
)

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", h.home)

	mux.HandleFunc("/signup", h.signup)

	mux.HandleFunc("/signin", h.signIn)
	mux.HandleFunc("/logout", h.requireAuth(h.logout))
	mux.HandleFunc("/post/create", h.requireAuth(h.postCreate))
	mux.HandleFunc("/post", h.post)
	mux.HandleFunc("/post/like", h.requireAuth(h.postLike))
	mux.HandleFunc("/comment/like", h.requireAuth(h.commentLike))
	mux.HandleFunc("/comment/create", h.requireAuth(h.commentCreate))
	mux.HandleFunc("/created", h.requireAuth(h.createdPosts))
	mux.HandleFunc("/mylikes", h.requireAuth(h.likedPosts))

	mux.HandleFunc("/google/callback", h.handleGoogleCallback)
	mux.HandleFunc("/google/login", h.handleGoogleLogin)
	mux.HandleFunc("/github/callback", h.handleGithubCallback)
	mux.HandleFunc("/github/login", h.handleGithubLogin)

	mux.HandleFunc("/roles",h.requireAuth(h.Roles)) //only Administrator
	mux.HandleFunc("/post/delete",h.deletePost) //only Moderator or Administator
	mux.HandleFunc("/comment/delete",h.deleteComment)//only Moderator or Administator
	mux.HandleFunc("/report/",h.report)//only Moderator 
	mux.HandleFunc("/apply",h.apply) //only User
	mux.HandleFunc("/messages",h.messages)
	mux.HandleFunc("/responses",h.responses)
	mux.HandleFunc("/post_moderator",h.postModerator)
	return h.middleware(mux)
}
