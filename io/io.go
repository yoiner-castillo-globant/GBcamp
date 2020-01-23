package io

import (
	"encoding/json"
	"github.com/yoiner-castillo-globant/GBcamp/db"
	"io/ioutil"
	"log"
)

func SaveMapInFile() {
	jsonString, _ := json.Marshal(db.Datos)
	err := ioutil.WriteFile("./io/Info.txt", jsonString, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadMapFromFile() {
	datosComoBytes, err := ioutil.ReadFile("./io/Info.txt")
	if err != nil {
	//	fmt.Println("[Informativo, no error] No existe el archivo a leer...")
	} else {
		err = json.Unmarshal(datosComoBytes, &db.Datos)
		if err != nil {
			panic(err)
		}
	}
}
