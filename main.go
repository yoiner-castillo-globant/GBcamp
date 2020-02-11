package main

import (
	//"github.com/yoiner-castillo-globant/GBcamp/db"
	//"github.com/yoiner-castillo-globant/GBcamp/constants"
	"github.com/yoiner-castillo-globant/GBcamp/restful/cart"
)



func main() {
	icart := cart.CreateCart()

	icart.AddItem("1", 3)
	icart.PrintCart()
	icart.ChangeItemAmount("1",5)
	icart.PrintCart()

	cart.PrintCart(icart.GetAllItems())
	icart.AddItem("2", 7)
	icart.DeleteItem("1")
	icart.PrintCart()
	icart.DeleteAllItems()
	icart.PrintCart()


}
