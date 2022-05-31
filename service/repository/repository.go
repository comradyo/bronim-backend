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
(firebase_id, name, email)
values ($1, $2, $3)
returning (firebase_id)
`
	var insertedFirebaseID string
	err := r.db.Get(&insertedFirebaseID, query,
		profile.FirebaseID,
		profile.Name,
		//profile.Surname,
		//profile.DateOfBirth,
		//profile.Sex,
		//profile.PhoneNumber,
		profile.Email,
		//profile.Password,
		//profile.AvatarUrl,
	)
	if err != nil {
		return models.Profile{}, err
	}
	return r.GetProfile(insertedFirebaseID)
}

func (r *Repository) GetProfile(profileID string) (models.Profile, error) {
	query := `
select * from profiles where firebase_id = $1
`
	var profile models.Profile
	err := r.db.Get(&profile, query, profileID)
	return profile, err
}

func (r *Repository) UpdateProfile(profileID string, profile models.Profile) (models.Profile, error) {
	query := `
update profiles set name = $1, surname = $2, date_of_birth = $3, sex = $4, phone_number = $5, email = $6, avatar_url = $7
where firebase_id = $8
returning id;
`
	var updatedID int
	err := r.db.Get(&updatedID, query,
		profile.Name,
		profile.Surname,
		profile.DateOfBirth,
		profile.Sex,
		profile.PhoneNumber,
		profile.Email,
		profile.AvatarUrl,
		profileID)
	if err != nil {
		return models.Profile{}, err
	}
	query = `
select * from profiles where id = $1;
`
	err = r.db.Get(&profile, query, updatedID)
	return profile, err
}

func (r *Repository) CreateRestaurant(restaurant models.Restaurant) (models.Restaurant, error) {
	query := `
insert into restaurants (google_id, address, description, img_url, phone_number, email, website_url, kitchen, tags, rating, date, lat, lng)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10::text[], $11)
returning (id)
`
	var insertedID string
	err := r.db.Get(&insertedID, query,
		restaurant.GoogleID,
		restaurant.Address,
		restaurant.Description,
		restaurant.ImgUrl,
		restaurant.PhoneNumber,
		restaurant.Email,
		restaurant.WebsiteUrl,
		restaurant.Kitchen,
		restaurant.Tags,
		restaurant.Rating,
		restaurant.Date,
		restaurant.Lat,
		restaurant.Lng,
	)
	if err != nil {
		return models.Restaurant{}, err
	}
	return r.GetRestaurant(insertedID)
}

func (r *Repository) GetRestaurant(restaurantID string) (models.Restaurant, error) {
	query := `
select * from restaurants where id = $1;
`
	var restaurant models.Restaurant
	err := r.db.Get(&restaurant, query, restaurantID)
	return restaurant, err
}

//TODO: популярность
func (r *Repository) GetPopularRestaurants() ([]models.Restaurant, error) {
	query := `
	select * from restaurants 
	order by rating desc 
	LIMIT 10;
`
	return r.scanRestaurants(query)
}

//TODO: insert if not exists
/*
TODO: сначала инсертим, потом достаём по айдишникам из apiRestaurants
*/
//В деливери идем на GoogleAPI с координатами, полученными из запроса, берем айдишники близжайших ресторанов,
func (r *Repository) GetNearestRestaurants(apiRestaurants []models.Restaurant) ([]models.Restaurant, error) {
	query := `
select * from restaurants;
`
	return r.scanRestaurants(query)
}

func (r *Repository) GetNewRestaurants() ([]models.Restaurant, error) {
	query := `
	select * from restaurants 
	order by date desc 
	LIMIT 10;
`
	return r.scanRestaurants(query)
}

func (r *Repository) GetFavouritesRestaurants(userID int) ([]models.Restaurant, error) {
	query := `
	select res.id, name, description, address, img_url, phone_number, 
	email, website_url, kitchen, tags, rating, starts_at_cell_id, ends_at_cell_id, 
	date, lat, lng from restaurants as res
	join favourites as fav on res.id = fav.restaurant_id
	where fav.profile_id = $1;
	`
	return r.scanRestaurants(query, userID)
}

func (r *Repository) GetKitchenRestaurants(kitchen string) ([]models.Restaurant, error) {
	query := `
