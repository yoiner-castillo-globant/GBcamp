package db

import (
	"fmt"
	"github.com/yoiner-castillo-globant/GBcamp/db"
	"testing"
)

func TestCreate(t *testing.T) {

	data := db.NewMemoryDB()

	err := data.Create("test1", "50")
	if err != nil {
		t.Errorf("Create was incorrect, got this error: %s", err)
	}

	if data.Len() != 1 {
		t.Errorf("Create was incorrect, got: %d, want: %d.", data.Len(), 1)
	}
}

func TestRetrieve(t *testing.T) {
	data := db.NewMemoryDB()
	data.Create("test2", "Works")
	x, err := data.Retrieve("test2")
	value := fmt.Sprintf("%v", x)

	if err != nil {
		t.Errorf("Retrieve was incorrect, got this error: %s, ", err)

	}
	if value != "Works" {
		t.Errorf("Retrieve was incorrect, got: %s, want: %s.", value, "Funciona")
	}
}

func TestUpdate(t *testing.T) {
	data := db.NewMemoryDB()

	data.Create("test3", "works")
	err := data.Update("test3", "works2")
	x, _ := data.Retrieve("test3")
	dato := fmt.Sprintf("%v", x)

	if err != nil {
		t.Errorf("Retrieve was incorrect, got this error: %s, ", err)
	}
	if dato != "works2" {
		t.Errorf("Update was incorrect, got: %s, want: %s.", dato, "works2")
	}
}

func TestDelete(t *testing.T) {
	data := db.NewMemoryDB()

	data.Create("test4", "Deleting Test")
	previousAmount := data.Len()
	err := data.Delete("test4")
	postAmount := data.Len()
	previousAmount--

	if err != nil {
		t.Errorf("Update was incorrect, got this error: %s", err)
	}
	if previousAmount != postAmount {
		t.Errorf("Update was incorrect, got: %d, want: %d.", previousAmount, postAmount)
	}
}
