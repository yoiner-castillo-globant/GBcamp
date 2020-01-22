package main

import (
	"github.com/yoiner-castillo-globant/GBcamp/io"

	
)

/*
	"fmt"
	"github.com/yoiner-castillo-globant/GBcamp/testing"
	"github.com/yoiner-castillo-globant/GBcamp/db"
	"github.com/yoiner-castillo-globant/GBcamp/mensaje"
	"github.com/yoiner-castillo-globant/GBcamp/io"
*/

//import "fmt"

func main() {
//io.EscribirArchivo("Prueba3\nPrueba2\ntesting\ntesting2", "EscribirEnArchivo")
io.LeerArchivo("EscribirEnArchivo")
	//testing.TestCreate()
	/*
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
	*/

	//mensaje.Mensaje(bd.Create("1", 123))
	/*
		mensaje.Mensaje(bd.Create("2", 44))
	*/
}
