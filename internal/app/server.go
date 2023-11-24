package app

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Server(handler http.Handler) error {
	server := &http.Server{
		Addr:         ":8000",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		ErrorLog:     log.Default(),
	}

	log.Println("the server is running on http://localhost" + server.Addr)
	// log.Fatal(server.ListenAndServe())
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server err:", err)
		return err
	}

	return nil
}
