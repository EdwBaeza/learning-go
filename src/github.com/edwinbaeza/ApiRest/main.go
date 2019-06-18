package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeopleEndPoint(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(people)
}
func GetPersonEndPoint(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	for _, element := range people {
		if element.ID == params["id"] {
			json.NewEncoder(response).Encode(element)
			return
		}
	}
	json.NewEncoder(response).Encode(&Person{})
}
func CreatePersonEndPoint(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	var persona Person
	_ = json.NewDecoder(request.Body).Decode(&persona)
	persona.ID = params["id"]
	people = append(people, persona)
	json.NewEncoder(response).Encode(people)
}
func DeletePersonEndPoint(response http.ResponseWriter, request *http.Request) {

}
func main() {
	router := mux.NewRouter()
	people = append(people,
		Person{
			ID: "1",
			FirstName: "Edwin",
			LastName:  "Baeza",
			Address:   &Address{City: "Merida", State: "Yucatan"}})

	people = append(people,
		Person{ID: "2",
			FirstName: "David",
			LastName:  "Rodriguez",
			Address:   &Address{City: "Merida", State: "Yucatan"}})
	//endpoints = Nuevas Rutas
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
