package foodlisting

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/barbaraalp/go-postgres-redis-rest-api/repository/foodlisting"
)

func GetAllFoods(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	oi, err := foodlisting.GetAll(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("fail to get food listing: %v", err)))
		return
	}

	json.NewEncoder(w).Encode(oi)
}

func GetFoodByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	oi, err := foodlisting.GetByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("fail to get food listing by id (%d): %v", id, err)))
		return
	}

	json.NewEncoder(w).Encode(oi)
}
