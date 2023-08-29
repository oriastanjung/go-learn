package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/oriastanjung/go-course/pkg/config"
	"github.com/oriastanjung/go-course/pkg/handlers"
	"github.com/oriastanjung/go-course/pkg/render"
)

const portNumber = ":8080"

func main() {
	// create a template cache
	var app config.AppConfig
	tc, err := render.CreateTemplateCacheForComplex()
	if err != nil {
		log.Fatal("cannot create template cache ")
	}

	app.TemplateCache = tc
	// karena masih dev mode kita set false supaya ketika refresh kita rebuild cache nya
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	// cara lama
	// http.HandleFunc("/", handlers.Home)
	// http.HandleFunc("/about", handlers.About)

	fmt.Printf("Running on localhost%s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
