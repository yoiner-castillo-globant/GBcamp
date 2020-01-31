package db

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sync"
	"encoding/json"
	"log"
)

type iDB interface {
	Create(string, interface{}) error
	Retrieve(string) (interface{}, error)
	Update(string, interface{}) error
	Delete(string) error
	PrintDATA()
	SaveMapInFile()
	ReadMapFromFile()
}

type MemoryDB struct {
	data map[string]interface{}
	mtx  sync.Mutex
}

func NewMemoryDB() *MemoryDB {
	totalMemory := make(map[string]interface{})
	return &MemoryDB{data: totalMemory}
}

func (md *MemoryDB) Len() int {
	return len(md.data)
}

func (md *MemoryDB) Create(key string, value interface{}) error {
	md.mtx.Lock()

	if md.data[key] == nil {
		md.data[key] = value
	} else {
		md.mtx.Unlock()
		return errors.New("Error, cannot be created, the key already exists")
	}
	md.mtx.Unlock()
	return nil
}

func (md *MemoryDB) Retrieve(key string) (interface{}, error) {

	if md.data[key] == nil {
		return nil, errors.New("No information was found with the key received")
	}
	return md.data[key], nil
}

func (md *MemoryDB) Update(key string, value interface{}) error {
	md.mtx.Lock()
	if md.data[key] != nil {
		md.data[key] = value
	} else {
		md.mtx.Unlock()
		return errors.New("No information was found with the key received")
	}
	md.mtx.Unlock()
	return nil
}

func (md *MemoryDB) Delete(key string) error {
	md.mtx.Lock()
	if md.data[key] != nil {
		delete(md.data, key)
	} else {
		md.mtx.Unlock()
		return errors.New("No information was found with the key received")
	}
	md.mtx.Unlock()
	return nil
}

func (md *MemoryDB) PrintDATA() {
	md.mtx.Lock()
	if len(md.data) > 0 {
		fmt.Println(md.data)
	}
	md.mtx.Unlock()
}

func (md *MemoryDB) SaveMapInFile() {
	jsonString, _ := json.Marshal(md.data)
	err := ioutil.WriteFile("./io/Info.txt", jsonString, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (md *MemoryDB) ReadMapFromFile() {
	datosComoBytes, err := ioutil.ReadFile("./io/Info.txt")
	if err == nil {
		err = json.Unmarshal(datosComoBytes, &md.data)
		if err != nil {
			panic(err)
		}
	}
}
