package io

import (
	"encoding/json"
	"github.com/yoiner-castillo-globant/GBcamp/db"
	"io/ioutil"
	"log"
)

func SaveMapInFile() {
	jsonString, _ := json.Marshal(db.DATA)
	err := ioutil.WriteFile("./io/Info.txt", jsonString, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadMapFromFile() {
	datosComoBytes, err := ioutil.ReadFile("./io/Info.txt")
	if err == nil {
		err = json.Unmarshal(datosComoBytes, &db.DATA)
		if err != nil {
			panic(err)
		}
	}
}
