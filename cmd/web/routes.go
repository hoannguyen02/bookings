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
	return r
}
