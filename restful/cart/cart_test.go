package cart

import (
	"fmt"
	"github.com/yoiner-castillo-globant/GBcamp/restful/cart"
	"testing"
)

func TestAddItem(t *testing.T) {

	icart := cart.CreateCart()
	icart.AddItem(1, 3)

	if 1 != 1 {
		t.Errorf("Create was incorrect, got this error: %s", err)
	}


}