package main

import (
	"fmt"
	"io"
	"net/http" 
	"os"
)

func getRoot( w http.ResponseWriter, r *http.Request) {
	fmt.Print("got / request\n")
	io.WriteString(w, "Hello")
}

func getHello(w, htttp.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
