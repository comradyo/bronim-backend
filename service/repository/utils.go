package repository

import (
	"bronim/pkg/models"
)

func (r *Repository) scanRestaurants(query string, args ...interface{}) ([]models.Restaurant, error) {
	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var restaurants []models.Restaurant
	for rows.Next() {
		var restaurant models.Restaurant
		err := rows.StructScan(&restaurant)
		if err != nil {
			return nil, err
		}
		restaurants = append(restaurants, restaurant)
	}
	return restaurants, nil
}
