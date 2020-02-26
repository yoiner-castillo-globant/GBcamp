package Control

import (
	"errors"
	"github.com/yoiner-castillo-globant/GBcamp/App/ApiRest/Logic"
	"github.com/yoiner-castillo-globant/GBcamp/App/Request"
)

type IControl interface{
	AddCart(string)
	AddItem(string, string) error
	GetItems(string) ([]*Request.ArticleCart, error)
	ChangeItemAmount(string, string, int ) error
	DeleteItem(string, string) error
	DeleteAll(string) error
}

type Admin struct {
	Carts []*Logic.Cart
}

func New() *Admin {
	var carts = []*Logic.Cart{}
	return &Admin{Carts: carts}
}

func (a *Admin) AddCart(keyCart string) {	
	a.Carts = append(a.Carts,  Logic.NewCart(keyCart))
}

func (a *Admin) AddItem(keyCart string, keyProduct string) error{	
	icart :=  findItem(keyCart, a.Carts) 
	if err:= validateEmptyCart(icart.IdCart); err!= nil{
		return err
	}
	if err := icart.AddItem(keyProduct); err!= nil{
		return err
	}
	return nil
}

func (a *Admin) GetItems(keyCart string) ([]*Request.ArticleCart, error){	
	icart :=  findItem(keyCart, a.Carts) 
	if err:= validateEmptyCart(icart.IdCart); err!= nil{
		return nil, err
	}
	return icart.Elements, nil
}
func (a *Admin) ChangeItemAmount(keyCart string, keyProduct string, amount int ) error{	
	icart :=  findItem(keyCart, a.Carts) 
	if err:= validateEmptyCart(icart.IdCart); err!= nil{
		return err
	}
	return icart.ChangeItemAmount(keyProduct,amount)
}
func (a *Admin) DeleteItem(keyCart string, keyProduct string) error {	
	icart :=  findItem(keyCart, a.Carts) 
	if err:= validateEmptyCart(icart.IdCart); err!= nil{
		return err
	}
	return icart.DeleteItem(keyProduct)
}
func (a *Admin) DeleteAll(keyCart string) error {	
	icart :=  findItem(keyCart, a.Carts) 

	if err:= validateEmptyCart(icart.IdCart); err!= nil{
		return err
	}
	icart.DeleteAllItems()
	return nil
}

func validateEmptyCart(idCart string) error  {
	if idCart == ""{
		return 	errors.New("you need to send a valid id or")
	}
	return nil
}

func findItem(key string, data []*Logic.Cart) *Logic.Cart {
	for _, cart := range data {
		if cart.IdCart == key {
			// Found!
			return cart
		}
	}
	return &Logic.Cart{}
}