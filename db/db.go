package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"github.com/yoiner-castillo-globant/GBcamp/constants/constants"	
)

type IDB interface {
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
	defer md.SaveMapInFile()
	md.mtx.Lock()
	if md.data[key] == "" || md.data[key] == nil {
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
	defer md.SaveMapInFile()
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
	defer md.SaveMapInFile()
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
	err := ioutil.WriteFile(constants.FilePath, jsonString, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (md *MemoryDB) ReadMapFromFile() {
	dataLikeBytes, err := ioutil.ReadFile(constants.FilePath)
	if err == nil {
		err = json.Unmarshal(dataLikeBytes, &md.data)
		if err != nil {
			panic(err)
		}
	}
}
