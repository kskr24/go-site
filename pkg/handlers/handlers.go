package handlers

import (
	"github.com/kskr24/go-site/pkg/render"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello")
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
