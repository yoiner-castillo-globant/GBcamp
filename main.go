package main

import (
	"github.com/yoiner-castillo-globant/GBcamp/db"
	//"github.com/yoiner-castillo-globant/GBcamp/restful/cart"
)



func main() {
/*	icart := cart.CreateCart()
	item1 := cart.Element{Id: 2, Title:"", Price:3.4, Amount:2}
	item2 := cart.Element{3,"daa", 5.2,3}

	items := []cart.Element{}
	items = append(items, item1)
	items = append(items, item2)
	item3 := cart.Element{Id: 4, Title:"33", Price:3.7, Amount:2}


	icart.AddItems(items)
	icart.PrintCart()

	icart.AddItem(item3)
	cart.PrintCart(icart.GetAllItems())

	icart.ChangeItemAmount(1,2)
	cart.PrintCart(icart.GetAllItems())

	icart.DeleteItem(1)
	cart.PrintCart(icart.GetAllItems())

	icart.DeleteAllItems()
	cart.PrintCart(icart.GetAllItems())*/

	
		data := db.NewMemoryDB()
		data.ReadMapFromFile()
		data.PrintDATA()
		data.Create("23", 23445)
		data.Create("532", "Polo")
		data.Create("Jhoi", "temporal")
		data.PrintDATA()
		data.Update("23", "una cadena")
		data.PrintDATA()
		data.Delete("532")
		data.PrintDATA()
		

}
