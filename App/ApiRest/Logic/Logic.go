package Logic

import (
	"time"
	"bytes"
	"errors"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/yoiner-castillo-globant/GBcamp/App/Request"
	"github.com/yoiner-castillo-globant/GBcamp/App/Constants"
)

type ICart interface {
	AddItem(string) error
	GetAllItems() map[string]Request.ArticleCart
	ChangeItemAmount(string, int) error
	DeleteItem(string) error
	DeleteAllItems() 
}


type Cart struct {
	Id			string
	Elements	map[string]Request.ArticleCart
}

//Creates a new cart
func NewCart(key string) *Cart {
	items := make(map[string]Request.ArticleCart)
	return &Cart{ Id: key, Elements: items}
}


//Adds products to a Cart
func (ct *Cart) AddItem(key string) error{
	item, err := getElementFromApi(key)
	if err!= nil{
		return err
	}
	ct.Elements[key] = item
	return nil
}
//Gets all Cart´s Items
func (ct *Cart) GetAllItems() map[string]Request.ArticleCart{
	return ct.Elements
}

//Changes the product´s amount
func (ct *Cart) ChangeItemAmount(key string,amount int) error{
	if ct.Elements[key].Element.ArticleId == ""{
		return	errors.New("Error, ChangeItemAmount, Key not found")
	}else{
		item := ct.Elements[key]
		item.Amount = amount
		ct.Elements[key] = item
	}
	return nil
}

//Deletes a product from a Cart
func (ct *Cart) DeleteItem( key string) error{
	if ct.Elements[key].Element.ArticleId == ""{
		return	errors.New("Error, DeleteItem, Key not found")
	}else{
		delete(ct.Elements, key)
	}
	return nil
}
//Deletes All cart´s products
func (ct *Cart) DeleteAllItems() {
	ct.Elements = NewCart(ct.Id).Elements
}


//Gets Information about of a product from a API
func getElementFromApi(key string) (Request.ArticleCart, error) {
	var product Request.Article
	Url := Constants.ApiUrlProducts + "/" + key
	var Client = &http.Client{Timeout: 10 * time.Second}
	resp, err := Client.Get(Url)

	if err != nil{
		return Request.ArticleCart{}, err
	}else{
		data, _ := ioutil.ReadAll(resp.Body)
		if string(bytes.TrimSpace(data)) == "undefined"{
			return Request.ArticleCart{},  errors.New("undefined") 
		}else{
			json.Unmarshal(data, &product)
		}
		
	}
	return  Request.ArticleCart{Element: &product, Amount: 1}, nil
}