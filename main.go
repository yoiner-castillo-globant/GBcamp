package main

import (
	"github.com/yoiner-castillo-globant/GBcamp/io"
	"github.com/yoiner-castillo-globant/GBcamp/db"
)


func main() {
io.ReadMapFromFile()
db.PrintDatos()


db.Create("23", 23445)
db.Create("532", "Polo")
db.Create("Jhoi", "temporal")
io.SaveMapInFile();
db.PrintDatos()


}
