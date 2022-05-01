package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// SessionLoad loads and saves session on every request
func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w)
		fmt.Println(r)
		next.ServeHTTP(w, r)
	})
}
