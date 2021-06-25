package main

import (
	"bookings/internal/config"
	"bookings/internal/handlers"
	"bookings/internal/render"
	"fmt"

	"log"
	"time"

	"net/http"

	"github.com/alexedwards/scs"
)

const portNumber = ":8080"

//create variable for config pkg
var app config.AppConfig
var session *scs.SessionManager

func main() {

	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour

	//default session package
	//cookie will persist after browser closed
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	//call func from handlers
	repo := handlers.NewRepo(&app)
	//after crearting repo variable, pass back to handlers
	handlers.NewHandlers(repo)

	//provide access to app config
	render.NewTemplate(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	fmt.Printf("Starting Application on port %s\n", portNumber)
	//start actual server
	err = srv.ListenAndServe()
	log.Fatal(err)

}
