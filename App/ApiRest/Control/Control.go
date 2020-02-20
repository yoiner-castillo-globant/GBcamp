package Control

import(
	"errors"
	"github.com/yoiner-castillo-globant/GBcamp/App/Request"
	"github.com/yoiner-castillo-globant/GBcamp/App/ApiRest/Logic"
)

type IControl interface{
	AddCart(string)
	AddItem(string, string) error
	GetItems(string) (map[string]Request.ArticleCart, error)
	ChangeItemAmount(string, string, int ) error
	DeleteItem(string, string) error
	DeleteAll(string) error
}

type Admin struct {
	Carts map[string]*Logic.Cart
}

func New() *Admin {
	carts := make(map[string]*Logic.Cart)
	return &Admin{Carts: carts}
}

func (a *Admin) AddCart(keyCart string) {	
	a.Carts[keyCart] = Logic.NewCart(keyCart)
}

func (a *Admin) AddItem(keyCart string, keyProduct string) error{	
	icart := a.Carts[keyCart] 
	if err:= validateEmptyCart(icart.Id); err!= nil{
		return err
	}
	if err := icart.AddItem(keyProduct); err!= nil{
		return err
	}
	return nil
}

func (a *Admin) GetItems(keyCart string) (map[string]Request.ArticleCart, error){	
	icart := a.Carts[keyCart] 
	if err:= validateEmptyCart(icart.Id); err!= nil{
		return nil, err
	}
	return icart.Elements, nil
}
func (a *Admin) ChangeItemAmount(keyCart string, keyProduct string, amount int ) error{	
	icart := a.Carts[keyCart] 
	if err:= validateEmptyCart(icart.Id); err!= nil{
		return err
	}
	return icart.ChangeItemAmount(keyProduct,amount)
}
func (a *Admin) DeleteItem(keyCart string, keyProduct string) error {	
	icart := a.Carts[keyCart]
	if err:= validateEmptyCart(icart.Id); err!= nil{
		return err
	}
	return icart.DeleteItem(keyProduct)
}
func (a *Admin) DeleteAll(keyCart string) error {	
	icart := a.Carts[keyCart]

	if err:= validateEmptyCart(icart.Id); err!= nil{
		return err
	}
	icart.DeleteAllItems()
	return nil
}

func validateEmptyCart(idCart string) error  {
	if idCart == ""{
		return 	errors.New("Error, keyCart not found")
	}
	return nil
}