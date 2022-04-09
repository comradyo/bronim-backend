package delivery

import (
	"bronim/pkg/models"
	"bronim/pkg/utils"
	"bronim/service"
	"github.com/gorilla/mux"
	"net/http"
)

func errBytes(err error) []byte {
	return []byte(err.Error())
}

type Delivery struct {
	repository service.Repository
}

func NewDelivery(repository service.Repository) *Delivery {
	return &Delivery{
		repository: repository,
	}
}

func (h *Delivery) CreateProfile(w http.ResponseWriter, r *http.Request) {
	profile := models.Profile{}
	err := utils.GetObjectFromRequest(r.Body, profile)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	createdProfile, err := h.repository.CreateProfile(profile)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	body, err := utils.Marshall(createdProfile)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	profileID := vars["profile"]
	profile, err := h.repository.GetProfile(profileID)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	body, err := utils.Marshall(profile)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	profileID := vars["profile"]
	profile := models.Profile{}
	err := utils.GetObjectFromRequest(r.Body, profile)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	updatedProfile, err := h.repository.UpdateProfile(profileID, profile)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	body, err := utils.Marshall(updatedProfile)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, nil)
		return
	}
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurant := models.Restaurant{}
	err := utils.GetObjectFromRequest(r.Body, restaurant)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	createdRestaurant, err := h.repository.CreateRestaurant(restaurant)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	body, err := utils.Marshall(createdRestaurant)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetRestaurant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	restaurantID := vars["restaurant"]
	restaurant, err := h.repository.GetRestaurant(restaurantID)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	body, err := utils.Marshall(restaurant)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetTables(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	restaurantID := vars["restaurant"]
	tables, err := h.repository.GetTables(restaurantID)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	tablesList := models.Tables{
		Arr: tables,
	}
	body, err := utils.Marshall(tablesList)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) CreateReservation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	restaurantID := vars["restaurant"]
	tableID := vars["table"]
	reservation := models.Reservation{}
	err := utils.GetObjectFromRequest(r.Body, reservation)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	createdReservation, err := h.repository.CreateReservation(restaurantID, tableID, reservation)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	body, err := utils.Marshall(createdReservation)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetReservations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	restaurantID := vars["restaurant"]
	tableID := vars["table"]
	reservations, err := h.repository.GetReservations(restaurantID, tableID)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	reservationsList := models.Reservations{
		Arr: reservations,
	}
	body, err := utils.Marshall(reservationsList)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetProfileReservations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	profileID := vars["profile"]
	profileReservations, err := h.repository.GetProfileReservations(profileID)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	profileReservationsList := models.ProfileReservations{
		Arr: profileReservations,
	}
	body, err := utils.Marshall(profileReservationsList)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	utils.SendResponse(w, http.StatusOK, body)
}
