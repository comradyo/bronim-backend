package models

type Restaurant struct {
	ID          string   `json:"id" db:"id"`
	GoogleID    string   `json:"google_id" db:"google_id"`
	Address     string   `json:"address" db:"address"`
	Description string   `json:"description" db:"description"`
	Tags        []string `json:"tags" db:"tags"`
	ImgUrl      string   `json:"img_url" db:"img_url"`
	PhoneNumber string   `json:"phone_number" db:"phone_number"`
	Email       string   `json:"email" db:"email"`
	WebsiteUrl  string   `json:"website_url" db:"website_url"`
	Geoposition string   `json:"geoposition" db:"geoposition"`
	Rating      string   `json:"rating" db:"rating"`
}

type Restaurants struct {
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

type Tables struct {
	Arr []Table `json:"tables"`
}

type Profile struct {
	ID          string `json:"id" db:"id"`
	FirebaseID  string `json:"firebase_id" db:"firebase_id"`
	Name        string `json:"restaurant_id" db:"restaurant_id"`
	Surname     string `json:"floor" db:"floor"`
	DateOfBirth string `json:"pos_x" db:"pos_x"`
	Sex         string `json:"pos_y" db:"pos_y"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	AvatarUrl   string `json:"avatar_url" db:"avatar_url"`
}

type Reservation struct {
	ID              string `json:"id" db:"id"`
	TableID         string `json:"table_id" db:"table_id"`
	ProfileID       string `json:"profile_id" db:"profile_id"`
	ReservationDate string `json:"reservation_date" db:"reservation_date"`
	CellID          string `json:"cell_id" db:"cell_id"`
	NumOfCells      string `json:"num_of_cells" db:"num_of_cells"`
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
