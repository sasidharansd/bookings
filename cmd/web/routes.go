package main

import (
	"github.com/aerosasi/bookings/internal/config"
	"github.com/aerosasi/bookings/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)

	mux.Get("/generals-quarters", handler.Repo.Generals)

	mux.Get("/majors-suite", handler.Repo.Majors)
	mux.Post("/search-availability", handler.Repo.Availability)
	mux.Post("/search-availability-json", handler.Repo.AvailabilityJson)

	mux.Get("/contact", handler.Repo.Contact)
	mux.Get("/make-reservation", handler.Repo.Reservation)

	// post

	mux.Post("/search-availability", handler.Repo.PostAvailability)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
