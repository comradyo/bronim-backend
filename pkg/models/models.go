package models

type Restaurant struct {
	ID          string   `json:"id" db:"id"`
	GoogleID    string   `json:"google_id" db:"google_id"`
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	Address     string   `json:"address" db:"address"`
	ImgUrl      string   `json:"img_url" db:"img_url"`
	PhoneNumber string   `json:"phone_number" db:"phone_number"`
	Email       string   `json:"email" db:"email"`
	WebsiteUrl  string   `json:"website_url" db:"website_url"`
	Geoposition string   `json:"geoposition" db:"geoposition"`
	Kitchen     string   `json:"kitchen" db:"kitchen"`
	Tags        []string `json:"tags" db:"tags"`
	Rating      string   `json:"rating" db:"rating"`
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
	TableID       string `json:"table_id"`
	ReservedTimes []int  `json:"reserved_times"`
}

type TableAndReservationsList struct {
	Arr []TableAndReservations `json:"reservations"`
}

type Profile struct {
	ID          string `json:"id,omitempty" db:"id"`
	FirebaseID  string `json:"firebase_id" db:"firebase_id"`
	Name        string `json:"name" db:"name"`
	Surname     string `json:"surname" db:"surname"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth"`
	Sex         string `json:"sex" db:"sex"`
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

type ReservationList struct {
	Arr []Reservation `json:"reservations"`
}

type ProfileReservation struct {
	Restaurant  Restaurant  `json:"restaurant"`
	Reservation Reservation `json:"reservation"`
}

type ProfileReservationList struct {
	Arr []ProfileReservation `json:"profile_reservations"`
}

type Err struct {
	ErrStr string `json:"error"`
}
