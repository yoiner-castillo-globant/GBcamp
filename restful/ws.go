package main

import (
    "strconv"
    "encoding/json"
    "net/http"
	"github.com/gorilla/mux"
    "github.com/yoiner-castillo-globant/GBcamp/restful/manager"
    "github.com/kjk/betterguid"
)


func main() {
    r := mux.NewRouter()
    IManager := manager.New()

     r.HandleFunc("/NewCart", func(w http.ResponseWriter, r *http.Request) {
         newId := betterguid.New()
         IManager.AddCart(newId)
         json.NewEncoder(w).Encode("{cartId:"+newId+"}")
     })

    r.HandleFunc("/cart/{cartid}/article/{articleId}", func(w http.ResponseWriter, r *http.Request) {
        pathParams := mux.Vars(r)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK) 
        CartId :=   pathParams["cartid"] 
        IdElement :=    pathParams["articleId"] 
        
        IManager.AddItem(CartId, IdElement)

        json.NewEncoder(w).Encode("Added successfully")
    }).Methods("POST")
     r.HandleFunc("/cart/{cartid}/articles", func(w http.ResponseWriter, r *http.Request) {
        pathParams := mux.Vars(r)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK) 
        CartId :=   pathParams["cartid"] 
        json.NewEncoder(w).Encode(IManager.GetItems(CartId))
    }).Methods("GET")
    
    r.HandleFunc("/cart/{cartid}/article/{articleId}/amount/{amount}", func(w http.ResponseWriter, r *http.Request) {
        pathParams := mux.Vars(r)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK) 
        CartId :=   pathParams["cartid"] 
        IdElement :=    pathParams["articleId"] 
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
        if err =  IManager.ChangeItemAmount(CartId, IdElement, Amount); err!= nil{
            json.NewEncoder(w).Encode(`{"error": `+err.Error()+`}`)
        }else{
            json.NewEncoder(w).Encode(`{"done": "Added successfully"}`)
        }
    }    ).Methods("PUT")
	 r.HandleFunc("/cart/{cartid}/article/{articleId}", func(w http.ResponseWriter, r *http.Request) {
        pathParams := mux.Vars(r)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK) 
        CartId     :=  pathParams["cartid"] 
        ArticleId  :=  pathParams["articleId"] 
        IManager.DeleteItem(CartId, ArticleId)
        json.NewEncoder(w).Encode(`{"done": "the product was eliminated"}`)
    }).Methods("DELETE")
	 r.HandleFunc("/cart/{cartid}",func(w http.ResponseWriter, r *http.Request) {
        pathParams := mux.Vars(r)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK) 
        CartId     :=  pathParams["cartid"] 
        IManager.DeleteAll(CartId)
        json.NewEncoder(w).Encode(`{"done": "the products were eliminated"}`)
    }).Methods("DELETE")
	
    http.ListenAndServe(":9090", r)
}

