package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hoannguyen02/golang/bookings/pkg/handlers"
)

func routes() http.Handler {
	r := chi.NewRouter()
	r.Use(NoSurf)
	r.Use(SessionLoad)
	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)
	//  Create a file server which serves files out of the static directory
	fileServer := http.FileServer(http.Dir("./static/"))
	// Register as the handler by using handle function
	// Strip the static prefix before the request reaches the file server
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return r
}
