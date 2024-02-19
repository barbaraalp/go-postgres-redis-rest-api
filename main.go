package main

import (
	"log"
	"net/http"

	"github.com/barbaraalp/go-postgres-redis-rest-api/router"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	router.SetFoodRoutes(r)

	log.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", r)
}
