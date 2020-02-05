package cart

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"github.com/yoiner-castillo-globant/GBcamp/constants"
)

type ICart interface {
	AddItems([]Element) error
	AddItem(Element) error
	GetAllItems() []Element
	ChangeItemAmount(int, int) error
	DeleteItem(int)
	DeleteAllItems()
	PrintCart()
}

type Cart struct {
	elements map[Element]int
}

//Element _
type Element struct {
	Id     int     `json:"idItem"`
	Title  string  `json:"titelItem"`
	Price  float64 `json:"priceItem"`
}

type apiStruct struct {
	Id    string	`json:"id"`
	Title string	`json:"title"`
	Price string	`json:"price"`
}

func CreateCart() *Cart {
	items := make(map[Element]int)
	return &Cart{elements: items}
}

func (ct *Cart) AddItems(_items []Element) error {
		ct.elements = append(ct.elements, _items...)
	return nil
}

func (ct *Cart) AddItem(_id int, _amount int) error {
	item := getItemFromApi(_id, _amount)
	ct.elements = append(ct.elements, item)
	return nil
}

func (ct *Cart) GetAllItems() []Element {
	return ct.elements
}
func (ct *Cart) ChangeItemAmount(_idkey int, _amount int) error {
	changed := false
	for i, item := range ct.elements {
		if item.Id == _idkey {
			element := item
			element.Amount = _amount
			ct.elements[i] = element
			changed = true
		}
	}
	if !changed {
		return errors.New("Error, ChangeItemAmount, IdKey not found")
	}

	return nil
}
func (ct *Cart) DeleteItem(_idkey int) {
	var index int
	for i, item := range ct.elements {
		if item.Id == _idkey {
			index = i
		}
	}
	ct.elements = removeIndex(ct.elements, index)
}
func (ct *Cart) DeleteAllItems() {
	ct.elements = nil
}
func (ct *Cart) PrintCart() {
	fmt.Println(ct)
}
func PrintCart(_items []Element) {
	fmt.Println(_items)
}

func getItemFromApi(_id int, _amount int) Element {
	var product apiStruct
	Url := (constants.ApiUrlProducts + "/" + strconv.Itoa(_id))
	var Client = &http.Client{Timeout: 10 * time.Second}
	resp, err := Client.Get(Url)

	if err != nil {
		panic(err.Error())
	}
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &product)

	IdProduct, err := strconv.Atoi(product.Id)
	PriceProduct, err := strconv.ParseFloat(product.Price, 32)	

	return Element{Id: IdProduct, Title: product.Title, Price:PriceProduct, Amount: _amount}
}

func removeIndex(s []Element, index int) []Element {
	return append(s[:index], s[index+1:]...)
}
