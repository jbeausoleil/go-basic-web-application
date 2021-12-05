package main

import (
	"fmt"
	"github.com/jbeausoleil/go-basic-web-application/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port http://localhost%s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
