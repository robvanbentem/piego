package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"piego/db"
	"strconv"
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

func LedgerEntryCreateHandler(w http.ResponseWriter, r *http.Request) {
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

func LedgerEntryUpdateHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		log.Print(err.Error())
		return
	}

	var e db.LedgerEntry
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&e)

	err = db.LedgerEntryUpdate(id, e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
