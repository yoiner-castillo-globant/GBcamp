package manager

import (
	"github.com/yoiner-castillo-globant/GBcamp/restful/cart"
	"github.com/yoiner-castillo-globant/GBcamp/structs"
)

type IAdmin interface {
	DeleteItem(string, string)
	DeleteAll(string)
	AddCart(string)
	AddItemCart(string, string)
	ChangeItemAmount(string, string, int ) error
	GetItems(string) []structs.ResponseStruct
}

type Admin struct {
	Carts map[string]*cart.Cart
}

func New() *Admin {
	carts := make(map[string]*cart.Cart)
	return &Admin{Carts: carts}
}

func (a *Admin) AddCart(keyCart string) {	
	a.Carts[keyCart] = cart.CreateCart()
}
func (a *Admin) AddItem(keyCart string, keyProduct string) error{	
	icart := a.Carts[keyCart] 
	if err := icart.AddItem(keyProduct,1); err!= nil{
		return err
	}
	return nil
}
func (a *Admin) GetItems(keyCart string) []structs.ResponseStruct{	
	icart := a.Carts[keyCart] 
	return cart.EncodeMap(icart.Elements)
}
func (a *Admin) ChangeItemAmount(keyCart string, keyProduct string, amount int ) error{	
	icart := a.Carts[keyCart] 
	return icart.ChangeItemAmount(keyProduct,amount)
}
func (a *Admin) DeleteItem(keyCart string, idProduct string) {	
	icart := a.Carts[keyCart]
	icart.DeleteItem(idProduct)
}
func (a *Admin) DeleteAll(keyCart string) {	
	icart := a.Carts[keyCart]
	icart.DeleteAllItems()
}
