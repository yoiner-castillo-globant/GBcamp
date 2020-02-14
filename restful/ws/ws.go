package ws

import (
	"encoding/json"
//	"fmt"
	"github.com/gorilla/mux"
	"github.com/kjk/betterguid"
	"github.com/yoiner-castillo-globant/GBcamp/restful/manager"
	"log"
	"net/http"
	"strconv"
)

var IManager = manager.New()

type response struct {
	CartId string
}

func NewCart(w http.ResponseWriter, r *http.Request) {
	newId := betterguid.New()
	IManager.AddCart(newId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"CartId": "` + newId + `"}`))
	// json.NewEncoder(w).Encode(`{"CartId": "`+ newId + `"}`)
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	CartId := r.FormValue("CartId")
	ArticleId := r.FormValue("ArticleId")
	pathParams := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	if CartId == "" {
		CartId = pathParams["CartId"]
	}
	if ArticleId == "" {
		ArticleId = pathParams["ArticleId"]
	}

	if CartId == "" {
		w.Write([]byte(`{"Response": "you need to send an CartId"}`))
	} else if ArticleId == "" {
		w.Write([]byte(`{"Response": "you need to send an ArticleId"}`))
	} else {
		
		err := IManager.AddItem(CartId, ArticleId);
		
		if  err != nil {
			
			w.Write([]byte(`{"Response": "you need to send an ArticleId valid"}`))
		}
		
		w.Write([]byte(`{"Response": "Added successfully"}`))
		//json.NewEncoder(w).Encode(`{"response": "Added successfully"}`)
	}
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	CartId := pathParams["CartId"]
	json.NewEncoder(w).Encode(IManager.GetItems(CartId))
}

func ChangeAmountItem(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	CartId := pathParams["CartId"]
	ArticleId := pathParams["ArticleId"]
	Amount := -1
	var err error
	if val, ok := pathParams["amount"]; ok {
		Amount, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"amount": "amount need a number"}`))
			return
		}
	}
	if err = IManager.ChangeItemAmount(CartId, ArticleId, Amount); err != nil {
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	} else {
		w.Write([]byte(`{"done": "Added successfully"}`))
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	CartId := pathParams["CartId"]
	ArticleId := pathParams["ArticleId"]
	IManager.DeleteItem(CartId, ArticleId)
	w.Write([]byte(`{"done": "The product was eliminated"}`))
}

func DeleteAllItems(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	CartId := pathParams["CartId"]
	IManager.DeleteAll(CartId)
	w.Write([]byte(`{"done": "The products were eliminated"}`))
}
func LoadServer() {

	err := http.ListenAndServe(":9090", LoadRouter())
	if err != nil {
		log.Fatal(err)
	}
}

func LoadRouter() *mux.Router {
	r := mux.NewRouter()
	r.Headers("Content-Type", "application/json",
		  "X-Requested-With", "XMLHttpRequest")
		  
	r.HandleFunc("/NewCart", NewCart).Methods(http.MethodGet)
	r.HandleFunc("/cart/{CartId}/article/{ArticleId}", AddItem).Methods(http.MethodPost)
	r.HandleFunc("/AddItem", AddItem).Methods(http.MethodPost)
	r.HandleFunc("/cart/{CartId}/articles", GetItems).Methods(http.MethodGet)
	r.HandleFunc("/cart/{CartId}/article/{ArticleId}/amount/{amount}", ChangeAmountItem).Methods(http.MethodPut)
	r.HandleFunc("/cart/{CartId}/article/{ArticleId}", DeleteItem).Methods(http.MethodDelete)
	r.HandleFunc("/cart/{CartId}", DeleteAllItems).Methods(http.MethodDelete)
	return r
}
