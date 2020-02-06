package main

import (
    //"fmt"
    "net/http"
	"github.com/gorilla/mux"
	"github.com/yoiner-castillo-globant/GBcamp/restful/cart"
)


func main() {
    r := mux.NewRouter()

	icart := cart.CreateCart()
    icart.AddItem("1",3)
    icart.AddItem("3",8)
    icart.AddItem("4",2)

    r.HandleFunc("/article/{articleId}", icart.WSAddItem).Methods("POST")
    r.HandleFunc("/articles", icart.WSGetItems).Methods("GET")
    r.HandleFunc("/article/{articleId}/amount/{amount}", icart.WSChangeItemAmount).Methods("PUT")
	r.HandleFunc("/article/{articleId}",icart.WSDeleteItem).Methods("DELETE")
	r.HandleFunc("/articles",icart.WSDeleteAllItems).Methods("DELETE")
	
    http.ListenAndServe(":9090", r)
}