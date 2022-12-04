package routes

import (
	"net/http"

	"github.com/ashfaqrobin/booking-app/pkg/config"
	"github.com/ashfaqrobin/booking-app/pkg/handlers"
	"github.com/ashfaqrobin/booking-app/pkg/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middlewares.NoSurf)
	mux.Use(middlewares.SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileserver := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/*", http.StripPrefix("/public", fileserver))

	return mux
}
