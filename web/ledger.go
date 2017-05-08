package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"piego/db"
	"fmt"
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

func LedgerCreateHandler(w http.ResponseWriter, r *http.Request) {
	var entry db.LedgerEntry
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&entry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	id, err := db.LedgerEntryCreate(entry)
	http.Redirect(w, r, fmt.Sprintf("/ledger/entry/%d", id), 302)
}