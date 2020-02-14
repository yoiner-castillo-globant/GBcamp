package cart

import (
	"encoding/json"
	"bytes"
	"errors"
	"fmt"
	"github.com/yoiner-castillo-globant/GBcamp/constants"
	"github.com/yoiner-castillo-globant/GBcamp/structs"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type ICart interface {
	AddItem(int, int) error
	GetAllItems() map[structs.Element]int
	ChangeItemAmount(int, int) error
	DeleteItem(int)
	DeleteAllItems()
	PrintCart()
}
type Cart struct {
	Elements map[structs.Element]int
}

func CreateCart() *Cart {
	items := make(map[structs.Element]int)
	return &Cart{Elements: items}
}


func (ct *Cart) AddItem(idProduct string, amount int) error {

	item, err := getElementFromApi(idProduct)
	
	if err!= nil{
		return err
	}
	ct.Elements[item] = amount
	return nil
}

func (ct *Cart) GetAllItems() map[structs.Element]int {
	return ct.Elements
}

func (ct *Cart) ChangeItemAmount(idkey string, amount int) error {
	changed := false

	item := ct.getElementFromMap(idkey)
	if ct.Elements[item] == 0 {
		changed = false
	}else{
		ct.Elements[item] = amount
	}

	if ct.Elements[item] == amount{
		changed = true
	}
	if !changed {
		return errors.New("Error, ChangeItemAmount, IdKey not found")
	}

	return nil
}

func (ct *Cart) DeleteItem(idkey string) {
	item := ct.getElementFromMap(idkey)
	delete(ct.Elements, item)
}
func (ct *Cart) DeleteAllItems() {
	ct.Elements = CreateCart().Elements
}
func (ct *Cart) PrintCart() {
	fmt.Println(ct)
}

func PrintCart(data map[structs.Element]int) {
	fmt.Println(data)
}

func (ct *Cart) getElementFromMap(idProduct string) structs.Element {
	for key := range ct.Elements {
		if key.Id == idProduct {
			return key
		}
	}
	return structs.Element{}
}

func EncodeMap(data map[structs.Element]int) []structs.ResponseStruct {
	items := []structs.ResponseStruct{}
	for key, value := range data {
	item := structs.ResponseStruct{Amount: value, Id:key.Id, Title: key.Title, Price: key.Price}
	items = append(items, item )
	}
	return items
}


func getElementFromApi(_idProducto string) (structs.Element, error) {
	var product structs.ApiStruct
	Url := constants.ApiUrlProducts + "/" + _idProducto
	var Client = &http.Client{Timeout: 10 * time.Second}
	resp, err := Client.Get(Url)

	if err != nil{
		return structs.Element{}, err
	}else{
		data, _ := ioutil.ReadAll(resp.Body)
		if string(bytes.TrimSpace(data)) == "undefined"{
			return structs.Element{},  errors.New("undefined") 
		}else{
			json.Unmarshal(data, &product)
		}
		
	}
	PriceProduct, _ := strconv.ParseFloat(product.Price, 32)
	return structs.Element{Id: product.Id, Title: product.Title, Price: PriceProduct}, nil
}

func removeIndex(s []structs.Element, index int) []structs.Element {
	return append(s[:index], s[index+1:]...)
}
