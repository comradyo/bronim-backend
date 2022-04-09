package service

import (
	"bronim/pkg/models"
)

type Repository interface {
	CreateProfile(profile models.Profile) (models.Profile, error)
	GetProfile(id string) (models.Profile, error)
	UpdateProfile(id string, profile models.Profile) (models.Profile, error)
	CreateRestaurant(profile models.Restaurant) (models.Restaurant, error)
	GetRestaurant(id string) (models.Restaurant, error)
	GetTables(id string) ([]models.Table, error)
	CreateReservation(restaurantID string, tableID string, reservation models.Reservation) (models.Reservation, error)
	GetReservations(restaurantID string, tableID string) ([]models.Reservation, error)
	GetProfileReservations(id string) ([]models.ProfileReservation, error)
}
