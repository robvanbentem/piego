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

func ShoplistFindHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	entry, err := db.ShoplistEntryFind(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err.Error())
		return
	}

	s, err := json.Marshal(entry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(s))
}

func ShoplistCreateHandler(w http.ResponseWriter, r *http.Request) {
	var e db.ShoplistEntry
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	id, err := db.ShoplistEntryCreate(e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/shoplist/entry/%d", id), 302)
}

func ShoplistDateHandler(w http.ResponseWriter, r *http.Request) {
	date := mux.Vars(r)["date"]

	entries, err := db.ShoplistForDate(date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s, err := json.Marshal(entries)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(s))
}

func ShoplistDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = db.ShoplistEntryDelete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ShoplistUpdateHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	var e db.ShoplistEntry
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&e)

	err = db.ShoplistEntryUpdate(id, e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/shoplist/entry/%d", id), 302)
}
