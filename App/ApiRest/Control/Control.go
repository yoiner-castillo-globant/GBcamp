package Control

import(
	"github.com/yoiner-castillo-globant/GBcamp/App/Request"
	"github.com/yoiner-castillo-globant/GBcamp/App/ApiRest/Logic"

)

type IControl interface{
	AddCart(string)
	AddItem(string, string) error
	GetItems(string) map[string]Request.ArticleCart
	ChangeItemAmount(string, string, int ) error
	DeleteItem(string, string)	
	DeleteAll(string)
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
	if err := icart.AddItem(keyProduct); err!= nil{
		return err
	}
	return nil
}

func (a *Admin) GetItems(keyCart string) map[string]Request.ArticleCart{	
	icart := a.Carts[keyCart] 
	return icart.Elements
}
func (a *Admin) ChangeItemAmount(keyCart string, keyProduct string, amount int ) error{	
	icart := a.Carts[keyCart] 
	return icart.ChangeItemAmount(keyProduct,amount)
}
func (a *Admin) DeleteItem(keyCart string, keyProduct string) {	
	icart := a.Carts[keyCart]
	icart.DeleteItem(keyProduct)
}
func (a *Admin) DeleteAll(keyCart string) {	
	icart := a.Carts[keyCart]
	icart.DeleteAllItems()
}