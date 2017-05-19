package web

import (
	"encoding/json"
	"net/http"
	"piego/db"
)

func ShopsAllHandler(w http.ResponseWriter, r *http.Request) {
	shops, err := db.ShopsAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	shopsJson, err := json.Marshal(shops)
	w.Write([]byte(shopsJson))
}
