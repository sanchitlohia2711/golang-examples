package main

import (
	"fmt"
	"net/http"
	"strings"
)

func multiple() {
	getProductsHandler := http.HandlerFunc(getProducts2)
	http.Handle("/products", getProductsHandler)
	http.ListenAndServe(":8080", nil)
}

//HelloWorld hellow world handler
func getProducts2(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filters, present := query["filters"]

	if !present || len(filters) == 0 {
		fmt.Println("filters not present")
	}
	w.WriteHeader(200)
	w.Write([]byte(strings.Join(filters, ",")))
}
