package cart

import (
	"fmt"
	"github.com/yoiner-castillo-globant/GBcamp/restful/cart"
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