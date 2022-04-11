package repository

import (
	"bronim/pkg/models"
	sql "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) *Repository {
	return &Repository{
		db: database,
	}
}

type profileIDNamed struct {
	id int `db:"id"`
}

func (r *Repository) CreateProfile(profile models.Profile) (models.Profile, error) {
	query := `
insert into profiles 
(firebase_id, name, surname, date_of_birth, sex, phone_number, email, password, avatar_url)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
returning (id)
`
	var insertedID int
	err := r.db.Get(insertedID, query, profile)
	if err != nil {
		return models.Profile{}, err
	}
	query = `
select * from profiles where id = $1;
`
	err = r.db.Get(profile, query, insertedID)
	return profile, err
}

func (r *Repository) GetProfile(profileID string) (models.Profile, error) {
	query := `
select * from profiles where id = $1
`
	var profile models.Profile
	err := r.db.Get(profile, query, profileID)
	return profile, err
}

func (r *Repository) UpdateProfile(profile models.Profile) (models.Profile, error) {
	query := `
update profiles set name = $1, surname = $2, date_of_birth = $3, sex = $4, phone_number = $5, email = $6, avatar_url = $7
where id = $8
returning id;
`
	var updatedID int
	err := r.db.Get(updatedID, query, profile)
	if err != nil {
		return models.Profile{}, err
	}
	query = `
select * from profiles where id = $1;
`
	err = r.db.Get(profile, query, updatedID)
	return profile, err
}

func (r *Repository) CreateRestaurant(restaurant models.Restaurant) (models.Restaurant, error) {
	query := `
insert into restaurants (google_id, address, description, tags, img_url, phone_number, email, website_url, geoposition, rating)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
returning (id)
`
	var insertedID int
	err := r.db.Get(insertedID, query, restaurant)
	if err != nil {
		return models.Restaurant{}, err
	}
	query = `
select * from restaurants where id = $1;
`
	err = r.db.Get(restaurant, query, insertedID)
	return restaurant, err
}

func (r *Repository) GetRestaurant(restaurantID string) (models.Restaurant, error) {
	query := `
select * from restaurants where id = $1;
`
	var restaurant models.Restaurant
	err := r.db.Get(restaurant, query, restaurantID)
	return restaurant, err
}

func (r *Repository) GetTable(tableID string) (models.Table, error) {
	query := `
select * from tables where id = $1;
`
	var table models.Table
	err := r.db.Get(table, query, tableID)
	return table, err
}

func (r *Repository) GetTables(restaurantID string) ([]models.Table, error) {
	query := `
select * from tables where restaurant_id = $1;
`
	rows, err := r.db.Queryx(query, restaurantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []models.Table
	for rows.Next() {
		var t models.Table
		err := rows.StructScan(&t)
		if err != nil {
			return nil, err
		}
		tables = append(tables, t)
	}
	return tables, nil
}

func (r *Repository) CreateReservation(reservation models.Reservation) (models.Reservation, error) {
	query := `
insert into reservations
(table_id, profile_id, reservation_date, cell_id, num_of_cells, comment) 
VALUES 
($1, $2, $3, $4, $5, $6)
returning id
`
	var insertedID int
	err := r.db.Get(insertedID, query, reservation)
	if err != nil {
		return models.Reservation{}, err
	}
	query = `
select * from reservations where id = $1;
`
	err = r.db.Get(reservation, query, insertedID)
	return reservation, err
}

func (r *Repository) GetReservations(tableID string) ([]models.Reservation, error) {
	query := `
select * from reservations where table_id = $1;
`
	rows, err := r.db.Queryx(query, tableID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var reservations []models.Reservation
	for rows.Next() {
		var t models.Reservation
		err := rows.StructScan(&t)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, t)
	}
	return reservations, nil
}

func (r *Repository) GetProfileReservations(profileID string) ([]models.ProfileReservation, error) {
	query := `
select r.*, rsv.* from 
                      reservations rsv 
                          join tables t on rsv.table_id = t.id 
                          join restaurants r on t.restaurant_id = r.id 
                  where profile_id = $1;
`
	rows, err := r.db.Queryx(query, profileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var reservations []models.ProfileReservation
	for rows.Next() {
		var t models.ProfileReservation
		err := rows.StructScan(&t)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, t)
	}
	return reservations, nil
}
