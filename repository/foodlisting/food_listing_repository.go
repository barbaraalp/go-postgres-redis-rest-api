package foodlisting

import (
	"context"
	"database/sql"
	"log"

	"github.com/barbaraalp/go-postgres-redis-rest-api/model/foodlisting"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=postgres dbname=postgres host=192.168.100.14 port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

const (
	getAll = `
		SELECT
			id,
			establishment_name,
			food_type,
			location,
			price,
			quantity_available,
			description
		FROM food_listing
	`

	getById = `
		SELECT id,
		establishment_name,
		food_type,
		location, price,
		quantity_available,
		description
	FROM food_listing
	WHERE id = $1
	`
)

func GetAll(ctx context.Context) ([]foodlisting.FoodListing, error) {
	rows, err := db.QueryContext(ctx, getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foods []foodlisting.FoodListing
	for rows.Next() {
		var food foodlisting.FoodListing
		if err := rows.Scan(&food.ID, &food.EstablishmentName, &food.FoodType, &food.Location, &food.Price, &food.QuantityAvailable, &food.Description); err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return foods, nil
}

func GetByID(ctx context.Context, id int) (foodlisting.FoodListing, error) {
	var food foodlisting.FoodListing
	err := db.QueryRowContext(ctx, getById, id).
		Scan(&food.ID, &food.EstablishmentName, &food.FoodType, &food.Location, &food.Price, &food.QuantityAvailable, &food.Description)
	if err != nil {
		return foodlisting.FoodListing{}, err
	}

	return food, nil
}
