package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/hoannguyen02/golang/bookings/pkg/config"
	"github.com/hoannguyen02/golang/bookings/pkg/handlers"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// Todo: Change it to true for production mode
	app.InProduction = false

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	repo := handlers.GetRepo(&app)
	handlers.SetRepo(repo)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
