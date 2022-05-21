package models

import (
	"fmt"
	"github.com/lib/pq"
)

type Restaurant struct {
	ID             string         `json:"id" db:"id"`
	GoogleID       string         `json:"google_id" db:"google_id"`
	Name           string         `json:"name" db:"name"`
	Description    string         `json:"description" db:"description"`
	Address        string         `json:"address" db:"address"`
	ImgUrl         string         `json:"img_url" db:"img_url"`
	PhoneNumber    string         `json:"phone_number" db:"phone_number"`
	Email          string         `json:"email" db:"email"`
	WebsiteUrl     string         `json:"website_url" db:"website_url"`
	Geoposition    string         `json:"geoposition" db:"geoposition"`
	Kitchen        string         `json:"kitchen" db:"kitchen"`
	Tags           pq.StringArray `json:"tags" db:"tags"`
	Rating         string         `json:"rating" db:"rating"`
	StartsAtCellID string         `json:"starts_at_cell_id" db:"starts_at_cell_id"`
	EndsAtCellID   string         `json:"ends_at_cell_id" db:"ends_at_cell_id"`
	Date 		   string		  `json:"date" db:"date"`
	Lat			   string		  `json:"lat" db:"lat"`
	Lng			   string		  `json:"lng" db:"lng"`
}

type RestaurantList struct {
	Arr []Restaurant `json:"restaurants"`
}

type Table struct {
	ID           string `json:"id" db:"id"`
	RestaurantID string `json:"restaurant_id" db:"restaurant_id"`
	Floor        string `json:"floor" db:"floor"`
	PosX         string `json:"pos_x" db:"pos_x"`
	PosY         string `json:"pos_y" db:"pos_y"`
	Places       string `json:"places" db:"places"`
}

type TableList struct {
	Arr []Table `json:"tables"`
}

type TableAndReservations struct {
	TableID       string         `json:"table_id" db:"table_id"`
	ReservedTimes pq.StringArray `json:"reserved_cells" db:"reserved_cells"`
}

type TableAndReservationsList struct {
	Arr []TableAndReservations `json:"reservations"`
}

type Profile struct {
	ID          string `json:"id,omitempty" db:"id"`
	FirebaseID  string `json:"firebase_id" db:"firebase_id"`
	Name        string `json:"name" db:"name"`
	Surname     string `json:"surname,omitempty" db:"surname"`
	DateOfBirth string `json:"date_of_birth,omitempty" db:"date_of_birth"`
	Sex         string `json:"sex,omitempty" db:"sex"`
	PhoneNumber string `json:"phone_number,omitempty" db:"phone_number"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	AvatarUrl   string `json:"avatar_url,omitempty" db:"avatar_url"`
}

type Reservation struct {
	ID              string        `json:"id,omitempty" db:"id"`
	TableID         string        `json:"table_id" db:"table_id"`
	ProfileID       string        `json:"profile_id" db:"profile_id"`
	ReservationDate string        `json:"reservation_date" db:"reservation_date"`
	Cells           pq.Int64Array `json:"cells" db:"cells"`
	Comment         string        `json:"comment" db:"comment"`
	NumOfGuests     int           `json:"num_of_guests" db:"num_of_guests"`
}

type ReservationList struct {
	Arr []Reservation `json:"reservations"`
}

type ProfileReservation struct {
	Restaurant  Restaurant  `json:"restaurant" db:"restaurant"`
	Reservation Reservation `json:"reservation" db:"reservation"`
}

type ProfileReservationList struct {
	Arr []ProfileReservation `json:"profile_reservations"`
}

type Err struct {
	ErrStr string `json:"error"`
}

func ErrEmptyValue(valueName string) error {
	return fmt.Errorf(`value '%s' can't be null or empty`, valueName)
}
