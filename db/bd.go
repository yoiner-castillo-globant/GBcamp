package db

import "fmt"

//import "strconv"

var datos = make(map[string]interface{})

func Create(key string, data interface{}) string {
	datos[key] = data
	conv := fmt.Sprintf("%v", data)
	return ("Ingresado la información[" + conv + "] en la llave [" + key + "]")
}

//strconv.Itoa( convierto int a string
func Retrieve(key string) interface{} {
	//dato := datos[key]
	var x interface{} = datos[key]
	return x
}

func Update(key string, data interface{}) string {
	datos[key] = data
	upda := fmt.Sprintf("%v", data)
	return ("actualizado [" + upda + "] en la llave [" + key + "]")
}

func Delete(key string) string {
	delete(datos, key)
	return ("Eliminando.. indice[" + key + "]")
}

func PrintDatos() {
	fmt.Println(datos)
	//	str := fmt.Sprintf("%v", datos)
	//fmt.Println(str) // "[1 2 3]"
}
