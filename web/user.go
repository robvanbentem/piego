package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"piego/db"
	"strconv"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users := db.UsersAll()

	for idx, _ := range users {
		users[idx].FetchBalance()
	}

	s, err := json.Marshal(users)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(s))
}

func UsersFindHandler(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, _ := strconv.ParseInt(idStr, 10, 64)
	user, err := db.UsersFind(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userJson, _ := json.Marshal(user)
	w.Write([]byte(userJson))
}
