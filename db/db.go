package db

import "fmt"

//import "strconv"

var Datos = make(map[string]interface{})

func Create(key string, data interface{}) string {
	Datos[key] = data
	conv := fmt.Sprintf("%v", data)
	return ("Ingresado la información[" + conv + "] en la llave [" + key + "]")
}

//strconv.Itoa( convierto int a string
func Retrieve(key string) interface{} {
	//dato := datos[key]
	var x interface{} = Datos[key]
	return x
}

func Update(key string, data interface{}) string {
	Datos[key] = data
	upda := fmt.Sprintf("%v", data)
	return ("actualizado [" + upda + "] en la llave [" + key + "]")
}

func Delete(key string) string {
	delete(Datos, key)
	return ("Eliminando.. indice[" + key + "]")
}

func PrintDatos() {
	fmt.Println(Datos)
}
