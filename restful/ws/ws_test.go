package ws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/yoiner-castillo-globant/GBcamp/restful/ws"
	"github.com/yoiner-castillo-globant/GBcamp/structs"
)

func TestNewCart(t *testing.T) {
	tt := []struct {
		name string
		err  string
	}{
		{name: "Normal"},
	}
	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {

			r, err := http.NewRequest("GET", "/NewCart", nil)
			if err != nil {
				t.Errorf("Could not create request %v", err)
			}

			w := httptest.NewRecorder()

			ws.NewCart(w, r)

			resp := w.Result()
			defer resp.Body.Close()

			b, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				t.Fatalf("Could not read response: %v", err)
			}

			if tc.err != "" {
				if resp.StatusCode != http.StatusOK {
					t.Errorf("Expected status bad request; got %v", resp.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, msg)
				}
				return
			}

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status Ok; got %v", resp.Status)
			}

			data := structs.ResponseNewCartWs{}

			if err = json.Unmarshal(b, &data); err != nil {
				t.Fatalf("Could not encoded the response in a particular structure: %v", err)
			} else if reflect.TypeOf(data.CartId).Kind() != reflect.String {
				t.Fatalf("Expected a string value: %v", data.CartId)
			}

		})
	} //End for

}

func getNewCart() string {
	r, _ := http.NewRequest("GET", "/NewCart", nil)

	w := httptest.NewRecorder()
	ws.NewCart(w, r)
	resp := w.Result()
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	data := structs.ResponseNewCartWs{}

	json.Unmarshal(b, &data)

	return data.CartId
}

func TestAddItem(t *testing.T) {

	IdCart := getNewCart()

	tt := []struct {
		name   string
		cartId string
		values string
		err    string
	}{
		{name: "With out {cartId}", cartId: "", values: "4", err: "you need to send an cartId"},
		{name: "With out {articleId}", cartId: IdCart, values: "", err: "you need to send an articleId"},
		{name: "With out {articleId} valid", cartId: IdCart, values: "x", err: "you need to send an articleId valid"},
		{name: "Successfull", cartId: IdCart, values: "2"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			data := url.Values{}
			data.Set("cartId", tc.cartId)
			data.Set("articleId", tc.values)

			r, err := http.NewRequest("POST", "AddItem", strings.NewReader(data.Encode()))
			if err != nil {
				t.Errorf("Could not create request %v", err)
			}
			//r.Header.Set("Content-Type", "application/json")
			r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			w := httptest.NewRecorder()
			ws.AddItem(w, r)

			resp := w.Result()
			defer resp.Body.Close()

			b, err := ioutil.ReadAll(resp.Body)
			dataStruct := structs.BadResponseAddItem{}
			if err = json.Unmarshal(b, &dataStruct); err != nil {
				t.Fatalf("Could not encoded the response in a particular structure: %v", err)
			}

			if err != nil {
				t.Fatalf("Could not read response: %v", err)
			}

			if tc.err != "" {
				if resp.StatusCode != http.StatusOK {
					t.Errorf("Expected status bad request; got %v", resp.StatusCode)
				}
				fmt.Println("Respuesta: ", dataStruct.Response)
				if dataStruct.Response != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, dataStruct.Response)
				}
				return
			}

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status Ok; got %v", resp.Status)
			}

		})
	} //End for

}
