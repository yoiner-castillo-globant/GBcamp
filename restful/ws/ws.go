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
	cartId string
}

func NewCart(w http.ResponseWriter, r *http.Request) {
	newId := betterguid.New()
	IManager.AddCart(newId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"cartId": "` + newId + `"}`))
	// json.NewEncoder(w).Encode(`{"cartId": "`+ newId + `"}`)
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	CartId := r.FormValue("cartId")
	IdElement := r.FormValue("articleId")
	pathParams := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if CartId == "" {
		CartId = pathParams["cartId"]
	}
	if IdElement == "" {
		IdElement = pathParams["articleId"]
	}

	// fmt.Println("cartID::> ", CartId)
	//  fmt.Println("IdElement: ", IdElement)

	if CartId == "" {
		w.Write([]byte(`{"response": "you need to send an cartId"}`))
	} else if IdElement == "" {
		w.Write([]byte(`{"response": "you need to send an articleId"}`))
	} else {
		
		err := IManager.AddItem(CartId, IdElement);
		
		if  err != nil {
			
			w.Write([]byte(`{"response": "you need to send an articleId valid"}`))
		}
		
		w.Write([]byte(`{"response": "Added successfully"}`))
		//json.NewEncoder(w).Encode(`{"response": "Added successfully"}`)
	}
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	CartId := pathParams["cartId"]
	json.NewEncoder(w).Encode(IManager.GetItems(CartId))
}

func ChangeAmountItem(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	CartId := pathParams["cartId"]
	IdElement := pathParams["articleId"]
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
	if err = IManager.ChangeItemAmount(CartId, IdElement, Amount); err != nil {
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	} else {
		w.Write([]byte(`{"done": "Added successfully"}`))
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	CartId := pathParams["cartId"]
	ArticleId := pathParams["articleId"]
	IManager.DeleteItem(CartId, ArticleId)
	w.Write([]byte(`{"done": "The product was eliminated"}`))
}

func DeleteAllItems(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	CartId := pathParams["cartId"]
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
	r.HandleFunc("/NewCart", NewCart)
	r.HandleFunc("/cart/{cartId}/article/{articleId}", AddItem).Methods("POST")
	r.HandleFunc("/AddItem", AddItem).Methods("POST")
	r.HandleFunc("/cart/{cartId}/articles", GetItems).Methods("GET")
	r.HandleFunc("/cart/{cartId}/article/{articleId}/amount/{amount}", ChangeAmountItem).Methods("PUT")
	r.HandleFunc("/cart/{cartId}/article/{articleId}", DeleteItem).Methods("DELETE")
	r.HandleFunc("/cart/{cartId}", DeleteAllItems).Methods("DELETE")
	return r
}
