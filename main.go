package main

import (
	"fmt"

	"github.com/yoiner-castillo-globant/GBcamp/db"
	"github.com/yoiner-castillo-globant/GBcamp/mensaje"
)

//import "fmt"

func main() {

	mensaje.Mensaje(db.Create("1", "123"))
	mensaje.Mensaje(db.Create("3", "3423"))
	mensaje.Mensaje(db.Create("55", "125443"))

	db.PrintDatos()

	var x interface{} = db.Retrieve("55")
	str := fmt.Sprintf("%v", x)
	fmt.Println("Dato encontrado::> " + str+" con el indice[55]") // "[1 2 3]"

	db.PrintDatos()
	mensaje.Mensaje("Dato Modificado::> " + db.Update("55", "92"))
	db.PrintDatos()
	mensaje.Mensaje("Dato Eliminado::> " + db.Delete("55"))
	db.PrintDatos()

	//mensaje.Mensaje(bd.Create("1", 123))
	/*
		mensaje.Mensaje(bd.Create("2", 44))
	*/
}
