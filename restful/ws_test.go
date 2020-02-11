package main

import (
	"fmt"
	"testing"
)


func TestCreate(t *testing.T) {

	data := db.NewMemoryDB()

	if err := data.Create("test1", "50"); err != nil {
		t.Errorf("Create was incorrect, got this error: %s", err)
	}
	if data.Len() != 1 {
		t.Errorf("Create was incorrect, got: %d, want: %d.", data.Len(), 1)
	}
}