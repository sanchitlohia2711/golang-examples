package main

import (
	"net/http"
)

func formvaluesingle() {
	getProductsHandler := http.HandlerFunc(getProducts4)
	http.Handle("/products", getProductsHandler)
	http.ListenAndServe(":8080", nil)
}

//HelloWorld hellow world handler
func getProducts4(w http.ResponseWriter, r *http.Request) {
	filters := r.FormValue("filters") //filters=["color"]
	w.WriteHeader(200)
	w.Write([]byte(filters))
}
