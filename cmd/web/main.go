package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/vrodnenok/go-ship/pkg/config"
	"github.com/vrodnenok/go-ship/pkg/handlers"
	"github.com/vrodnenok/go-ship/pkg/render"
)

var app config.AppConfig

func main() {
	app.PORT_NUMBER = ":8080"
	app.UseCache = false

	// app.InProduction set to to "true" when go in production
	app.InProduction = false

	app.Session = scs.New()
	app.Session.Lifetime = 24 * time.Hour
	app.Session.Cookie.Persist = true
	app.Session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session.Cookie.Secure = app.InProduction

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache", err)
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	srv := &http.Server{
		Addr:    app.PORT_NUMBER,
		Handler: routes(&app),
	}

	fmt.Printf("Listening at port %s \n", app.PORT_NUMBER)
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println("Failer to start server")
	}
}
