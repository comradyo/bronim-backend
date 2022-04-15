package service

import (
	"bronim/pkg/models"
)

type RestaurantsFilter struct {
}

type Repository interface {
	CreateProfile(profile models.Profile) (models.Profile, error)
	GetProfile(profileID string) (models.Profile, error)
	//MVP2// UpdateProfile(profileID string, profile models.Profile) (models.Profile, error)
	CreateRestaurant(restaurant models.Restaurant) (models.Restaurant, error)
	GetRestaurant(restaurantID string) (models.Restaurant, error)
	GetTable(tableID string) (models.Table, error)
	GetTables(restaurantID string) ([]models.Table, error)
	CreateReservation(reservation models.Reservation) (models.Reservation, error)
	GetReservations(tableID string) ([]models.Reservation, error)
	GetProfileReservations(profileID string) ([]models.ProfileReservation, error)
	GetPopularRestaurants() ([]models.Restaurant, error)
	//В деливери идем на GoogleAPI с координатами, полученными из запроса, берем айдишники близжайших ресторанов,
	GetNearestRestaurants(apiRestaurants []models.Restaurant) ([]models.Restaurant, error)
	GetNewRestaurants() ([]models.Restaurant, error)
	GetKitchenRestaurants(kitchen string) ([]models.Restaurant, error)
	//MVP2// GetRestaurantsByFilter(filter RestaurantsFilter) ([]models.Restaurant, error)
	//MVP2// GetFavouriteRestaurants(profileID string) ([]models.Restaurant, error)
}