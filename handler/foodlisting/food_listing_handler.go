package foodlisting

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	mfoodlisting "github.com/barbaraalp/go-postgres-redis-rest-api/model/foodlisting"
	"github.com/barbaraalp/go-postgres-redis-rest-api/repository/foodlisting"
	"github.com/go-chi/chi"
	redis "github.com/go-redis/redis/v8"
)

var redisHost = "redis:6379"
var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis: ", err)
	}
}

func GetAllFoods(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cacheKey := "foods:all"
	cacheResult, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		log.Println("Cache hit for foods:all")
		var foods []mfoodlisting.FoodListing
		err = json.Unmarshal([]byte(cacheResult), &foods)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error unmarshalling cached foods: %v", err)))
		}

		json.NewEncoder(w).Encode(foods)
		return
	}

	log.Println("No cache")
	foods, err := foodlisting.GetAll(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("fail to get food listing: %v", err)))
		return
	}
	data, err := json.Marshal(foods)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error marshalling foods for cache: %v", err)))
	}
	err = redisClient.Set(ctx, cacheKey, data, time.Second*10).Err()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error setting cache in Redis: %v", err)))
	}

	json.NewEncoder(w).Encode(foods)
}

func GetFoodByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	foods, err := foodlisting.GetByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("fail to get food listing by id (%d): %v", id, err)))
		return
	}

	json.NewEncoder(w).Encode(foods)
}
