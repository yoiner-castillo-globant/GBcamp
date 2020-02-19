package ApiRest

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kjk/betterguid"
	"github.com/yoiner-castillo-globant/GBcamp/App/ApiRest/Control"
	"github.com/yoiner-castillo-globant/GBcamp/App/Request"
)

var IControl = Control.New()

func NewCartEP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	newId := betterguid.New()
	IControl.AddCart(newId)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"CartId": "` + newId + `"}`))
}

func GetItemsCartEP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pathParams := mux.Vars(req)
	CartId := pathParams["CartId"]

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(IControl.GetItems(CartId))
}

func validateAddItemRequest(w http.ResponseWriter, cartId, articleId string) bool {
	if cartId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"Response": "you need to send an CartId"}`))
		return false
	}

	if articleId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"Response": "you need to send an id"}`))
		return false
	}

	return true
}
func AddItemCartEP(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var article Request.PostArticle
	cartId := params["CartId"]

	_ = json.NewDecoder(req.Body).Decode(&article)

	if !validateAddItemRequest(w, cartId, article.ArticleId) {
		return
	}

	if err := IControl.AddItem(cartId, article.ArticleId); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"Response": "you need to send a valid id"}`))
	
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"Response": "Added successfully"}`))

}

func ChangeAmountItemEP(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var article Request.PostArticle
	_ = json.NewDecoder(req.Body).Decode(&article)
	CartId := params["CartId"]

	if CartId == "" {
		w.Write([]byte(`{"Response": "you need to send an CartId"}`))
	} else if article.ArticleId == "" {
		w.Write([]byte(`{"Response": "you need to send an existing id in the cart"}`))
	} else if article.Amount == 0 {
		w.Write([]byte(`{"Response": "you need to send the product´s quantity like quantity. quantity != 0 "}`))
	} else {

		if err := IControl.ChangeItemAmount(CartId, article.ArticleId, article.Amount); err != nil {
			w.Write([]byte(`{"Response": "you need to send an id valid"}`))
		} else {
			w.Write([]byte(`{"Response": "Changed successfully"}`))
		}
	}
}

func DeleteItemCartEP(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")
	CartId := params["CartId"]
	ArticleId := params["ArticleId"]
	IControl.DeleteItem(CartId, ArticleId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"done": "The product was eliminated"}`))

}

func DeleteItemsCartEP(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")
	CartId := params["CartId"]
	IControl.DeleteAll(CartId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"done": "The products were eliminated"}`))

}

func LoadServer() error {

	err := http.ListenAndServe(":3000", LoadEndPoints())
	if err != nil {
		return err
	}
	return nil
}

func LoadEndPoints() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/CreateCart", NewCartEP).Methods(http.MethodGet)
	router.HandleFunc("/AddItem/{CartId}", AddItemCartEP).Methods(http.MethodPost)
	router.HandleFunc("/GetItems/{CartId}", GetItemsCartEP).Methods(http.MethodGet)
	router.HandleFunc("/ChangeQuantity/{CartId}", ChangeAmountItemEP).Methods(http.MethodPut)
	router.HandleFunc("/Delete/{CartId}/article/{ArticleId}", DeleteItemCartEP).Methods(http.MethodDelete)
	router.HandleFunc("/Delete/{CartId}", DeleteItemsCartEP).Methods(http.MethodDelete)
	return router
}
