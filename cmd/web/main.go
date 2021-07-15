package main

import (
	"github.com/Igor-Koniukhov/cacheSample/pkg/config"
	"github.com/Igor-Koniukhov/cacheSample/pkg/handlers"
	"github.com/Igor-Koniukhov/cacheSample/pkg/render"

	"log"
	"net/http"
)
var Port = ":8080"
var app config.AppConfig
func main()  {

	tc, err:= render.CreateTemplateCache()
	if err !=nil {
		log.Print(err)
	}
	app.TemplateCache = tc
	app.UseCache = true


	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.HomePage)
	http.HandleFunc("/about", handlers.Repo.AboutPage)

	log.Fatal(http.ListenAndServe(Port, nil))

}

