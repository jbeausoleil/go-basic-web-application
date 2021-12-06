package main

import (
	"fmt"
	"github.com/jbeausoleil/go-basic-web-application/pkg/config"
	"github.com/jbeausoleil/go-basic-web-application/pkg/handlers"
	"github.com/jbeausoleil/go-basic-web-application/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	app := config.AppConfig{}

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port http://localhost%s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
