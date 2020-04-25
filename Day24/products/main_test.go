package main

import (
	"strconv"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"net/http"
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

func TestEmptyTable(t *testing.T){
	clearTable()

	req, _ := http.NewRequest("GET", "/products", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
        t.Errorf("Expected an empty array. Got %s", body)
    }
}

func TestGetNonExistentProduct(t *testing.T){
	clearTable()

	req, _ := http.NewRequest("GET", "/product/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
    json.Unmarshal(response.Body.Bytes(), &m)
    if m["error"] != "Product not found" {
        t.Errorf("Expected the 'error' key of the response to be set to 'Product not found'. Got '%s'", m["error"])
    }
}

func TestCreateProduct(t *testing.T){
	clearTable()

	body := []byte(`{"name":"iPad", "price": 399}`)
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	
	if m["name"] != "iPad" {
        t.Errorf("Expected product name to be 'test product'. Got '%v'", m["name"])
    }
    if m["price"] != 399.0 {
        t.Errorf("Expected product price to be '399'. Got '%v'", m["price"])
    }
}

func TestGetProduct(t *testing.T) {
    clearTable()
    addProducts(1)

    req, _ := http.NewRequest("GET", "/product/1", nil)
    response := executeRequest(req)

    checkResponseCode(t, http.StatusOK, response.Code)
}

func addProducts(count int) {
    if count < 1 {
        count = 1
    }

    for i := 0; i < count; i++ {
        a.DB.Exec("INSERT INTO products(name, price) VALUES($1, $2)", "Product "+strconv.Itoa(i), (i+1.0)*10)
    }
}

func TestUpdateProduct(t *testing.T){
	clearTable()
	addProducts(1)

	req, _ := http.NewRequest("GET", "/product/1", nil)
	response := executeRequest(req)
	var original map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &original)

	body := []byte(`{"name":"iPad Pro", "price": 999}`)
	req, _ = http.NewRequest("PUT", "/product/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var updated map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &updated)

	if original["name"] == updated["name"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", original["name"], updated["name"], updated["name"])
	}
	if original["price"] == updated["price"] {
		t.Errorf("Expected the price to change from '%v' to '%v'. Got '%v'", original["price"], updated["price"], updated["price"])
	}
}

func TestDeleteProduct(t *testing.T) {
    clearTable()
    addProducts(1)

    req, _ := http.NewRequest("GET", "/product/1", nil)
	response := executeRequest(req)
	
    checkResponseCode(t, http.StatusOK, response.Code)

    req, _ = http.NewRequest("DELETE", "/product/1", nil)
    response = executeRequest(req)

    checkResponseCode(t, http.StatusOK, response.Code)

    req, _ = http.NewRequest("GET", "/product/1", nil)
	response = executeRequest(req)

    checkResponseCode(t, http.StatusNotFound, response.Code)
}

func executeRequest(r *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, r)

	return rr
}

func checkResponseCode(t *testing.T, expectedStatusCode, actualStatusCode int){
	if expectedStatusCode != actualStatusCode {
        t.Errorf("Expected response code %d. Got %d\n", expectedStatusCode, actualStatusCode)
    }
}
