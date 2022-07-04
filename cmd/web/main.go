package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Random-7/GoRcon/pkg/config"
	"github.com/Random-7/GoRcon/pkg/handlers"
	"github.com/Random-7/GoRcon/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	//Create new repo and pass it back to handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	//Pass the AppConfig to the render so it can update the template cache
	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println("Starting Web...")
	_ = http.ListenAndServe(portNumber, nil)
}
