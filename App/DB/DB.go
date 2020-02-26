package DB

import (
	"log"
	"sync"
	"errors"
	"io/ioutil"
	"encoding/json"
	"github.com/yoiner-castillo-globant/GBcamp/App/Constants"
)

type IDB interface {
	Create(string, interface{}) error
	Retrieve(string) (interface{}, error)
	Update(string, interface{}) error
	Delete(string) error
	
}

type MemoryDB struct {
	data map[string]interface{}
	mtx  sync.Mutex
}

func NewMemoryDB() *MemoryDB {
	totalMemory := make(map[string]interface{})
	return &MemoryDB{data: totalMemory}
}

//Get the dimensions of an object
func (md *MemoryDB) Len() int {
	return len(md.data)
}

// Create a new Element, if it doesn't exist
func (md *MemoryDB) Create(key string, value interface{}) error {
	md.mtx.Lock()
	if md.data[key] == "" || md.data[key] == nil {
		md.data[key] = value
	} else {
		md.mtx.Unlock()
		return errors.New("Error, cannot be created, the key already exists")
	}
	md.mtx.Unlock()
	return md.SaveMapInFile()
}

// Retrieve an Element if it exist
func (md *MemoryDB) Retrieve(key string) (interface{}, error) {
	if md.data[key] == nil {
		return nil, errors.New("No information was found with the key received")
	}
	return md.data[key], nil
}

// Update the Element if the key exist
func (md *MemoryDB) Update(key string, value interface{}) error {
	md.mtx.Lock()
	if md.data[key] != nil {
		md.data[key] = value
	} else {
		md.mtx.Unlock()
		return errors.New("No information was found with the key received")
	}
	md.mtx.Unlock()
	return md.SaveMapInFile()
}

//Delete an Element if the key exist
func (md *MemoryDB) Delete(key string) error {
	
	md.mtx.Lock()
	if md.data[key] != nil {
		delete(md.data, key)
	} else {
		md.mtx.Unlock()
		return errors.New("No information was found with the key received")
	}
	md.mtx.Unlock()
	return md.SaveMapInFile()
}

func (md *MemoryDB) SaveMapInFile() error {
	jsonString, _ := json.Marshal(md.data)
	if err := ioutil.WriteFile(Constants.FilePath, jsonString, 0644); err != nil {
		return err
	}
	return nil
}
func (md *MemoryDB) ReadMapFromFile() error {
	dataLikeBytes, err := ioutil.ReadFile(Constants.FilePath)
	if err = json.Unmarshal(dataLikeBytes, &md.data); err != nil {
		return err
	}
	return nil
}