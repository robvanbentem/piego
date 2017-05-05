package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"piego/db"
	"piego/web"
)

func main() {
	fmt.Printf("Starting Piego webserver\n")
	db.InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/shops", web.ShopsAllHandler).Methods("GET")

	r.HandleFunc("/users", web.UsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", web.UsersFindHandler).Methods("GET")

	r.HandleFunc("/shoplist/{date}", web.ShoplistDateHandler).Methods("GET")

	r.HandleFunc("/shoplist/entry", web.ShoplistCreateHandler).Methods("POST")
	r.HandleFunc("/shoplist/entry/{id}", web.ShoplistFindHandler).Methods("GET")
	r.HandleFunc("/shoplist/entry/{id}", web.ShoplistDeleteHandler).Methods("DELETE")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	panic(http.ListenAndServe("0.0.0.0:8004", handlers.CORS(originsOk, headersOk, methodsOk)(loggedRouter)))
}
