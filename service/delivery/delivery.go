package delivery

import (
	log "bronim/pkg/logger"
	"bronim/pkg/models"
	"bronim/pkg/places"
	"bronim/pkg/utils"
	"bronim/service"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func errBytes(err error) []byte {
	return []byte(err.Error())
}

type Delivery struct {
	repository         service.Repository
	googlePlacesClient places.GooglePlacesClient
}

func NewDelivery(repository service.Repository, googlePlacesClient places.GooglePlacesClient) *Delivery {
	return &Delivery{
		repository:         repository,
		googlePlacesClient: googlePlacesClient,
	}
}

func (h *Delivery) CreateProfile(w http.ResponseWriter, r *http.Request) {
	log.DebugAtFunc(h.CreateProfile, "started")
	profile := models.Profile{}
	err := utils.GetObjectFromRequest(r.Body, &profile)
	if err != nil {
		log.ErrorAtFunc(h.CreateProfile, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	createdProfile, err := h.repository.CreateProfile(profile)
	if err != nil {
		log.ErrorAtFunc(h.CreateProfile, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	body, err := utils.Marshall(createdProfile)
	if err != nil {
		log.ErrorAtFunc(h.CreateProfile, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.DebugAtFunc(h.CreateProfile, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetProfile(w http.ResponseWriter, r *http.Request) {
	log.DebugAtFunc(h.GetProfile, "started")
	vars := mux.Vars(r)
	profileID := vars["profile"]
	profile, err := h.repository.GetProfile(profileID)
	if err != nil {
		log.ErrorAtFunc(h.GetProfile, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	body, err := utils.Marshall(profile)
	if err != nil {
		log.ErrorAtFunc(h.GetProfile, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.DebugAtFunc(h.GetProfile, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

//MVP2//
/*
func (h *Delivery) UpdateProfile(w http.ResponseWriter, r *http.Request) { log.DebugAtFunc(h.CreateProfile, "started")
	vars := mux.Vars(r)
	profileID := vars["profile"]
	profile := models.Profile{}
	err := utils.GetObjectFromRequest(r.Body, profile)
	if err != nil { log.ErrorAtFunc(h.CreateProfile, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	updatedProfile, err := h.repository.UpdateProfile(profileID, profile)
	if err != nil { log.ErrorAtFunc(h.CreateProfile, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	body, err := utils.Marshall(updatedProfile)
	if err != nil { log.ErrorAtFunc(h.CreateProfile, err)
		utils.SendResponse(w, http.StatusInternalServerError, nil)
		return
	}
	log.DebugAtFunc(h.GetProfileReservations, "started") utils.SendResponse(w, http.StatusOK, body)
}
*/

func (h *Delivery) CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	log.DebugAtFunc(h.CreateRestaurant, "started")
	restaurant := models.Restaurant{}
	err := utils.GetObjectFromRequest(r.Body, &restaurant)
	if err != nil {
		log.ErrorAtFunc(h.CreateRestaurant, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	createdRestaurant, err := h.repository.CreateRestaurant(restaurant)
	if err != nil {
		log.ErrorAtFunc(h.CreateRestaurant, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	body, err := utils.Marshall(createdRestaurant)
	if err != nil {
		log.ErrorAtFunc(h.CreateRestaurant, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.DebugAtFunc(h.CreateRestaurant, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetRestaurant(w http.ResponseWriter, r *http.Request) {
	log.DebugAtFunc(h.GetRestaurant, "started")
	vars := mux.Vars(r)
	restaurantID := vars["restaurant"]
	restaurant, err := h.repository.GetRestaurant(restaurantID)
	if err != nil {
		log.ErrorAtFunc(h.GetRestaurant, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	body, err := utils.Marshall(restaurant)
	if err != nil {
		log.ErrorAtFunc(h.GetRestaurant, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.DebugAtFunc(h.GetRestaurant, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

//TODO: популярность
func (h *Delivery) GetPopularRestaurants(w http.ResponseWriter, r *http.Request) {
	log.DebugAtFunc(h.GetPopularRestaurants, "started")
	rests, err := h.repository.GetPopularRestaurants()
	if err != nil {
		log.ErrorAtFunc(h.GetPopularRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	restaurants := models.Restaurants{
		Arr: rests,
	}
	body, err := utils.Marshall(restaurants)
	if err != nil {
		log.ErrorAtFunc(h.GetPopularRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.DebugAtFunc(h.GetPopularRestaurants, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

//TODO: работа с гугловской апишкой
//В деливери идем на GoogleAPI с координатами, полученными из запроса, берем айдишники близжайших ресторанов,
func (h *Delivery) GetNearestRestaurants(w http.ResponseWriter, r *http.Request) {
	log.DebugAtFunc(h.GetNearestRestaurants, "started")
	q := r.URL.Query()

	var lat string
	var lon string

	if len(q["lat"]) > 0 {
		lat = q["lat"][0]
	}
	if len(q["lon"]) > 0 {
		lon = q["lon"][0]
	}

	apiRestaurants, err := h.googlePlacesClient.GetNearestRestaurants(lat, lon)

	rests, err := h.repository.GetNearestRestaurants(apiRestaurants)
	if err != nil {
		log.ErrorAtFunc(h.GetNearestRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	restaurants := models.Restaurants{
		Arr: rests,
	}
	body, err := utils.Marshall(restaurants)
	if err != nil {
		log.ErrorAtFunc(h.GetNearestRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.DebugAtFunc(h.GetNearestRestaurants, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetNewRestaurants(w http.ResponseWriter, r *http.Request) {
	log.DebugAtFunc(h.GetNewRestaurants, "started")
	rests, err := h.repository.GetNewRestaurants()
	if err != nil {
		log.ErrorAtFunc(h.GetNewRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	restaurants := models.Restaurants{
		Arr: rests,
	}
	body, err := utils.Marshall(restaurants)
	if err != nil {
		log.ErrorAtFunc(h.GetNewRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.DebugAtFunc(h.GetProfileReservations, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetKitchenRestaurants(w http.ResponseWriter, r *http.Request) {
	log.DebugAtFunc(h.GetKitchenRestaurants, "started")
	vars := mux.Vars(r)
	kitchen := vars["kitchen"]
	rests, err := h.repository.GetKitchenRestaurants(kitchen)
	if err != nil {
		log.ErrorAtFunc(h.GetKitchenRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	restaurants := models.Restaurants{
		Arr: rests,
	}
	body, err := utils.Marshall(restaurants)
	if err != nil {
		log.ErrorAtFunc(h.GetKitchenRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.DebugAtFunc(h.GetKitchenRestaurants, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetTables(w http.ResponseWriter, r *http.Request) {
	log.DebugAtFunc(h.GetTables, "started")
	vars := mux.Vars(r)
	restaurantID := vars["restaurant"]
	tables, err := h.repository.GetTables(restaurantID)
	if err != nil {
		log.ErrorAtFunc(h.GetTables, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	tablesList := models.Tables{
		Arr: tables,
	}
	body, err := utils.Marshall(tablesList)
	if err != nil {
		log.ErrorAtFunc(h.GetTables, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.DebugAtFunc(h.GetTables, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) CreateReservation(w http.ResponseWriter, r *http.Request) {
	log.DebugAtFunc(h.CreateReservation, "started")
	vars := mux.Vars(r)
	restaurantID := vars["restaurant"]
	tableID := vars["table"]
	reservation := models.Reservation{}
	err := utils.GetObjectFromRequest(r.Body, &reservation)
	if err != nil {
		log.ErrorAtFunc(h.CreateReservation, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}

	table, err := h.repository.GetTable(tableID)
	if err != nil {
		log.ErrorAtFunc(h.CreateReservation, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	if table.RestaurantID != restaurantID {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(
			fmt.Errorf("table %s is not in restaurant %s", tableID, restaurantID),
		))
		return
	}

	reservation.TableID = tableID
	createdReservation, err := h.repository.CreateReservation(reservation)
	if err != nil {
		log.ErrorAtFunc(h.CreateReservation, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	body, err := utils.Marshall(createdReservation)
	if err != nil {
		log.ErrorAtFunc(h.CreateReservation, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.DebugAtFunc(h.CreateReservation, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetReservations(w http.ResponseWriter, r *http.Request) {
	log.DebugAtFunc(h.GetReservations, "started")
	vars := mux.Vars(r)
	restaurantID := vars["restaurant"]
	tableID := vars["table"]

	table, err := h.repository.GetTable(tableID)
	if err != nil {
		log.ErrorAtFunc(h.GetReservations, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	if table.RestaurantID != restaurantID {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(
			fmt.Errorf("table %s is not in restaurant %s", tableID, restaurantID),
		))
		return
	}

	reservations, err := h.repository.GetReservations(tableID)
	if err != nil {
		log.ErrorAtFunc(h.GetReservations, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	reservationsList := models.Reservations{
		Arr: reservations,
	}
	body, err := utils.Marshall(reservationsList)
	if err != nil {
		log.ErrorAtFunc(h.GetReservations, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.DebugAtFunc(h.GetReservations, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetProfileReservations(w http.ResponseWriter, r *http.Request) {
	log.DebugAtFunc(h.GetProfileReservations, "started")
	vars := mux.Vars(r)
	profileID := vars["profile"]
	profileReservations, err := h.repository.GetProfileReservations(profileID)
	if err != nil {
		log.ErrorAtFunc(h.GetProfileReservations, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	profileReservationsList := models.ProfileReservations{
		Arr: profileReservations,
	}
	body, err := utils.Marshall(profileReservationsList)
	if err != nil {
		log.ErrorAtFunc(h.GetProfileReservations, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.DebugAtFunc(h.GetProfileReservations, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}