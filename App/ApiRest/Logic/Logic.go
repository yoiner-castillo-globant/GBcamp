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
	GetAllItems() []*Request.ArticleCart
	ChangeItemAmount(string, int) error
	DeleteItem(string) error
	DeleteAllItems()
	
}

type Cart struct {
	IdCart   string `json:"idcart,omitempty" bson:"idcart"`
	Elements []*Request.ArticleCart `json:"elements,omitempty" bson:"elements"`
}

//Creates a new cart
func NewCart(key string) *Cart {
	items := []*Request.ArticleCart{}
	return &Cart{IdCart: key, Elements: items}
}

//Adds products to a Cart
func (ct *Cart) AddItem(key string) error {
	article, _ := findItem(key, ct.Elements)
	if article.Element.ArticleId != "" {
		return errors.New("The article already exist in the cart")
	} 
	item, err := getElementFromApi(key)
	if err != nil {
		return err
	}
	if item.Element.ArticleId == ""{
		return errors.New("you need to send a valid id or")
	}

	ct.Elements = append(ct.Elements,  item)
	return nil
}

//Gets all Cart´s Items
func (ct *Cart) GetAllItems() []*Request.ArticleCart {
	return ct.Elements
}

//Changes the product´s amount
func (ct *Cart) ChangeItemAmount(key string, amount int) error {
	item, index := findItem(key, ct.Elements)
	if item.Element.ArticleId == "" {
		return errors.New("Error, ChangeItemAmount, Key not found")
	} 
	

	item.Amount = amount
	
	ct.Elements[index] = item

	return nil
}

//Deletes a product from a Cart
func (ct *Cart) DeleteItem(key string) error {
	item, index := findItem(key, ct.Elements)
	if item.Element.ArticleId == "" {
		return errors.New("Error, DeleteItem, Key not found")
	} 
	ct.Elements = remove(ct.Elements, index)
	
	return nil
}

//Deletes All cart´s products
func (ct *Cart) DeleteAllItems() {
	ct.Elements = NewCart(ct.IdCart).Elements
}

//Gets Information about of a product from a API
func getElementFromApi(key string) (*Request.ArticleCart, error) {
	var product Request.Article
	Url := Constants.ApiUrlProducts + "/" + key
	var Client = &http.Client{Timeout: 10 * time.Second}
	resp, err := Client.Get(Url)

	if err != nil{
		return &Request.ArticleCart{}, err
	}
	data, _ := ioutil.ReadAll(resp.Body)
	if string(bytes.TrimSpace(data)) == "undefined"{
		return &Request.ArticleCart{},  errors.New("you need to send a valid id or") 
	}
		json.Unmarshal(data, &product)
	
	return  &Request.ArticleCart{Element: product, Amount: 1}, nil
}

func remove(slice []*Request.ArticleCart, index int) []*Request.ArticleCart {
    return append(slice[:index], slice[index+1:]...)
}

func findItem(key string, data []*Request.ArticleCart) (*Request.ArticleCart, int) {
	for index, cart := range data {
		if cart.Element.ArticleId == key {
			// Found!
			return cart, index
		}
	}
	return &Request.ArticleCart{}, -1
}