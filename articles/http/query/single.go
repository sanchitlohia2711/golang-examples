package main

import (
	"net/http"
)

func single() {
	getProductsHandler := http.HandlerFunc(getProducts)
	http.Handle("/products", getProductsHandler)
	http.ListenAndServe(":8080", nil)
}

//HelloWorld hellow world handler
func getProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	filters := query.Get("filters")
	w.WriteHeader(200)
	w.Write([]byte(filters))
}
