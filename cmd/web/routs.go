package main

import (
	"net/http"

	"github.com/alexei/hello-world/pkg/config"
	"github.com/alexei/hello-world/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routs(app *config.AppConfig) http.Handler {
	/*mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.HomeHandler))
	mux.Get("/about", http.HandlerFunc(handlers.AboutHandler))
	return mux
	*/
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Get("/", handlers.HomeHandler)
	mux.Get("/about", handlers.AboutHandler)
	mux.Get("/json", handlers.Json)

	return mux

}
