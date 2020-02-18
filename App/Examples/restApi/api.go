package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
	"github.com/yoiner-castillo-globant/GBcamp/App/Structs"
)



var people []Structs.Person

// EndPoints
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  for _, item := range people {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Structs.Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request){
  json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  var person Structs.Person
  _ = json.NewDecoder(req.Body).Decode(&person)
  person.ID = params["id"]
  people = append(people, person)
  json.NewEncoder(w).Encode(people)

}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  for index, item := range people {
    if item.ID == params["id"] {
      people = append(people[:index], people[index + 1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(people)
}

func main() {
  router := mux.NewRouter()
  
  // adding example data
  people = append(people, Structs.Person{ID: "1", FirstName:"Ryan", LastName:"Ray", Address: &Structs.Address{City:"Dubling", State:"California"}})
  people = append(people, Structs.Person{ID: "2", FirstName:"Maria", LastName:"Ray"})

  // endpoints
  router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
  router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":3000", router))
}