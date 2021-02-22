package year

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
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

func PostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rec := struct {
		Year uint64
	}{}
	err = json.Unmarshal(body, &rec)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := "Year is leap"
	if !IsLeapYear(rec.Year) {
		response = "Year isn't leap"
	}
	w.Write([]byte(response))
}
