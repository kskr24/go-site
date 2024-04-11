package main

import (
	"fmt"
	"github.com/kskr24/go-site/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	// fmt.Println("Hello World")
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application at port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
