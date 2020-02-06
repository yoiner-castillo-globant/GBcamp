package main

import (
    //"fmt"
    "net/http"
	"github.com/gorilla/mux"
	"github.com/yoiner-castillo-globant/GBcamp/restful/cart"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "post called"}`))
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
    w.Write([]byte(`{"message": "put called"}`))
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "delete called"}`))
}


func main() {
    r := mux.NewRouter()
	icart := cart.CreateCart()
    icart.AddItem("1",3)
    icart.AddItem("3",8)
    icart.AddItem("4",2)

	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
    r.HandleFunc("/articles", icart.GetItems).Methods("GET")
    
	r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")
	
    http.ListenAndServe(":9090", r)
}