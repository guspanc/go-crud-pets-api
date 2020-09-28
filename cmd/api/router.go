package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/guspanc/go-crud-pets-api/cmd/api/controller"
	"github.com/guspanc/go-crud-pets-api/cmd/api/models"
	"github.com/guspanc/go-crud-pets-api/cmd/api/service"
	"github.com/guspanc/go-crud-pets-api/pkg/logger"

	"github.com/gorilla/mux"
)

func registerRoutes(r *mux.Router, app *Application) {
	registerMiddlewares(r)

	// custom 404
	customNotFound(r)

	v1Router := r.PathPrefix("/v1").Subrouter()

	// pets
	registerPetRoutes(v1Router, app)
}

func registerPetRoutes(r *mux.Router, app *Application) {
	svc := service.NewPetService(app.db)
	ctrl := controller.NewPetController(svc)

	r.HandleFunc("/pets", ctrl.HandleGetPets).Methods(http.MethodGet)
	r.HandleFunc("/pets", ctrl.HandleAddPet).Methods(http.MethodPost)
}

// Custom 404
func customNotFound(r *mux.Router) {
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiError := models.NewAPIError("NotFound", "resource not found")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(apiError)
	})
}

// Middleware

func registerMiddlewares(r *mux.Router) {
	r.Use(loggingMiddleware, mediaTypeHeaderMiddleware, recoveryMiddleware)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.INFO.Printf("starting request %s %s\n", r.Method, r.URL.Path)
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Now().Sub(start).Milliseconds()
		logger.INFO.Printf("finished request %s %s %dms\n", r.Method, r.URL.Path, duration)
	})
}

func mediaTypeHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				apiError := models.NewAPIError("InternalServerError", "internal server error")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(apiError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
