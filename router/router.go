package router

import (
	"github.com/barbaraalp/go-postgres-redis-rest-api/handler/foodlisting"
	"github.com/go-chi/chi"
)

func SetFoodRoutes(r chi.Router) {
	r.Get("/food-listings", foodlisting.GetAllFoods)
	r.Get("/food-listings/{id}", foodlisting.GetFoodByID)
}
