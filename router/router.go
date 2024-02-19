package router

import (
	"github.com/barbaraalp/go-postgres-redis-rest-api/handler/foodlisting"
	"github.com/go-chi/chi"
)

func SetFoodRoutes(r chi.Router) {
	r.Get("/foods", foodlisting.GetAllFoods)
	r.Get("/foods/{id}", foodlisting.GetFoodByID)
}
