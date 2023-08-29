package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/oriastanjung/go-course/pkg/config"
)

// kurang efisien karena read terus dri setiap request
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template : ", err)
	}
}

// cara render lebih bagus

var tc = make(map[string]*template.Template)

func RenderTemplateMoreEficient(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		// need to create the template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}

	} else {
		// template already in cache
		log.Println("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	// parse the template

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// add template to cache

	tc[t] = tmpl
	return nil
}

// render eficient tapi lebih complex
// disini kita gaperlu nambahin lagi klo ada base layout baru atau template layout baru karena otomatis read dari sistem

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplateMoreEficientAndComplex(w http.ResponseWriter, tmpl string) {
	var templateCache map[string]*template.Template

	if app.UseCache {
		// get template cache dari app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCacheForComplex()
	}

	// get requested template from cache
	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buffer := new(bytes.Buffer)

	_ = t.Execute(buffer, nil)

	// render the template
	_, err := buffer.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser ", err)
	}
}

func CreateTemplateCacheForComplex() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from "./templates"

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl

	for _, page := range pages {
		filename := filepath.Base(page)
		ts, err := template.New(filename).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[filename] = ts
	}

	return myCache, nil

}
