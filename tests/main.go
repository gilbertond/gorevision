package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func calculate(val int) (res int) {
	res = val * 2
	return res
}

func generateCustomFields(requestBody string, serverUrl string, token string) (*http.Response, error) {

	request, reqError := http.NewRequest(
		"POST",
		serverUrl,
		strings.NewReader(requestBody),
	)

	if reqError != nil {
		log.Fatalf("Internal error: %s", reqError)
	}

	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	response, respError := http.DefaultClient.Do(request)

	if respError != nil {
		log.Fatalf("Remote Server returned error: %s", respError)
		return nil, respError
	}

	if response.StatusCode != http.StatusOK {
		//Read content from response
		body, _ := ioutil.ReadAll(response.Body)
		bodyString := string(body)
		log.Println(bodyString)

		log.Fatalf("Remote Server returned error code: %d", response.StatusCode)
		return nil, respError
	}

	response.Body.Close()
	return response, nil
}
