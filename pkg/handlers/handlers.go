package handlers

import (
	"github.com/kskr24/go-site/pkg/render"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
