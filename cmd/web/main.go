package main

import (
	"log"
	"net/http"
	"time"

	"github.com/aerosasi/bookings/pkg/config"
	handler "github.com/aerosasi/bookings/pkg/handlers"
	render "github.com/aerosasi/bookings/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true when in productionm
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc , err := render.CreateTemplateCache()
	if err != nil{
		log.Printf("error in CreateTemplateCache")
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)
	render.NewTemplates(&app)


	// http.HandleFunc("/", handler.Repo.Home)
	// http.HandleFunc("/about", handler.Repo.About)

	// http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}