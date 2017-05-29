package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"piego/db"
	"piego/web"
	"piego/ws"
)

type config struct {
	ServerAddress string
	ServerPort    int
	DBHost        string
	DBPort        int
	DBUser        string
	DBPass        string
	DBScheme      string
}

func main() {
	cfg := loadConfig()

	fmt.Printf("Starting Piego webserver at %s:%d\n", cfg.ServerAddress, cfg.ServerPort)
	db.InitDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBScheme)
	defer db.CloseDB()

	ws.Init()

	r := mux.NewRouter()
	registerRoutes(r)

	// middleware
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	corsHandler := getCorsHandler(loggedRouter)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.ServerAddress, cfg.ServerPort), corsHandler))
}

func registerRoutes(r *mux.Router) {
	r.HandleFunc("/shops", web.ShopsAllHandler).Methods("GET")
	r.HandleFunc("/users", web.UsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", web.UsersFindHandler).Methods("GET")
	r.HandleFunc("/items", web.ItemsAllHandler).Methods("GET")
	r.HandleFunc("/items/search", web.ItemsSearchHandler).Methods("POST")
	r.HandleFunc("/shoplist/{date}", web.ShoplistDateHandler).Methods("GET")
	r.HandleFunc("/shoplist/entry", web.ShoplistCreateHandler).Methods("POST")
	r.HandleFunc("/shoplist/entry/{id}", web.ShoplistFindHandler).Methods("GET")
	r.HandleFunc("/shoplist/entry/{id}", web.ShoplistDeleteHandler).Methods("DELETE")
	r.HandleFunc("/shoplist/entry/{id}", web.ShoplistUpdateHandler).Methods("PUT")
	r.HandleFunc("/ledger", web.LedgerAllHandler).Methods("GET")
	r.HandleFunc("/ledger/entry", web.LedgerEntryCreateHandler).Methods("POST")
	r.HandleFunc("/ledger/entry/{id}", web.LedgerEntryUpdateHandler).Methods("PUT")
	r.HandleFunc("/ledger/{date}", web.LedgerDateHandler).Methods("GET")

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.WSHandler(ws.GetHub(), w, r)
	})
}

func getCorsHandler(loggedRouter http.Handler) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	return handlers.CORS(originsOk, headersOk, methodsOk)(loggedRouter)
}

func loadConfig() config {
	f, err := os.Open("config.json")
	if err != nil {
		panic("No config.json found or not readable.")
	}

	cfg := config{
		ServerAddress: "0.0.0.0",
		ServerPort:    8004,
	}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic("Could not parse config.json.example, " + err.Error())
	}

	return cfg
}
