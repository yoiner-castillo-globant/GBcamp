package db

import (
	"fmt"
	"sync"
)

//import "strconv"

var Datos = make(map[string]interface{})
var mu sync.Mutex

func Create(key <-chan string, data <-chan interface{}) bool {

	//func Create(key string, data interface{}) bool {
	creo := false
	mu.Lock()
	llave := <-key
	if Datos[llave] == nil {
		Datos[llave] = <-data
		creo = true
	}
	mu.Unlock()
	return creo
}

//strconv.Itoa( convierto int a string
//func Retrieve(key string) interface{} {
	func Retrieve(key <-chan string) interface{} {	
	return Datos[<-key]
}

//func Update(key string, data interface{}) bool {
func Update(key <-chan string, data <-chan interface{}) bool {

	actualizo := false
	mu.Lock()
	llave := <-key
	if Datos[llave] != nil {
		Datos[llave] = <-data
		actualizo = true
	}
	mu.Unlock()
	return actualizo
}

//func Delete(key string) bool {
func Delete(key <-chan string) bool {

	elimino := false
	mu.Lock()
	llave := <-key
	if Datos[llave] != nil {
		delete(Datos, llave)
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
