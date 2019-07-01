package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculate(t *testing.T) {
	res := calculate(2)

	if res != 4 {
		t.Errorf("Expected 4 but got %d", res)
	}
}

func TestCalculateWithTestTable(t *testing.T) {
	var testValues = []struct {
		input    int
		expected int
	}{
		{1, 2},
		{8, 16},
		{3, 6},
		{5, 10},
	}

	for _, test := range testValues {
		res := calculate(test.input)
		if test.expected != res {
			t.Errorf("Expected %d but got %d", test.expected, res)
		}
	}
}

func TestCreateCustomFieldsMock(t *testing.T) {

	mockRequest := `{
					  "id": -737373,
					  "name": "some name",
					  "type": "some type",
					  "extraField1": "field1",
					  "extraField2": 33,
					  "extraField3": 44.11,
					  "extraField4": true,
					   "extraField5":{ "innerField5":2}
					}`

	handler := func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, mockRequest)
	}

	httpRequest := httptest.NewRequest("POST", "https://api.sendgrid.com/v3/contactdb/custom_fields", nil)

	httpWriter := httptest.NewRecorder()

	handler(httpWriter, httpRequest)

	response := httpWriter.Result()
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))

	if 200 != response.StatusCode {
		t.Fatalf("Remote Server returned error: %d", response.StatusCode)
	}
}

func TestCreateCustomFieldsReturns200(t *testing.T) {
	requestBody := fmt.Sprintf(`{
					  "id": -737373,
					  "name": "some name",
					  "type": "some type",
					  "extraField1": "field1",
					  "extraField2": 33,
					  "extraField3": 44.11,
					  "extraField4": true,
 					  "extraField5":{ "innerField5":2}
					}`,
	)

	mockResponseBody := fmt.Sprintf(`{
					  "id": -737373,
					  "name": "some name",
					  "type": "some type",
					  "extraField1": "field1",
					  "extraField2": 33,
					  "extraField3": 44.11,
					  "extraField4": true,
 					  "extraField5":{ "innerField5":2}
					}`,
	)

	token := "someToken"

	request, reqError := http.NewRequest(
		"POST",
		"http://server/contactdb/custom_fields",
		strings.NewReader(requestBody),
	)

	if reqError != nil {
		log.Fatalf("Ooops some Error: %s", reqError)
	}

	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	mockServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			r = request

			io.WriteString(w, mockResponseBody)

			if r.Method != "POST" {
				t.Errorf("Expected 'POST' request but was %s ", r.Method)
			}

			if r.URL.EscapedPath() != "/contactdb/custom_fields" {
				t.Errorf("Expected endpoint '/contactdb/custom_fields' but was '%s'", r.URL.EscapedPath())
			}

			r.ParseForm()
			reqJson, err := simplejson.NewFromReader(r.Body)

			if err != nil {
				t.Errorf("Invalid Json request body!! ****%s", err)
			}

			requestField := reqJson.GetPath("extraField1").MustString()
			if requestField != "field1" {
				t.Errorf("Expected field 'extraField1=field1' but found '%s'", requestField)
			}

			if requestField != "field1" {
				t.Errorf("Expected field 'extraField1=field1' but found '%s'", requestField)
			}

			//Can you try to match request request and response fields to make sure the API created the expected custom fields
			requestFieldId := reqJson.Get("id").MustInt()
			if requestFieldId != -737373 {
				t.Errorf("Expected to create a custom 'id' of value '%d' but '%s'", -737373, string(requestFieldId))
			}

			//Nested fields
			nestedRequestField := reqJson.GetPath("extraField5", "innerField5").MustInt()
			if nestedRequestField != 2 {
				t.Errorf("Expected to create a custom 'extraField5/field5' of value '%d' but '%s'", 2, string(nestedRequestField))
			}

			//response, respError := http.DefaultClient.Do(r)
			//if respError != nil {
			//	log.Fatalf("Oooops some error: %s", respError)
			//}
			//resoBody, _ := simplejson.NewFromReader(response.Body)
			//responseFieldId := resoBody.Get("id").MustInt()
			//if responseFieldId != -737373 && requestFieldId != responseFieldId {
			//	t.Errorf("Expected to create a custom 'id' of value '%d' but '%s'", -737373, string(responseFieldId))
			//}
		}))
	defer mockServer.Close()

	response, err := generateCustomFields(requestBody, mockServer.URL, token)

	if err != nil {
		t.Errorf("CreateCustomFields service returned an error: %s", err)
	}

	if response == nil {
		t.Errorf("CreateCustomFields service returned null response")
	}
}

func TestCreateCustomFieldsIntegration(t *testing.T) {

	requestBody := fmt.Sprintf(`{
					  "id": -737373,
					  "name": "some name",
					  "type": "some type",
					  "extraField1": "field1",
					  "extraField2": 33,
					  "extraField3": 44.11,
					  "extraField4": true
					}`,
	)

	serverUrl := "https://api.sendgrid.com/v3/contactdb/custom_fields"

	token := "token"
	response, error := generateCustomFields(requestBody, serverUrl, token)

	if error != nil {
		t.Fatalf("Remote Server returned error: %s", error)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatalf("Remote Server returned error code: %d", response.StatusCode)
	}

}
