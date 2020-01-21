package main


//import "github.com/yoiner-castillo-globant/GBcamp/mensaje"
import "github.com/yoiner-castillo-globant/GBcamp/bd"

import "fmt"


func main(){
	//mensaje.Mensaje("vamos")
bd.Init()
	fmt.Println(bd.Create("1", 123))
	//mensaje.Mensaje(bd.Create("1", 123))
	/*
	mensaje.Mensaje(bd.Create("2", 44))
*/
}