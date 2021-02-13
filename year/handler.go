package year

import (
	"log"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	strYear, err := strconv.ParseUint(r.URL.Query().Get("year"), 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := "Year is leap"
	if !IsLeapYear(strYear) {
		response = "Year isn't leap"
	}
	w.Write([]byte(response))
}
