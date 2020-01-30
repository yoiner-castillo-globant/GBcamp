package db

import (
	"fmt"
	"github.com/yoiner-castillo-globant/GBcamp/db"
	"testing"
)

func TestCreate(t *testing.T) {
	resp, err := db.Create("test1", "50")
	if !resp {
		t.Errorf("Create was incorrect, got this error: %s", err)
	}

	length := len(db.Datos)

	if length == 1 {
		t.Errorf("Create was incorrect, got: %d, want: %d.", length, 1)
	}
}

func TestRetrieve(t *testing.T) {
	db.Create("test2", "Funciona")
	x, err := db.Retrieve("test2")
	dato := fmt.Sprintf("%v", x)

	if err != nil {
		t.Errorf("Retrieve was incorrect, got this error: %s, ", err)

	}
	if dato != "Funciona" {
		t.Errorf("Retrieve was incorrect, got: %s, want: %s.", dato, "Funciona")
	}
}

func TestUpdate(t *testing.T) {
	db.Create("test3", "Funciona")
	val, err := db.Update("test3", "Funciona2")
	x, _ := db.Retrieve("test3")
	dato := fmt.Sprintf("%v", x)

	if !val {
		t.Errorf("Retrieve was incorrect, got this error: %s, ", err)

	}
	if dato != "Funciona2" {
		t.Errorf("Update was incorrect, got: %s, want: %s.", dato, "Funciona2")
	}
}

func TestDelete(t *testing.T) {
	db.Create("test4", "Eliminar Test")
	cantidadAnterior := len(db.Datos)
	val, err := db.Delete("test4")
//Incorrect
	//	val, err := db.Delete("test")

	cantidadPosterior := len(db.Datos)
	cantidadAnterior--

	if !val {
		t.Errorf("Update was incorrect, got this error: %s", err)
	}
	if cantidadAnterior != cantidadPosterior {
		t.Errorf("Update was incorrect, got: %d, want: %d.", cantidadAnterior, cantidadPosterior)
	}
}
