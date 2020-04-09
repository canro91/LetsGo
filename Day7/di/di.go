package main

import (
	"net/http"
	"io"
	"fmt"
)

func Greet(writer io.Writer, name string){
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHander(w http.ResponseWriter, r *http.Request){
	Greet(w, "World")
}

func main(){
	// 1. Write to Stdout
	// Greet(os.Stdout, "World")
	// 2. Write to Web
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHander))
}