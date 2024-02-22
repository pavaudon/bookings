package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/pavaudon/bookings/internal/config"
	"github.com/pavaudon/bookings/internal/handlers"
	"github.com/pavaudon/bookings/internal/models"
	"github.com/pavaudon/bookings/internal/render"
)

const port = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	fmt.Printf("printf Starting application on port %s\n", port)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	//what am i going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	return nil
}
