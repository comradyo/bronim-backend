package places

import (
	"bronim/pkg/models"
)

type GooglePlacesClient struct {
}

func NewGooglePlacesClient(credentials string) *GooglePlacesClient {
	return &GooglePlacesClient{}
}

func (c *GooglePlacesClient) GetNearestRestaurants(lat string, lon string) ([]models.Restaurant, error) {
	return nil, nil
}
