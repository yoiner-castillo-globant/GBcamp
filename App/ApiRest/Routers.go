package ApiRest

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kjk/betterguid"
	"github.com/yoiner-castillo-globant/GBcamp/App/Request"
	"github.com/yoiner-castillo-globant/GBcamp/App/ApiRest/Control"
)

var IControl = Control.New()

func NewCartEP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	newId := betterguid.New()
	IControl.AddCart(newId)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"CartId": "` + newId + `"}`))
}

func getItemsCartEP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pathParams := mux.Vars(req)
	cartId := pathParams["CartId"]

	items, err := IControl.GetItems(cartId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"Response": "you need to send a valid idCart"}`))
		log.Println("error: invalid idCart") 
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func validateAddItemRequest(w http.ResponseWriter, cartId, articleId string) bool {
	if cartId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"Response": "you need to send an CartId"}`))
		log.Println("error: CartId not sent") 
		return false
	}

	if articleId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"Response": "you need to send a product id"}`))
		log.Println("error: product id not sent") 
		return false
	}

	return true
}

func addItemCartEP(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	article := Request.PostArticle{}
	cartId := params["CartId"]

	_ = json.NewDecoder(req.Body).Decode(&article)

	if !validateAddItemRequest(w, cartId, article.ArticleId) {
		return
	}

	if err := IControl.AddItem(cartId, article.ArticleId); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"Response": `+err.Error()+` }`))
		log.Println("error: product id invalid") 
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"Response": "Added successfully"}`))

}

func changeAmountItemEP(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	article := Request.PostArticle{}
	cartId := params["CartId"]

	_ = json.NewDecoder(req.Body).Decode(&article)

	if !validateAddItemRequest(w, cartId, article.ArticleId) {
		return
	}

	if article.Amount <= 0{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"Response": "you need to send a quantity greater than zero"}`))
		log.Println("error: quantity <= 0") 
		return
	}

	if err := IControl.ChangeItemAmount(cartId, article.ArticleId, article.Amount); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"Response": "you need to send an existing id in the cart"}`))
		log.Println("error: product id doesn´t exist in the cart") 
		return
	} 

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"Response": "Changed successfully"}`))
		
	
}

func deleteItemCartEP(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")
	cartId := params["CartId"]
	articleId := params["ArticleId"]

	if !validateAddItemRequest(w, cartId, articleId) {
		return
	}

	if err:= IControl.DeleteItem(cartId, articleId); err != nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"Response": "you need to send an existing id in the cart"}`))
		log.Println("error: product id doesn´t exist in the cart") 
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"done": "The product was eliminated"}`))

}

func deleteItemsCartEP(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")
	cartId := params["CartId"]

	if !validateAddItemRequest(w, cartId, "articleId") {
		return
	}

	IControl.DeleteAll(cartId)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"done": "The products were eliminated"}`))

}

func LoadServer() error {

	err := http.ListenAndServe(":3000", LoadEndPoints())
	if err != nil {
		log.Println(err) 
		return err
	}
	return nil
}

func LoadEndPoints() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/CreateCart", NewCartEP).Methods(http.MethodGet)
	router.HandleFunc("/AddItem/{CartId}", addItemCartEP).Methods(http.MethodPost)
	router.HandleFunc("/GetItems/{CartId}", getItemsCartEP).Methods(http.MethodGet)
	router.HandleFunc("/ChangeQuantity/{CartId}", changeAmountItemEP).Methods(http.MethodPut)
	router.HandleFunc("/Delete/{CartId}/article/{ArticleId}", deleteItemCartEP).Methods(http.MethodDelete)
	router.HandleFunc("/Delete/{CartId}", deleteItemsCartEP).Methods(http.MethodDelete)
	return router
}
