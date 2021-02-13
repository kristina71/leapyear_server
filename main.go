package main

import (
	"leapyear/year"
	"net/http"
)

func main() {
	http.HandleFunc("/leapYear", year.Handler)
	http.ListenAndServe(":3000", nil)
}
