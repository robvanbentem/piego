package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"piego/db"
	"strconv"
)

type entry struct {
	user_id int
	shop_id int
	name    string
	qty     int
	date    string
}

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
	decoder := json.NewDecoder(r.Body)

	var e entry
	err := decoder.Decode(&e)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, err := db.ShoplistEntryCreate(e.user_id, e.shop_id, e.name, e.qty, e.date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/shoplist/entry/"+string(id), 302)
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
