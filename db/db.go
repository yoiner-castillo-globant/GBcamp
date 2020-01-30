package db

import (
	"errors"
	"fmt"
	"sync"
)

var DATA = make(map[string]interface{})
var mu sync.Mutex

func Create(key string, value interface{})  error {
	mu.Lock()
	if DATA[key] == nil {
		DATA[key] = value
	} else {
		mu.Unlock()
		return  errors.New("Error, cannot be created, the key already exists")
	}
	mu.Unlock()
	return  nil
}


func Retrieve(key string) (interface{}, error) {

	if DATA[key] == nil {
		return nil, errors.New("No information was found with the key received")
	}
	return DATA[key], nil
}

func Update(key string, value interface{})  error {
	mu.Lock()
	if DATA[key] != nil {
		DATA[key] = value
	} else {
		mu.Unlock()
		return errors.New("No information was found with the key received")
	}
	mu.Unlock()
	return nil
}

func Delete(key string)  error {
	mu.Lock()
	if DATA[key] != nil {
		delete(DATA, key)
	}else{
		mu.Unlock()
		return errors.New("No information was found with the key received")
	}
	mu.Unlock()
	return nil
}

func PrintDATA() {
	mu.Lock()
	if len(DATA) > 0 {
		fmt.Println(DATA)
	}
	mu.Unlock()
}
