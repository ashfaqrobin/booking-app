package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ashfaqrobin/booking-app/internal/config"
	"github.com/ashfaqrobin/booking-app/internal/handlers"
	"github.com/ashfaqrobin/booking-app/internal/models"
	"github.com/ashfaqrobin/booking-app/internal/render"
	"github.com/ashfaqrobin/booking-app/internal/routes"
)

func main() {
	// Register what I want to put in the session
	gob.Register(models.Reservation{})

	// Creating app config
	var app config.AppConfig

	tmplCache, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tmplCache
	app.UseCache = false
	app.InProduction = false

	// Adding session
	session := scs.New()
	session.Lifetime = 24 * time.Hour

	session.Cookie.Secure = app.InProduction
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	// Setting template
	repo := handlers.NewRepo()
	handlers.NewHandler(repo)

	// Setting app config
	config.SetConfig(&app)

	// Creating Server
	fmt.Println("Server is listening on 8080")

	serve := http.Server{
		Addr:    ":8080",
		Handler: routes.Routes(&app),
	}

	err = serve.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
