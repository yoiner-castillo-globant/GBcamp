package main

import (
	//"github.com/yoiner-castillo-globant/GBcamp/App/ApiRest"
	"fmt"

	"github.com/yoiner-castillo-globant/GBcamp/App/Mongo/Ado"
	//"github.com/yoiner-castillo-globant/GBcamp/App/ApiRest/Control"
	"github.com/yoiner-castillo-globant/GBcamp/App/ApiRest/Logic"
	"github.com/yoiner-castillo-globant/GBcamp/App/Request"
)

func main(){
//	ApiRest.TurnOn()

	//Insertar() 
	//Update() 

	//showCart("Cart1")
	//Update()

	//ShowElementsCart("Cart1")
	//  val, err:=Ado.GetCart("Cart1")
	//  fmt.Println("val:" + val.IdCart )
	//  fmt.Println("err:" + err )

	// dato:= Ado.Find("Cart1", "product3")
	// fmt.Println(dato.Title)
Ado.ChangeItemAmount("Cart1", "product3", 500)
}

func Update()  {
	//Elements := make(map[string]Request.ArticleCart)
	 article2 := Request.Article{ArticleId:"product3", Price: "3,4", Title: "TV-60" }
	 cartitem2 := Request.ArticleCart{Amount: 12, Element: article2}
	// Elements[cartitem2.Element.ArticleId] = cartitem2

	Ado.AddItemCart("Cart1",cartitem2)
}

func ShowElementsCart(key string)  {
	
	items, err := Ado.GetItemsCarts(key)
	if err!= nil{
		fmt.Println(err )
	  }
	  for index, cartarticles := range items {
		fmt.Println(cartarticles.Element)
		fmt.Println(cartarticles.Amount)
		fmt.Println(index)

	}
}

func showCart(key string)  {
	val, err:=Ado.GetCart(key)
	
	if err!= nil{
	  fmt.Println(err )
	}

	for index, articles := range val.Elements {
		fmt.Println(articles.Element)
		fmt.Println(articles.Amount)
		fmt.Println(index)

	}
	
	fmt.Println(val)
}
func Insertar()  {
	Elements := []*Request.ArticleCart{}
	article := Request.Article{ArticleId:"product1", Price: "3,4", Title: "TV" }
	article2 := Request.Article{ArticleId:"product2", Price: "3,4", Title: "TV-60" }

	cartitem1 := &Request.ArticleCart{Amount: 30, Element: article}
	cartitem2 := &Request.ArticleCart{Amount: 12, Element: article2}

	Elements = append(Elements,cartitem1)
	Elements = append(Elements,cartitem2)

	
	data :=Logic.Cart{IdCart: "Cart1", Elements: Elements}

	val, err :=  Ado.InsertCart(data)
	fmt.Println(val)
	fmt.Println(err)
}