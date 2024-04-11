package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates:", err)
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in the cache
	_, inMap := tc[t]
	if !inMap {
		// need to create it
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Using cached template")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
}

func createTemplateCache(t string) error {
	templates := []string{fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	// parse the templates
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = tmpl
	return nil
}
