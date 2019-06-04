package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	data "./data"
)

// New .
func New() http.Handler {
	r := mux.NewRouter()

	// website
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Println("Hello") }).Methods("GET")

	// product
	r.HandleFunc("/data", data.GetAll).Methods("GET")
	r.HandleFunc("/data/{id}", data.Get).Methods("GET")
	r.HandleFunc("/data", data.New).Methods("POST")
	r.HandleFunc("/data", data.Edit).Methods("PUT")
	r.HandleFunc("/data/{id}", data.Delete).Methods("DELETE")

	return r
}
