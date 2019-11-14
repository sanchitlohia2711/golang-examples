package main

import (
	"fmt"
	"net/http"
	"strings"
)

func formvaluemultiple() {
	getProductsHandler := http.HandlerFunc(getProducts3)
	http.Handle("/products", getProductsHandler)
	http.ListenAndServe(":8080", nil)
}

//HelloWorld hellow world handler
func getProducts3(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filters, present := r.Form["filters"] //filters=["color", "price", "brand"]

	if !present || len(filters) == 0 {
		fmt.Println("filters not present")
	}
	w.WriteHeader(200)
	w.Write([]byte(strings.Join(filters, ",")))
}
