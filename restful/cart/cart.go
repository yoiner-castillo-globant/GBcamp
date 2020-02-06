package cart

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yoiner-castillo-globant/GBcamp/constants"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)

type ICart interface {
	AddItem(int, int) error
	GetAllItems() map[Element]int
	ChangeItemAmount(int, int) error
	DeleteItem(int)
	DeleteAllItems()
	PrintCart()
	ReadBook(http.ResponseWriter, *http.Request)
}

type Cart struct {
	elements map[Element]int
}

//Element _
type Element struct {
	Id    string     
	Title string  
	Price float64 
}

type apiStruct struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Price string `json:"price"`
}

func CreateCart() *Cart {
	items := make(map[Element]int)
	return &Cart{elements: items}
}

func (ct *Cart) AddItem(_idProduct string, _amount int) error {
	item := getElementFromApi(_idProduct)
	ct.elements[item] = _amount
	return nil
}

func (ct *Cart) GetAllItems() map[Element]int {
	return ct.elements
}

func (ct *Cart) ChangeItemAmount(_idkey string, _amount int) error {
	changed := false

	item := ct.getElementFromMap(_idkey)
	ct.elements[item] = _amount
	if item.Id == "" {
		changed = false
	}

	if !changed {
		return errors.New("Error, ChangeItemAmount, IdKey not found")
	}

	return nil
}
func (ct *Cart) DeleteItem(_idkey string) {
	item := ct.getElementFromMap(_idkey)
	delete(ct.elements, item)
}
func (ct *Cart) DeleteAllItems() {
	ct.elements = nil
}
func (ct *Cart) PrintCart() {
	fmt.Println(ct)
}

func PrintCart(_data map[Element]int) {
	fmt.Println(_data)
}

func (ct *Cart) getElementFromMap(_idProduct string) Element {
	for key := range ct.elements {
		if key.Id == _idProduct {
			return key
		}
	}
	return Element{}
}

type responseStruct struct {
	Id string
	Title string
	Price float64
	Amount int
}
func encodeMap(_data map[Element]int) []responseStruct {
	items := []responseStruct{}
	for key, value := range _data {
	item := responseStruct{Amount: value, Id:key.Id, Title: key.Title, Price: key.Price}
	items = append(items, item )
	}
	return items
}

func (ct *Cart) GetItems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonString := encodeMap(ct.elements)
	json.NewEncoder(w).Encode(jsonString)
}


func getElementFromApi(_idProducto string) Element {
	var product apiStruct
	Url := constants.ApiUrlProducts + "/" + _idProducto
	var Client = &http.Client{Timeout: 10 * time.Second}
	resp, err := Client.Get(Url)

	if err != nil {
		panic(err.Error())
	}
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &product)

	PriceProduct, err := strconv.ParseFloat(product.Price, 32)

	return Element{Id: product.Id, Title: product.Title, Price: PriceProduct}
}

func removeIndex(s []Element, index int) []Element {
	return append(s[:index], s[index+1:]...)
}
