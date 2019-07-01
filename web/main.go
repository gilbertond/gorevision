package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", getHomePage)
	http.ListenAndServe(":3000", nil)
}

func getHomePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "data back")
}