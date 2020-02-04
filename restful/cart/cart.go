package cart

import "errors"

import "fmt"

type iCart interface {
	AddItems([]Element) error
	AddItem(Element) error
	GetAllItems() []Element
	ChangeItemAmount(int, int) error
	DeleteItem(int)
	DeleteAllItems()
	PrintCart()
}

type Cart struct {
	elements []Element
}

//Element _
type Element struct {
	Id     int     `json:"idItem"`
	Title  string  `json:"titelItem"`
	Price  float32 `json:"priceItem"`
	Amount int     `json:"amountItem"`
}

func CreateCart() *Cart {
	items := []Element{}
	return &Cart{elements: items}
}

func (ct *Cart) AddItems(items []Element) error {
	for _, item := range items {
		ct.elements = append(ct.elements, item)
	}
	return nil
}

func (ct *Cart) AddItem(item Element) error {
	ct.elements = append(ct.elements, item)
	return nil
}

func (ct *Cart) GetAllItems() []Element {
	return ct.elements
}
func (ct *Cart) ChangeItemAmount(Idkey int, amount int) error {
	changed := false
	for i, item := range ct.elements {
		if item.Id == Idkey {
			element := item
			element.Amount = amount
			ct.elements[i] = element
			changed = true
		}
	}
	if !changed{
		return errors.New("Error, ChangeItemAmount, IdKey not found")
	}

	return nil
}
func (ct *Cart) DeleteItem(Idkey int) {
	var index int
	for i, item := range ct.elements {
		if item.Id == Idkey {
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
func PrintCart(items []Element) {
	fmt.Println(items)
}

func removeIndex(s []Element, index int) []Element {
	return append(s[:index], s[index+1:]...)
}


