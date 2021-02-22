package main

import (
	"leapyear/year"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/leapYear", year.GetHandler).Methods(http.MethodGet)
	r.HandleFunc("/leapYear", year.PostHandler).Methods(http.MethodPost)
	http.Handle("/", r)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
