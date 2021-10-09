package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF prtoection to all POST request
func NoSurf(next http.Handler) http.Handler {
	// var app config.AppConfig
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
