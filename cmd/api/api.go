package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// setting up the configuration for the application and injecting dependencies
type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)
	})

	return router
}

func (app *application) serve(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Starting server on port %s", app.config.addr)
	return srv.ListenAndServe()
}
