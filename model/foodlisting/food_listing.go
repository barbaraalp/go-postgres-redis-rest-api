package foodlisting

type FoodListing struct {
	ID                int     `json:"id"`
	EstablishmentName string  `json:"establishment_name"`
	FoodType          string  `json:"food_type"`
	Location          string  `json:"location"`
	Price             float64 `json:"price"`
	QuantityAvailable int     `json:"quantity_available"`
	Description       string  `json:"description"`
}
