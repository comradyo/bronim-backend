package delivery

import (
	log "bronim/pkg/logger"
	"bronim/pkg/models"
	"bronim/pkg/places"
	"bronim/pkg/utils"
	"bronim/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func errBytes(err error) []byte {
	bytes, _ := utils.Marshall(models.Err{ErrStr: err.Error()})
	return bytes
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
	log.InfoAtFunc(h.CreateProfile, "started")
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
	log.InfoAtFunc(h.CreateProfile, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetProfile(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.GetProfile, "started")
	vars := mux.Vars(r)
	profileID := vars["uuid"]
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
	log.InfoAtFunc(h.GetProfile, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.UpdateProfile, "started")
	vars := mux.Vars(r)
	profileID := vars["uuid"]
	profile := models.Profile{}
	err := utils.GetObjectFromRequest(r.Body, &profile)
	if err != nil {
		log.ErrorAtFunc(h.UpdateProfile, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	updatedProfile, err := h.repository.UpdateProfile(profileID, profile)
	if err != nil {
		log.ErrorAtFunc(h.UpdateProfile, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	body, err := utils.Marshall(updatedProfile)
	if err != nil {
		log.ErrorAtFunc(h.UpdateProfile, err)
		utils.SendResponse(w, http.StatusInternalServerError, nil)
		return
	}
	log.InfoAtFunc(h.UpdateProfile, "started")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.CreateRestaurant, "started")
	restaurant := models.Restaurant{}
	err := utils.GetObjectFromRequest(r.Body, &restaurant)
	if err != nil {
		log.ErrorAtFunc(h.CreateRestaurant, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	fmt.Println(restaurant)
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
	log.InfoAtFunc(h.CreateRestaurant, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetRestaurant(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.GetRestaurant, "started")
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
	log.InfoAtFunc(h.GetRestaurant, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetRestaurants(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.GetRestaurants, "started")

	q := r.URL.Query()
	var cuisine string
	if len(q["cuisine"]) > 0 {
		cuisine = q["cuisine"][0]
	}

	/*
		filter := service.GetRestaurantsFilter{
			Cuisine: cuisine,
		}
	*/

	rests, err := h.repository.GetKitchenRestaurants(cuisine)
	if err != nil {
		log.ErrorAtFunc(h.GetRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	restaurants := models.RestaurantList{
		Arr: rests,
	}
	body, err := utils.Marshall(restaurants)
	if err != nil {
		log.ErrorAtFunc(h.GetRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}

	log.InfoAtFunc(h.GetRestaurants, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetPopularRestaurants(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.GetPopularRestaurants, "started")
	rests, err := h.repository.GetPopularRestaurants()
	if err != nil {
		log.ErrorAtFunc(h.GetPopularRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	restaurants := models.RestaurantList{
		Arr: rests,
	}
	body, err := utils.Marshall(restaurants)
	if err != nil {
		log.ErrorAtFunc(h.GetPopularRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.InfoAtFunc(h.GetPopularRestaurants, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

//TODO: работа с гугловской апишкой
//В деливери идем на GoogleAPI с координатами, полученными из запроса, берем айдишники близжайших ресторанов,
func (h *Delivery) GetNearestRestaurants(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.GetNearestRestaurants, "started")
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
	restaurants := models.RestaurantList{
		Arr: rests,
	}
	body, err := utils.Marshall(restaurants)
	if err != nil {
		log.ErrorAtFunc(h.GetNearestRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.InfoAtFunc(h.GetNearestRestaurants, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetNewRestaurants(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.GetNewRestaurants, "started")
	rests, err := h.repository.GetNewRestaurants()
	if err != nil {
		log.ErrorAtFunc(h.GetNewRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	restaurants := models.RestaurantList{
		Arr: rests,
	}
	body, err := utils.Marshall(restaurants)
	if err != nil {
		log.ErrorAtFunc(h.GetNewRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.InfoAtFunc(h.GetProfileReservations, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetKitchenRestaurants(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.GetKitchenRestaurants, "started")
	vars := mux.Vars(r)
	kitchen := vars["cuisine"]
	rests, err := h.repository.GetKitchenRestaurants(kitchen)
	if err != nil {
		log.ErrorAtFunc(h.GetKitchenRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	restaurants := models.RestaurantList{
		Arr: rests,
	}
	body, err := utils.Marshall(restaurants)
	if err != nil {
		log.ErrorAtFunc(h.GetKitchenRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.InfoAtFunc(h.GetKitchenRestaurants, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetRestaurantReservations(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.GetRestaurantReservations, "started")
	vars := mux.Vars(r)
	restaurantID := vars["restaurant"]
	query := r.URL.Query()

	var date string
	var numOfGuests string
	if len(query["date"]) > 0 {
		date = query["date"][0]
	} else {
		err := models.ErrEmptyValue("date")
		log.ErrorAtFunc(h.GetRestaurantReservations, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	if len(query["guests"]) > 0 {
		numOfGuests = query["guests"][0]
	} else {
		err := models.ErrEmptyValue("guests")
		log.ErrorAtFunc(h.GetRestaurantReservations, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}

	reservations, err := h.repository.GetRestaurantReservations(restaurantID, date, numOfGuests)
	if err != nil {
		log.ErrorAtFunc(h.GetRestaurantReservations, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	for i := 0; i < len(reservations); i++ {
		if len(reservations[i].ReservedTimes) == 48 {
			reservations[i] = reservations[len(reservations)-1]
			reservations = reservations[:len(reservations)-1]
			i--
		}
	}
	reservationsList := models.TableAndReservationsList{
		Arr: reservations,
	}
	body, err := utils.Marshall(reservationsList)
	if err != nil {
		log.ErrorAtFunc(h.GetRestaurantReservations, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.InfoAtFunc(h.GetRestaurantReservations, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) GetFavourites(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.GetFavourites, "started")
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	user, err := h.repository.GetProfile(uuid)
	if err != nil {
		log.ErrorAtFunc(h.GetRestaurantReservations, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	userId, _ := strconv.Atoi(user.ID)
	rests, err := h.repository.GetFavouritesRestaurants(userId)
	if err != nil {
		log.ErrorAtFunc(h.GetRestaurantReservations, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	restaurants := models.RestaurantList{
		Arr: rests,
	}
	body, err := utils.Marshall(restaurants)
	if err != nil {
		log.ErrorAtFunc(h.GetNewRestaurants, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.InfoAtFunc(h.GetFavourites, "ended")
	utils.SendResponse(w, http.StatusOK, body)

}

//MVP2//
/*
func (h *Delivery) GetTables(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.GetTables, "started")
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
	log.InfoAtFunc(h.GetTables, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}
*/

func (h *Delivery) CreateReservation(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.CreateReservation, "started")
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
	log.InfoAtFunc(h.CreateReservation, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

//MVP2//
/*
func (h *Delivery) GetTableReservations(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.GetTableReservations, "started")
	vars := mux.Vars(r)
	restaurantID := vars["restaurant"]
	tableID := vars["table"]

	table, err := h.repository.GetTable(tableID)
	if err != nil {
		log.ErrorAtFunc(h.GetTableReservations, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	if table.RestaurantID != restaurantID {
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(
			fmt.Errorf("table %s is not in restaurant %s", tableID, restaurantID),
		))
		return
	}

	reservations, err := h.repository.GetTableReservations(tableID)
	if err != nil {
		log.ErrorAtFunc(h.GetTableReservations, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	reservationsList := models.Reservations{
		Arr: reservations,
	}
	body, err := utils.Marshall(reservationsList)
	if err != nil {
		log.ErrorAtFunc(h.GetTableReservations, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.InfoAtFunc(h.GetTableReservations, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}
*/

func (h *Delivery) GetProfileReservations(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.GetProfileReservations, "started")
	vars := mux.Vars(r)
	profileID := vars["uuid"]
	profileReservations, err := h.repository.GetProfileReservations(profileID)
	if err != nil {
		log.ErrorAtFunc(h.GetProfileReservations, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	profileReservationsList := models.ProfileReservationList{
		Arr: profileReservations,
	}
	fmt.Printf("reservations = %v\n", profileReservations)
	body, err := utils.Marshall(profileReservationsList)
	if err != nil {
		log.ErrorAtFunc(h.GetProfileReservations, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.InfoAtFunc(h.GetProfileReservations, "ended")
	utils.SendResponse(w, http.StatusOK, body)
}

func (h *Delivery) Subscribe(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.Subscribe, "started")
	vars := mux.Vars(r)
	profileID := vars["uuid"]
	restaurantID := vars["restid"]
	user, err := h.repository.GetProfile(profileID)
	if err != nil {
		log.ErrorAtFunc(h.Subscribe, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	userID, _ := strconv.Atoi(user.ID)
	restID, _ := strconv.Atoi(restaurantID)
	err = h.repository.Subscribe(userID, restID)
	if err != nil {
		log.ErrorAtFunc(h.Subscribe, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.InfoAtFunc(h.Subscribe, "ended")
	message, err := utils.Marshall(models.OK{OkMessage: "OK"})
	if err != nil {
		log.ErrorAtFunc(h.Subscribe, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	utils.SendResponse(w, http.StatusOK, message)
}

func (h *Delivery) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	log.InfoAtFunc(h.Unsubscribe, "started")
	vars := mux.Vars(r)
	profileID := vars["uuid"]
	restaurantID := vars["restid"]
	user, err := h.repository.GetProfile(profileID)
	if err != nil {
		log.ErrorAtFunc(h.Unsubscribe, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	userID, _ := strconv.Atoi(user.ID)
	restID, _ := strconv.Atoi(restaurantID)
	err = h.repository.Unsubscribe(userID, restID)
	if err != nil {
		log.ErrorAtFunc(h.Unsubscribe, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	log.InfoAtFunc(h.Unsubscribe, "ended")
	message, err := utils.Marshall(models.OK{OkMessage: "OK"})
	if err != nil {
		log.ErrorAtFunc(h.Unsubscribe, err)
		utils.SendResponse(w, http.StatusInternalServerError, errBytes(err))
		return
	}
	utils.SendResponse(w, http.StatusOK, message)
}
