package ApiRest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	//"github.com/stretchr/testify/assert"
	"github.com/yoiner-castillo-globant/GBcamp/App/ApiRest"
	"github.com/yoiner-castillo-globant/GBcamp/App/Request"
)
 func TestNewCart(t *testing.T) {
 	tt := []struct {
 		name string
 		err  string
 	}{
 		{name: "Successfull"},
 	}
 	for _, tc := range tt {
 		t.Run(tc.name, func(t *testing.T) {
 			request, err := http.NewRequest("GET", "/CreateCart", nil)
 			if err != nil {
 				t.Errorf("Could not create request %v", err)
 			}
 			response := httptest.NewRecorder()
 			ApiRest.LoadEndPoints().ServeHTTP(response, request)
 			resp := response.Result()
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
 			if resp.StatusCode != http.StatusCreated {
 				t.Errorf("Expected status Created; got %v", resp.Status)
 			}
 			data := Request.TestResponse{}
 			if err = json.Unmarshal(b, &data); err != nil {
 				t.Fatalf("Could not encoded the response in a particular structure: %v", err)
 			} else if reflect.TypeOf(data.CartId).Kind() != reflect.String {
 				t.Fatalf("Expected a string value: %v", data.CartId)
 			}
 		})
 	} //End for
 }

func getNewCart() string {
	r, _ := http.NewRequest("GET", "/CreateCart", nil)

	w := httptest.NewRecorder()
	ApiRest.NewCartEP(w, r)
	resp := w.Result()
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	data := Request.TestResponse{}

	json.Unmarshal(b, &data)
	return data.CartId
}

func TestAddItem(t *testing.T) {

	IdCart := getNewCart()

	tt := []struct {
		name   string
		method string
		cartId string
		values string
		err    string
	}{
		{name: "With out {CartId}", method: http.MethodPost, cartId: "", values: "4", err: "you need to send an CartId"},
		{name: "With out {ArticleId}", method: http.MethodPost, cartId: IdCart, values: "", err: "you need to send an ArticleId"},
		{name: "With out {ArticleId} valid", method: http.MethodPost, cartId: IdCart, values: "x", err: "you need to send an ArticleId valid"},
		{name: "Successfull", method: http.MethodPost, cartId: IdCart, values: "2"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			path := fmt.Sprintf("/AddItem/%s", tc.cartId)

			data := &Request.PostArticle{ ArticleId: tc.values}
			jsonData, _ := json.Marshal(data)

			request, err := http.NewRequest(tc.method, path, bytes.NewBuffer(jsonData))
			if err != nil {
				t.Errorf("Could not create request %v", err)
			}
			response := httptest.NewRecorder()

			ApiRest.LoadEndPoints().ServeHTTP(response, request)
			
			//assert.Equal(t, 200, response.Code, "OK response is expected")

			resp := response.Result()
			defer resp.Body.Close()

			b, err := ioutil.ReadAll(resp.Body)
			dataStruct := Request.TestResponse{}
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

				if dataStruct.Response != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, dataStruct.Response)
				}
				return
			}

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status Created; got %v", resp.Status)
			}

		})
	}
}
