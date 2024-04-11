package main

import (
	"fmt"
	"github.com/kskr24/go-site/pkg/config"
	"github.com/kskr24/go-site/pkg/handlers"
	"github.com/kskr24/go-site/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create Template Cache")
	}
	app.TemplateCache = tc
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application at port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
