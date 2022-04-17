package repository

import (
	"bronim/pkg/models"
	sql "github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
	var insertedID string
	err := r.db.Get(&insertedID, query,
		profile.FirebaseID,
		profile.Name,
		profile.Surname,
		profile.DateOfBirth,
		profile.Sex,
		profile.PhoneNumber,
		profile.Email,
		profile.Password,
		profile.AvatarUrl,
	)
	if err != nil {
		return models.Profile{}, err
	}
	return r.GetProfile(insertedID)
}

func (r *Repository) GetProfile(profileID string) (models.Profile, error) {
	query := `
select * from profiles where id = $1
`
	var profile models.Profile
	err := r.db.Get(&profile, query, profileID)
	return profile, err
}

//MVP2//
/*
func (r *Repository) UpdateProfile(profile models.Profile) (models.Profile, error) {
	query := `
update profiles set name = $1, surname = $2, date_of_birth = $3, sex = $4, phone_number = $5, email = $6, avatar_url = $7
where id = $8
returning id;
`
	var updatedID int
	err := r.db.Get(&updatedID, query, profile)
	if err != nil {
		return models.Profile{}, err
	}
	query = `
select * from profiles where id = $1;
`
	err = r.db.Get(&profile, query, updatedID)
	return profile, err
}
*/

func (r *Repository) CreateRestaurant(restaurant models.Restaurant) (models.Restaurant, error) {
	query := `
insert into restaurants (google_id, address, description, img_url, phone_number, email, website_url, geoposition, kitchen, tags, rating)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10::text[], $11)
returning (id)
`
	var insertedID string
	pqRest := toPqRestaurant(restaurant)
	err := r.db.Get(&insertedID, query,
		pqRest.GoogleID,
		pqRest.Address,
		pqRest.Description,
		pqRest.ImgUrl,
		pqRest.PhoneNumber,
		pqRest.Email,
		pqRest.WebsiteUrl,
		pqRest.Geoposition,
		pqRest.Kitchen,
		pqRest.Tags,
		pqRest.Rating,
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
	var pqRest pqRestaurant
	err := r.db.Get(&pqRest, query, restaurantID)
	restaurant := toModelRestaurant(pqRest)
	return restaurant, err
}

//TODO: популярность
func (r *Repository) GetPopularRestaurants() ([]models.Restaurant, error) {
	query := `
select * from restaurants
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
select * from restaurants
LIMIT 10;
`
	return r.scanRestaurants(query)
}

func (r *Repository) GetNewRestaurants() ([]models.Restaurant, error) {
	query := `
select * from restaurants
         order by id desc
LIMIT 10;
`
	return r.scanRestaurants(query)
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
(table_id, profile_id, reservation_date, cell_id, num_of_cells, comment) 
VALUES 
($1, $2, $3, $4, $5, $6)
returning id
`
	var insertedID int
	err := r.db.Get(&insertedID, query, reservation)
	if err != nil {
		return models.Reservation{}, err
	}
	query = `
select * from reservations where id = $1;
`
	err = r.db.Get(&reservation, query, insertedID)
	return reservation, err
}

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

func (r *Repository) GetRestaurantReservations(restaurantID, date string, numOfGuests string) ([]models.TableAndReservations, error) {
	query := `
select tb.id as table_id, array_agg(rs.cell_id) as reserved_cells, array_agg(rs.num_of_cells) as num_of_cells from
             reservations rs
                 join
                 tables tb
                     on rs.table_id = tb.id
                 join restaurants rt
                     on rt.id = tb.restaurant_id
                where tb.restaurant_id = $1 and rs.reservation_date = $2 and tb.places = $3
group by tb.id;
`
	rows, err := r.db.Queryx(query, restaurantID, date, numOfGuests)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var reservations []models.TableAndReservations
	for rows.Next() {
		var rs restaurantReservation
		err := rows.StructScan(&rs)
		if err != nil {
			return nil, err
		}
		trs := toTableAndReservation(rs)
		reservations = append(reservations, trs)
	}
	return reservations, nil
}

type restaurantReservation struct {
	TableID    string        `db:"table_id"`
	CellIDs    pq.Int64Array `db:"reserved_cells"`
	NumOfCells pq.Int64Array `db:"num_of_cells"`
}

func toTableAndReservation(rs restaurantReservation) models.TableAndReservations {
	var trs models.TableAndReservations
	trs.TableID = rs.TableID
	for i := range rs.CellIDs {
		var cellIDs []int
		numOfCells := int(rs.NumOfCells[i])
		for j := 0; j < numOfCells; j++ {
			cellIDs = append(cellIDs, int(rs.CellIDs[i])+j)
		}
		trs.ReservedTimes = append(trs.ReservedTimes, cellIDs...)
	}
	return trs
}
