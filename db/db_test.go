package db

import (
	"fmt"
	"testing"
	"github.com/yoiner-castillo-globant/GBcamp/db"
)

func TestCreate(t *testing.T) {
	db.Create("test1", "50")
	length := len(db.Datos)

	if length > 1 {
		t.Errorf("Create was incorrect, got: %d, want: %d.", length, 1)
	}
}

func TestRetrieve(t *testing.T) {
	db.Create("test2", "Funciona")
	var x interface{} = db.Retrieve("test2")
	dato := fmt.Sprintf("%v", x)

	if dato != "Funciona" {
		t.Errorf("Retrieve was incorrect, got: %s, want: %s.", dato, "Funciona")
	}
}

func TestUpdate(t *testing.T) {
	db.Create("test3", "Funciona")
	db.Update("test3", "Funciona2")
	var x interface{} = db.Retrieve("test3")
	dato := fmt.Sprintf("%v", x)

	if dato != "Funciona2" {
		t.Errorf("Update was incorrect, got: %s, want: %s.", dato, "Funciona2")
	}
}

func TestDelete(t *testing.T) {
	db.Create("test4", "Eliminar Test")
	cantidadAnterior := len(db.Datos)
	db.Delete("test4")
	cantidadPosterior := len(db.Datos)
	cantidadAnterior--

	if cantidadAnterior != cantidadPosterior {
		t.Errorf("Update was incorrect, got: %d, want: %d.", cantidadAnterior, cantidadPosterior)
	}
}