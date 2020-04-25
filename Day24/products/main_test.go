package main

import (
	"testing"
	"log"
	"os"
)

var a App

func TestMain(m *testing.M) {
	a.Initialize()

	createTable()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func createTable() {
	_, err := a.DB.Exec(`CREATE TABLE IF NOT EXISTS products
	(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price DECIMAL(10,2) NOT NULL DEFAULT 0.00
	)`)
	if err != nil {
		log.Fatal(err)
	}

}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}