select * from restaurants
         where kitchen = $1
LIMIT 10;
`
	return r.scanRestaurants(query, kitchen)
}

func (r *Repository) GetTable(tableID string) (models.Table, error) {
	query := `
select * from tables where id = $1;
`
	var table models.Table
	err := r.db.Get(&table, query, tableID)
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
(table_id, profile_id, reservation_date, cells, comment, num_of_guests) 
VALUES 
($1, $2, $3, $4, $5, $6)
returning id
`
	var insertedID int
	err := r.db.Get(&insertedID, query,
		reservation.TableID,
		reservation.ProfileID,
		reservation.ReservationDate,
		reservation.Cells,
		reservation.Comment,
		reservation.NumOfGuests,
	)
	if err != nil {
		return models.Reservation{}, err
	}
	query = `
select * from reservations where id = $1;
`
	err = r.db.Get(&reservation, query, insertedID)
	return reservation, err
}

//MVP2//
/*
func (r *Repository) GetTableReservations(tableID string) ([]models.Reservation, error) {
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
*/

func (r *Repository) GetProfileReservations(profileID string) ([]models.ProfileReservation, error) {
	query := `
select restaurant.*, reservation.* from 
                      reservations reservation
                          join tables t on reservation.table_id = t.id 
                          join restaurants restaurant on t.restaurant_id = restaurant.id 
                  where reservation.profile_id = $1;
`
	rows, err := r.db.Queryx(query, profileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var reservations []models.ProfileReservation
	for rows.Next() {
		var t models.ProfileReservation
		err := rows.Scan(
			&t.Restaurant.ID,
			&t.Restaurant.GoogleID,
			&t.Restaurant.Name,
			&t.Restaurant.Description,
			&t.Restaurant.Address,
			&t.Restaurant.ImgUrl,
			&t.Restaurant.PhoneNumber,
			&t.Restaurant.Email,
			&t.Restaurant.WebsiteUrl,
			&t.Restaurant.Kitchen,
			&t.Restaurant.Tags,
			&t.Restaurant.Rating,
			&t.Restaurant.StartsAtCellID,
			&t.Restaurant.EndsAtCellID,
			&t.Restaurant.Date,
			&t.Restaurant.Lat,
			&t.Restaurant.Lng,
			&t.Reservation.ID,
			&t.Reservation.TableID,
			&t.Reservation.ProfileID,
			&t.Reservation.ReservationDate,
			&t.Reservation.Cells,
			&t.Reservation.Comment,
			&t.Reservation.NumOfGuests,
		)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, t)
	}
	return reservations, nil
}

func (r *Repository) GetRestaurantReservations(restaurantID, date string, numOfGuests string) ([]models.TableAndReservations, error) {
	query := `
select table_id, array_remove(array_agg(reserved_cells order by reserved_cells), null) as reserved_cells
from (
         select distinct tb.id as table_id, unnest(rs.cells)::text as reserved_cells
         from tables tb join reservations rs on tb.id = rs.table_id
         where tb.restaurant_id = $1 and rs.reservation_date = $2 and tb.places = $3
         union
         select distinct tb.id as table_id, null as reserved_cells
         from tables tb left join reservations rs on tb.id = rs.table_id
         where tb.restaurant_id = $1 and tb.places = $3
     ) as s
group by table_id;
`
	rows, err := r.db.Queryx(query, restaurantID, date, numOfGuests)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var reservations []models.TableAndReservations
	for rows.Next() {
		var trs models.TableAndReservations
		err := rows.StructScan(&trs)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, trs)
	}
	return reservations, nil
}

func (r *Repository) Subscribe(userID, restID int) error {
	query := `insert into "favourites" (profile_id, restaurant_id) values ($1,$2);`
	_, err := r.db.Queryx(query, userID, restID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Unsubscribe(userID, restID int) error {
	query := `delete from "favourites" where profile_id = $1 and restaurant_id = $2;`
	_, err := r.db.Queryx(query, userID, restID)
	if err != nil {
		return err
	}
	return nil
}
