package main

import (
	log "bronim/pkg/logger"
	"bronim/pkg/places"
	"bronim/service/delivery"
	"bronim/service/repository"
	"fmt"
	"github.com/gorilla/mux"
	sql "github.com/jmoiron/sqlx"
	"net/http"
)

func main() {
	log.Init(log.DebugLevel)

	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	//TODO: credentials из конфига
	googlePlacesClient := places.NewGooglePlacesClient("")
	deli := delivery.NewDelivery(repo, *googlePlacesClient)
	router := setRouter(deli)
	port := "5000"

	log.InfoAtFunc(main, "bronim started", "YO")
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("err = ", err)
	}

}

func NewDB() (*sql.DB, error) {
	user := "postgres"
	password := "password"
	//host := viper.GetString("postgres_db.host")
	//port := viper.GetString("postgres_db.port")
	dbname := "postgres"
	sslmode := "disable"
	//connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", user, dbname, password, sslmode)
	return sql.Connect("postgres", connStr)
}

func setRouter(delivery *delivery.Delivery) *mux.Router {
	r := mux.NewRouter().PathPrefix("/bronim").Subrouter()

	r.HandleFunc("/profile", delivery.CreateProfile).Methods("POST")
	r.HandleFunc("/profile/{profile:[0-9]+}", delivery.GetProfile).Methods("GET")
	//MVP2// r.HandleFunc("/profile/{profile:[0-9]+}", delivery.UpdateProfile).Methods("POST")

	r.HandleFunc("/restaurant", delivery.CreateRestaurant).Methods("POST")
	r.HandleFunc("/restaurant/{restaurant:[0-9]+}", delivery.GetRestaurant).Methods("GET")

	r.HandleFunc("/restaurants/popular", delivery.GetPopularRestaurants).Methods("GET")
	r.HandleFunc("/restaurants/nearest", delivery.GetNearestRestaurants).Methods("GET")
	r.HandleFunc("/restaurants/new", delivery.GetNewRestaurants).Methods("GET")
	r.HandleFunc("/kitchens/{kitchen}", delivery.GetKitchenRestaurants).Methods("GET")

	r.HandleFunc("/restaurant/{restaurant:[0-9]+}/tables", delivery.GetTables).Methods("GET")

	r.HandleFunc("/restaurant/{restaurant:[0-9]+}/tables/{table:[0-9]+}", delivery.CreateReservation).Methods("POST")
	r.HandleFunc("/restaurant/{restaurant:[0-9]+}/tables/{table:[0-9]+}", delivery.GetReservations).Methods("GET")

	r.HandleFunc("/profile/{profile:[0-9]+}/reservations", delivery.GetProfileReservations).Methods("GET")

	//TODO: Vote

	return r
}
