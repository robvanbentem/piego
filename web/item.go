package web

import (
	"encoding/json"
	"log"
	"net/http"
	"piego/db"
)

type ItemSearchQry struct {
	ShopID int64
	Name   string
}

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
	var searchQry ItemSearchQry
	d := json.NewDecoder(r.Body)

	err := d.Decode(&searchQry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	items, err := db.ItemsSearch(searchQry.ShopID, searchQry.Name)
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
