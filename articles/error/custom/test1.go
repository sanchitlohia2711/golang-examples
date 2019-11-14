package main

import (
	"net/http"
)

type httpInputValidationError struct {
	statusCode int
	message    string
}

func (h *httpInputValidationError) Error() string {
	return h.message
}

func test1() {
	addUserHandler := http.HandlerFunc(addUser)
	http.Handle("/addUser", addUserHandler)
	http.ListenAndServe(":8080", nil)
}

//HelloWorld hellow world handler
func addUser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	gender := r.URL.Query().Get("gender")
	err := validate(name, gender)
	if err != nil {
		if err, ok := err.(*httpInputValidationError); ok {
			w.WriteHeader(err.statusCode)
			w.Write([]byte(err.Error()))
			return
		}
	}
	w.Write([]byte("Welcome " + name))
	w.WriteHeader(200)
}

func dovalidate(name, gender string) error {
	if name == "" {
		return &httpInputValidationError{statusCode: 400, message: "Name is mandatory"}
	}
	if gender == "" {
		return &httpInputValidationError{statusCode: 400, message: "Gender is mandatory"}
	}
	return nil
}
