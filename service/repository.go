package service

import (
	"bronim/pkg/models"
)

type RestaurantsFilter struct {
}

type Repository interface {
	CreateProfile(profile models.Profile) (models.Profile, error)
	GetProfile(profileID string) (models.Profile, error)
	UpdateProfile(profileID string, profile models.Profile) (models.Profile, error)
	CreateRestaurant(restaurant models.Restaurant) (models.Restaurant, error)
	GetRestaurant(restaurantID string) (models.Restaurant, error)
	GetTable(tableID string) (models.Table, error)
	GetTables(restaurantID string) ([]models.Table, error)
	CreateReservation(reservation models.Reservation) (models.Reservation, error)
	Subscribe(profileID, restID int) error
	Unsubscribe(profileID, restID int) error
	GetProfileReservations(profileID string) ([]models.ProfileReservation, error)
	GetPopularRestaurants() ([]models.Restaurant, error)
	GetNearestRestaurants(apiRestaurants []models.Restaurant) ([]models.Restaurant, error)
	GetNewRestaurants() ([]models.Restaurant, error)
	GetFavouritesRestaurants(uuid int) ([]models.Restaurant, error)
	GetKitchenRestaurants(filter GetRestaurantsFilter) ([]models.Restaurant, error)
	GetRestaurantReservations(restaurantID, date string, numOfGuests string) ([]models.TableAndReservations, error)
}

type GetRestaurantsFilter struct {
	Cuisine string
	Tags    []string
}
