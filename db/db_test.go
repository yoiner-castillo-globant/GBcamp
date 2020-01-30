package db

import (
	"fmt"
	"github.com/yoiner-castillo-globant/GBcamp/db"
	"testing"
)

func TestCreate(t *testing.T) {
	err := db.Create("test1", "50")
	if err != nil {
		t.Errorf("Create was incorrect, got this error: %s", err)
	}

	length := len(db.DATA)

	if length == 1 {
		t.Errorf("Create was incorrect, got: %d, want: %d.", length, 1)
	}
}

func TestRetrieve(t *testing.T) {
	db.Create("test2", "Works")
	x, err := db.Retrieve("test2")
	value := fmt.Sprintf("%v", x)

	if err != nil {
		t.Errorf("Retrieve was incorrect, got this error: %s, ", err)

	}
	if value != "Works" {
		t.Errorf("Retrieve was incorrect, got: %s, want: %s.", value, "Funciona")
	}
}

func TestUpdate(t *testing.T) {
	db.Create("test3", "works")
	err := db.Update("test3", "works2")
	x, _ := db.Retrieve("test3")
	dato := fmt.Sprintf("%v", x)

	if err != nil {
		t.Errorf("Retrieve was incorrect, got this error: %s, ", err)
	}
	if dato != "works2" {
		t.Errorf("Update was incorrect, got: %s, want: %s.", dato, "works2")
	}
}

func TestDelete(t *testing.T) {
	db.Create("test4", "Deleting Test")
	previousAmount := len(db.DATA)
	err := db.Delete("test4")
	postAmount := len(db.DATA)
	previousAmount--

	if err != nil {
		t.Errorf("Update was incorrect, got this error: %s", err)
	}
	if previousAmount != postAmount {
		t.Errorf("Update was incorrect, got: %d, want: %d.", previousAmount, postAmount)
	}
}
