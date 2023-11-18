package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Timeout(15 * time.Second))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/v1/healthcheck", app.hdlGetHealthcheck)

	router.Get("/v1/employee/{employee_id}", app.hdlGetEmployee)
	router.Get("/v1/employee", app.hdlListEmployee)
	router.Post("/v1/employee", app.hdlPostEmployee)
	router.Post("/v1/employee/upload", app.hdlPostEmployeeCSV)
	router.Put("/v1/employee/{employee_id}", app.hdlPutEmployee)
	router.Delete("/v1/employee/{employee_id}", app.hdlDeleteEmployee)

	return router
}
