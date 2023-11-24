package handler

import (
	"net/http"
	"text/template"

	"forum/internal/types"
)

func ErrorPage(w http.ResponseWriter, r *http.Request, status int) {
	e := types.Err{
		StatusCode: status,
		StatusText: http.StatusText(status),
	}

	w.WriteHeader(status)

	tmpl, err := template.ParseFiles("ui/html/errorPage.html", "ui/html/layout.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, e)

	if err != nil {
	
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
