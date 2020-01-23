package db

import "fmt"

//import "strconv"

var Datos = make(map[string]interface{})

func Create(key string, data interface{}) bool {
	if Datos[key] == nil {
		Datos[key] = data
		return true
	}
	return false
}

//strconv.Itoa( convierto int a string
func Retrieve(key string) interface{} {
	return Datos[key]
}

func Update(key string, data interface{}) bool {
	if Datos[key] != nil {
		Datos[key] = data
		return true
	}
	return false
}

func Delete(key string) bool{
	if Datos[key] != nil{
		delete(Datos, key)
		return true
	}
	return false
}

func PrintDatos() {
	if len(Datos) > 0 {
		fmt.Println(Datos)
	}
}
