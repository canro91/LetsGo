package main

import (
	"encoding/json"
	"strconv"
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() {
	var err error
	a.DB, err = sql.Open("sqlite3", "products.db")
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

func (a *App) InitializeRoutes(){
	a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
}

func (a *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, a.Router))
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
    count, _ := strconv.Atoi(r.FormValue("count"))
    start, _ := strconv.Atoi(r.FormValue("start"))

    if count > 10 || count < 1 {
        count = 10
    }
    if start < 0 {
        start = 0
    }

    products, err := getProducts(a.DB, start, count)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, products)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}
