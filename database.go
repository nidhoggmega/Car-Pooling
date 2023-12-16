// database.go (Microservice 2: Database Connection to API)
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("my_db", "car:password@tcp(localhost:3306)/carpooling")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
}

// Sign Up
func saveUserToDB(user User) {
	_, err := db.Exec("INSERT INTO users (id, first_name, last_name, mobile, email, driver, license, car_plate, account_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		user.ID, user.FirstName, user.LastName, user.Mobile, user.Email, user.Driver, user.License, user.CarPlate, user.AccountDate)
	if err != nil {
		log.Fatal(err)
	}
}

// Log In
func getUserFromDB(userID string) User {
	var user User
	err := db.QueryRow("SELECT * FROM users WHERE id=?", userID).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Mobile, &user.Email, &user.Driver, &user.License, &user.CarPlate, &user.AccountDate)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

// Update User
func updateUserToDB(user User) {
	_, err := db.Exec("UPDATE users SET id=?,first_name=?,last_name=?,mobile=?,email=?,driver=?,license=?,car_plate=?,account_date=?",
		user.ID, user.FirstName, user.LastName, user.Mobile, user.Email, user.Driver, user.License, user.CarPlate, user.AccountDate)
	if err != nil {
		log.Fatal(err)
	}
}

// Delete User
func deleteUserFromDB(userID string) {
	_, err := db.Exec("DELETE FROM users WHERE ID=?", userID)
	if err != nil {
		log.Fatal(err)
	}
}

// Create Trip
func saveTripToDB(trip Trip) {
	_, err := db.Exec("INSERT INTO trips (id, car_owner_id, pickup_location, alt_pick_up_location, start_time, destination, max_passengers, seats_available) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		trip.ID, trip.OwnerID, trip.Pickup, trip.AltPickup, trip.StartTime, trip.Destination, trip.MaxSeats, trip.AvailableSeats)
	if err != nil {
		log.Fatal(err)
	}
}

// Update Trip info
func updateTripToDB(trip Trip) {
	_, err := db.Exec("UPDATE trips set id=?,car_owner_id=?,pickup_location=?,alt_pick_up_location=?, start_time=?, destination=?, max_passengers=?, seats_available=?",
		trip.ID, trip.OwnerID, trip.Pickup, trip.AltPickup, trip.StartTime, trip.Destination, trip.MaxSeats, trip.AvailableSeats)
	if err != nil {
		log.Fatal(err)
	}
}

// List all trips
func getAllTrips() map[string]Trip {
	results, err := db.Query("select * from trips")
	if err != nil {
		panic(err.Error())
	}

	var trips map[string]Trip = map[string]Trip{}

	for results.Next() {
		var t Trip
		var id string
		err = results.Scan(&id, &t.OwnerID, &t.Pickup, &t.AltPickup, &t.StartTime, &t.Destination, &t.MaxSeats, &t.AvailableSeats)
		if err != nil {
			panic(err.Error())
		}

		trips[id] = t
	}

	return trips
}
