package main

import "github.com/yoiner-castillo-globant/GBcamp/db"


func main() {

	data := db.NewMemoryDB()
	data.ReadMapFromFile()
	data.PrintDATA()
	data.Create("23", 23445)
	data.Create("532", "Polo")
	data.Create("Jhoi", "temporal")
	data.PrintDATA()
	data.Update("23", "una cadena")
	data.PrintDATA()
	data.Delete("532")
	data.PrintDATA()
	data.SaveMapInFile();

}
