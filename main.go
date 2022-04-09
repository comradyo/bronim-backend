package main

import (
	"bronim/service/delivery"
	"fmt"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, World!")
}

func setRouter(delivery *delivery.Delivery) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/profile", delivery.CreateProfile).Methods("POST")
	r.HandleFunc("/profile/{profile:[0-9]+}", delivery.GetProfile).Methods("GET")
	r.HandleFunc("/profile/{profile:[0-9]+}", delivery.UpdateProfile).Methods("POST")

	r.HandleFunc("/restaurant", delivery.CreateRestaurant).Methods("POST")
	r.HandleFunc("/restaurant/{restaurant:[0-9]+}", delivery.GetRestaurant).Methods("GET")

	r.HandleFunc("/restaurant/{restaurant:[0-9]+}/tables", delivery.GetTables).Methods("GET")

	r.HandleFunc("/restaurant/{restaurant:[0-9]+}/tables/{table:[0-9]+}", delivery.CreateReservation).Methods("POST")
	r.HandleFunc("/restaurant/{restaurant:[0-9]+}/tables/{table:[0-9]+}", delivery.GetReservations).Methods("GET")

	r.HandleFunc("/profile/{profile:[0-9]+}/reservations", delivery.GetProfileReservations).Methods("GET")

	//TODO: Vote

	return r
}
