package models

type Restaurant struct {
	ID          int      `json:"id" db:"id"`
	GoogleID    string   `json:"google_id" db:"google_id"`
	Address     string   `json:"address" db:"address"`
	Description string   `json:"description" db:"description"`
	Tags        []string `json:"tags" db:"tags"`
	ImgUrl      string   `json:"img_url" db:"img_url"`
	PhoneNumber string   `json:"phone_number" db:"phone_number"`
	Email       string   `json:"email" db:"email"`
	WebsiteUrl  string   `json:"website_url" db:"website_url"`
	Geoposition string   `json:"geoposition" db:"geoposition"`
	Rating      float64  `json:"rating" db:"rating"`
}

type Restaurants struct {
	Arr []Restaurant `json:"restaurants"`
}

type Table struct {
	ID           int    `json:"id" db:"id"`
	RestaurantID string `json:"restaurant_id" db:"restaurant_id"`
	Floor        int    `json:"floor" db:"floor"`
	PosX         string `json:"pos_x" db:"pos_x"`
	PosY         string `json:"pos_y" db:"pos_y"`
	Places       int    `json:"places" db:"places"`
}

type Tables struct {
	Arr []Table `json:"tables"`
}

type Profile struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"restaurant_id" db:"restaurant_id"`
	Surname     int    `json:"floor" db:"floor"`
	DateOfBirth string `json:"pos_x" db:"pos_x"`
	Sex         string `json:"pos_y" db:"pos_y"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Email       string `json:"email" db:"email"`
	AvatarUrl   string `json:"avatar_url" db:"avatar_url"`
}

type Reservation struct {
	ID              int    `json:"id" db:"id"`
	TableID         int    `json:"table_id" db:"table_id"`
	ProfileID       int    `json:"profile_id" db:"profile_id"`
	ReservationDate string `json:"reservation_date" db:"reservation_date"`
	CellID          int    `json:"cell_id" db:"cell_id"`
	NumOfCells      int    `json:"num_of_cells" db:"num_of_cells"`
	Comment         string `json:"comment" db:"comment"`
}

type Reservations struct {
	Arr []Reservation `json:"reservations"`
}

type ProfileReservation struct {
	Restaurant  Restaurant  `json:"restaurant"`
	Reservation Reservation `json:"reservation"`
}

type ProfileReservations struct {
	Arr []ProfileReservation `json:"profile_reservations"`
}
