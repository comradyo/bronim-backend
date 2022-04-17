package repository

import (
	"bronim/pkg/models"
	"github.com/lib/pq"
)

type pqRestaurant struct {
	ID          string         `db:"id"`
	GoogleID    string         `db:"google_id"`
	Name        string         `db:"name"`
	Description string         `db:"description"`
	Address     string         `db:"address"`
	ImgUrl      string         `db:"img_url"`
	PhoneNumber string         `db:"phone_number"`
	Email       string         `db:"email"`
	WebsiteUrl  string         `db:"website_url"`
	Geoposition string         `db:"geoposition"`
	Kitchen     string         `db:"kitchen"`
	Tags        pq.StringArray `db:"tags"`
	Rating      string         `db:"rating"`
}

func toPqRestaurant(restaurant models.Restaurant) pqRestaurant {
	var r pqRestaurant
	r.ID = restaurant.ID
	r.GoogleID = restaurant.GoogleID
	r.Name = restaurant.Name
	r.Description = restaurant.Description
	r.Address = restaurant.Address
	r.ImgUrl = restaurant.ImgUrl
	r.PhoneNumber = restaurant.PhoneNumber
	r.Email = restaurant.Email
	r.WebsiteUrl = restaurant.WebsiteUrl
	r.Geoposition = restaurant.Geoposition
	r.Kitchen = restaurant.Kitchen
	r.Tags = restaurant.Tags
	r.Rating = restaurant.Rating
	return r
}

func toModelRestaurant(restaurant pqRestaurant) models.Restaurant {
	var r models.Restaurant
	r.ID = restaurant.ID
	r.GoogleID = restaurant.GoogleID
	r.Name = restaurant.Name
	r.Description = restaurant.Description
	r.Address = restaurant.Address
	r.ImgUrl = restaurant.ImgUrl
	r.PhoneNumber = restaurant.PhoneNumber
	r.Email = restaurant.Email
	r.WebsiteUrl = restaurant.WebsiteUrl
	r.Geoposition = restaurant.Geoposition
	r.Kitchen = restaurant.Kitchen
	r.Tags = restaurant.Tags
	r.Rating = restaurant.Rating
	return r
}

func (r *Repository) scanRestaurants(query string, args ...interface{}) ([]models.Restaurant, error) {
	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var restaurants []models.Restaurant
	for rows.Next() {
		var t pqRestaurant
		err := rows.StructScan(&t)
		if err != nil {
			return nil, err
		}
		restaurant := toModelRestaurant(t)
		restaurants = append(restaurants, restaurant)
	}
	return restaurants, nil
}
