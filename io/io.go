package io

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
func Escribir() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}*/

func EscribirArchivo(info interface{}, nombreArchivo string) {
	file, err := os.Create("./io/" + nombreArchivo + ".txt")
	if err != nil {
		return
	}
	defer file.Close()
	data := fmt.Sprintf("%v", info)
	file.WriteString(data)
}

func LeerArchivo(nombreArchivo string) {
	file, err := os.Open("./io/" + nombreArchivo + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file) // default line by line
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
