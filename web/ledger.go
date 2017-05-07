package web

import (
	"net/http"
	"piego/db"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
)

func LedgerAllHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := db.LedgerAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	s, err := json.Marshal(entries)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	w.Write(s)
}

func LedgerDateHandler(w http.ResponseWriter, r *http.Request) {
	date := mux.Vars(r)["date"]

	entries, err := db.LedgerForDate(date)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	s, err := json.Marshal(entries)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	w.Write(s)
}

