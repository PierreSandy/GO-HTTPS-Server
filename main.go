package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", getRoot)       //calls get root func at "/"
	http.HandleFunc("/hello", getHello) //calls get hello func at "/hello"

	err := http.ListenAndServe(": 9097", nil) //Listens on Port 9097, will close program after server told to shutdown
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server is now closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Print("recieved / request\n")
	io.WriteString(w, "Hello\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("recieved /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
