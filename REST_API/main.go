package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Person struct {
	ID 			string		`json:"id, omitempty"`
	Firstname	string		`json:"firstname, omitempty"`
	Lastname	string		`json:"lastname, omitempty"`
	Address		*Address	`json:"address, omitempty"`
}

type Persons struct {
	Person []Person `json:"persons"`
}

type Address struct {
	City	string	`json:"city, omitempty"`
	State 	string	`json:"state, omitempty"`
}

var persons Persons

// our main function
func main()  {
	jsonFile, err := os.Open("data.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened data.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &persons)

	for i := 0; i < len(persons.Person); i++ {
		fmt.Println("Id: " + persons.Person[i].ID)
		fmt.Println("Firstname: " + persons.Person[i].Firstname)
	}

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePeople).Methods("POST")
	router.HandleFunc("/people/{id}", UpdatePeople).Methods("PUT")
	router.HandleFunc("/people/{id}", DeleteStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// Get an item
func GetPeople(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(persons)
}
// Get an item by id
func GetPerson(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	for _, item := range persons.Person{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

//Create a new item
func CreatePeople(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	persons.Person = append(persons.Person, person)
	json.NewEncoder(w).Encode(person)
}

func UpdatePeople(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	for index, item := range persons.Person{
		if item.ID == params["id"]{
			_ = json.NewDecoder(r.Body).Decode(&persons.Person[index])
			break
		}
	}
	json.NewEncoder(w).Encode(persons)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	for index, item := range persons.Person  {
		if item.ID == params["id"] {
			persons.Person = append(persons.Person[:index], persons.Person[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(persons)
}