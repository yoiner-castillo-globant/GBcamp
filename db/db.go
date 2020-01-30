package db

import (
	"errors"
	"fmt"
	"sync"
)

//import "strconv"

var Datos = make(map[string]interface{})
var mu sync.Mutex

//func Create(key <-chan string, data <-chan interface{}) bool {

func Create(key string, data interface{}) (bool, error) {
	creo := false
	mu.Lock()
	//llave := <-key
	if Datos[key] == nil {
		Datos[key] = data
		creo = true
	} else {
		mu.Unlock()
		return false, errors.New("Error, No se puede crear, la clave ya existe")
	}
	mu.Unlock()
	return creo, nil
}

//strconv.Itoa( convierto int a string
func Retrieve(key string) (interface{}, error) {
	//func Retrieve(key <-chan string) interface{} {

	if Datos[key] == nil {
		return nil, errors.New("No se encontró información con la clave recibida")
	}
	return Datos[key], nil
}

func Update(key string, data interface{}) (bool, error) {
	//func Update(key <-chan string, data <-chan interface{}) bool {

	actualizo := false
	mu.Lock()
	//llave := <-key
	if Datos[key] != nil {
		Datos[key] = data
		actualizo = true
	} else {
		mu.Unlock()
		return false, errors.New("No se encontró información con la clave recibida")
	}
	mu.Unlock()
	return actualizo, nil
}

func Delete(key string) (bool, error) {
	//func Delete(key <-chan string) bool {

	elimino := false
	mu.Lock()
	//llave := <-key
	if Datos[key] != nil {
		delete(Datos, key)
		elimino = true
	}else{
		mu.Unlock()
		return false, errors.New("No se encontró información con la clave recibida")
	}
	mu.Unlock()
	return elimino, nil
}

func PrintDatos() {
	mu.Lock()
	if len(Datos) > 0 {
		fmt.Println(Datos)
	}
	mu.Unlock()
}
