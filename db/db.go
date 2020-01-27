package db

import (
	"fmt"
	"sync"
)

//import "strconv"

var Datos = make(map[string]interface{})
var mu sync.Mutex

func Create(key string, data interface{}) bool {
	creo := false
	mu.Lock()
	if Datos[key] == nil {
		Datos[key] = data
		creo = true
	}
	mu.Unlock()
	return creo
}

//strconv.Itoa( convierto int a string
func Retrieve(key string) interface{} {
	return Datos[key]
}

func Update(key string, data interface{}) bool {
	actualizo := false
	mu.Lock()
	if Datos[key] != nil {
		Datos[key] = data
		actualizo = true
	}
	mu.Unlock()
	return actualizo
}

func Delete(key string) bool {
	elimino := false
	mu.Lock()
	if Datos[key] != nil {
		delete(Datos, key)
		elimino = true
	}
	mu.Unlock()
	return elimino
}

func PrintDatos() {
	mu.Lock()
	if len(Datos) > 0 {
		fmt.Println(Datos)
	}
	mu.Unlock()
}
