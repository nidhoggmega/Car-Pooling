// main.go (Microservice 1: User and Trip)
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// User represents a user in the car-pooling platform
type User struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Mobile      string    `json:"mobile"`
	Email       string    `json:"email"`
	Driver      bool      `json:"driver"`
	License     string    `json:"license,omitempty"`
	CarPlate    string    `json:"car_plate,omitempty"`
	AccountDate time.Time `json:"account_date"`
}

// Trip represents a car-pooling trip
type Trip struct {
	ID             string    `json:"id"`
	OwnerID        string    `json:"owner_id"`
	Pickup         string    `json:"pickup"`
	AltPickup      string    `json:"alt_pickup"`
	Destination    string    `json:"destination"`
	StartTime      time.Time `json:"start_time"`
	MaxSeats       int       `json:"max_seats`
	AvailableSeats int       `json:"available_seats"`
}

type Trips struct {
	Users map[string]Trip `json:"Trips"`
}

var users = make(map[string]User)
var trips = make(map[string]Trip)

// var users []User
// var trips []Trip
func main() {

	router := mux.NewRouter()

	// Define API endpoints
	router.HandleFunc("/api/v1/users", CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/v1/trips", CreateTrip).Methods("POST")
	router.HandleFunc("/api/v1/trips/{id}", UpdateTrip).Methods("PUT")
	router.HandleFunc("/trips", GetAllTrips).Methods("GET")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

// CreateUser creates a new user account
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser.ID = generateID()
	newUser.AccountDate = time.Now()

	users[newUser.ID] = newUser

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// GetUser retrieves a user by ID for LogIn&SignUp
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	user, exists := users[userID]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Create user ID
func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	existingUser, found := users[userID]
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var updatedUser User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update user information
	existingUser.FirstName = updatedUser.FirstName
	existingUser.LastName = updatedUser.LastName
	existingUser.Mobile = updatedUser.Mobile
	existingUser.Email = updatedUser.Email

	// If the user is a car owner, update car owner information
	if existingUser.Driver == true {
		existingUser.License = updatedUser.License
		existingUser.CarPlate = updatedUser.CarPlate
	}

	users[userID] = existingUser

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingUser)
}

// Delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	if _, exists := users[userID]; !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Delete user
	delete(users, userID)

	w.WriteHeader(http.StatusNoContent)
}

// CreateTrip creates a new car-pooling trip
func CreateTrip(w http.ResponseWriter, r *http.Request) {
	var newTrip Trip
	_ = json.NewDecoder(r.Body).Decode(&newTrip)

	newTrip.ID = fmt.Sprintf("%d", time.Now().UnixNano())

	trips[newTrip.ID] = newTrip

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTrip)
}

// update Trip Info
func UpdateTrip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tripID := params["id"]

	if _, exists := trips[tripID]; !exists {
		http.Error(w, "Trip not found", http.StatusNotFound)
		return
	}

	var updatedTrip Trip
	_ = json.NewDecoder(r.Body).Decode(&updatedTrip)

	// Update trip information
	trips[tripID] = updatedTrip

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTrip)
}

// GetAllTrips retrieves all trips
func GetAllTrips(w http.ResponseWriter, r *http.Request) {
	var allTrips []Trip
	for _, trip := range trips {
		allTrips = append(allTrips, trip)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allTrips)
}
