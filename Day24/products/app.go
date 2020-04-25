package main

import (
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
}

func (a *App) Run(address string) {

}
