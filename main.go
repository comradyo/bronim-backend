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
	"os"
)

func main() {
	log.Init(log.DebugLevel)

	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
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
		os.Exit(1)
	}

}

func NewDB() (*sql.DB, error) {
	user := "postgres"
	password := "yo_password"
	//host := viper.GetString("postgres_db.host")
	//port := viper.GetString("postgres_db.port")
	host := "194.87.111.83"
	port := "5432"
	dbname := "bronim"
	sslmode := "disable"
	//connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)
	return sql.Connect("postgres", connStr)
}

func setRouter(delivery *delivery.Delivery) *mux.Router {
	r := mux.NewRouter().PathPrefix("/bronim").Subrouter()

	r.HandleFunc("/profiles", delivery.CreateProfile).Methods("POST")
	r.HandleFunc("/profiles/{uuid}", delivery.GetProfile).Methods("GET")
	r.HandleFunc("/profiles/{uuid}", delivery.UpdateProfile).Methods("POST")
	r.HandleFunc("/profiles/{uuid}/favourites", delivery.GetFavourites).Methods("GET")
	r.HandleFunc("/profiles/{uuid}/subscribe/{restid}", delivery.Subscribe).Methods("GET")
	r.HandleFunc("/profiles/{uuid}/unsubscribe/{restid}", delivery.Unsubscribe).Methods("GET")

	r.HandleFunc("/restaurants", delivery.CreateRestaurant).Methods("POST")
	r.HandleFunc("/restaurants/{restaurant:[0-9]+}", delivery.GetRestaurant).Methods("GET")

	r.HandleFunc("/restaurants", delivery.GetRestaurants).Methods("GET")
	r.HandleFunc("/restaurants/popular", delivery.GetPopularRestaurants).Methods("GET")
	r.HandleFunc("/restaurants/nearest", delivery.GetNearestRestaurants).Methods("GET")
	r.HandleFunc("/restaurants/new", delivery.GetNewRestaurants).Methods("GET")
	r.HandleFunc("/cuisines/{cuisine}", delivery.GetKitchenRestaurants).Methods("GET")

	r.HandleFunc("/restaurants/{restaurant:[0-9]+}/reservations", delivery.GetRestaurantReservations).Methods("GET")
	r.HandleFunc("/restaurants/{restaurant:[0-9]+}/tables/{table:[0-9]+}", delivery.CreateReservation).Methods("POST")

	r.HandleFunc("/profiles/{uuid}/reservations", delivery.GetProfileReservations).Methods("GET")

	//TODO: Vote

	return r
}
