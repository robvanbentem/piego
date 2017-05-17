package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"piego/db"
	"strconv"
)

func ItemsAllHandler(w http.ResponseWriter, r *http.Request) {
	items, err := db.ItemsAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	s, err := json.Marshal(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write(s)
}

func ItemsSearchHandler(w http.ResponseWriter, r *http.Request) {
	shopId, err := strconv.ParseInt(mux.Vars(r)["shopId"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	search := mux.Vars(r)["qry"]

	items, err := db.ItemsSearch(shopId, search)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	s, err := json.Marshal(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Write(s)
}
