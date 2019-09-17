package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Keg Struct (Model)
type Keg struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Brand          *Brand `json:"brand"`
	AlcoholContent int    `json:"alcoholContent"`
	Variety        string `json:"variety"`
	Price          int    `json:"price"`
	Fill           int    `json:"fill"`
}

// Brand Struct
type Brand struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

// init kegs as a slice Keg struct
var kegs []Keg

func getKegs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kegs)
}

func getKeg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// loops through all books and return one with matching id
	for _, item := range kegs {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Keg{})
}

func createKeg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var keg Keg
	_ = json.NewDecoder(r.Body).Decode(&keg)
	keg.ID = strconv.Itoa(rand.Intn(10000000)) // MockId
	kegs = append(kegs, keg)
	json.NewEncoder(w).Encode(keg)
}

func updateKeg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range kegs {
		if item.ID == params["id"] {
			kegs = append(kegs[:index], kegs[index+1:]...)
			var keg Keg
			_ = json.NewDecoder(r.Body).Decode(&keg)
			keg.ID = params["id"]
			kegs = append(kegs, keg)
			json.NewEncoder(w).Encode(keg)
		}
	}
}

func deleteKeg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range kegs {
		if item.ID == params["id"] {
			kegs = append(kegs[:index], kegs[index+1:]...)
			break

		}
	}
	json.NewEncoder(w).Encode(kegs)
}

func main() {
	// Init router
	r := mux.NewRouter()

	// mock data
	kegs = append(kegs, Keg{
		ID:             "1",
		Name:           "Hoppathon",
		Brand:          &Brand{Name: "Brew Hop", Location: "Staya"},
		Price:          5,
		AlcoholContent: 5,
		Variety:        "IPA",
		Fill:           124,
	})

	kegs = append(kegs, Keg{
		ID:             "2",
		Name:           "SCOBY DO",
		Brand:          &Brand{Name: "Symbiote", Location: "Backyard"},
		Price:          5,
		AlcoholContent: 8,
		Variety:        "Kombucha",
		Fill:           124,
	})

	//Route Handlers / Endpoints
	r.HandleFunc("/api/kegs", getKegs).Methods("GET")
	r.HandleFunc("/api/kegs/{id}", getKeg).Methods("GET")
	r.HandleFunc("/api/kegs", createKeg).Methods("POST")
	r.HandleFunc("/api/kegs/{id}", updateKeg).Methods("PUT")
	r.HandleFunc("/api/kegs/{id}", deleteKeg).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
