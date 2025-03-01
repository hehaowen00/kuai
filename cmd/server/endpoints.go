package main

import (
	"log"
	"net/http"
	"time"
)

func BindEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/register",
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseMultipartForm(2048)

			first := r.Form.Get("first")
			last := r.Form.Get("last")
			log.Println(first, last)

			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		})

	mux.HandleFunc("POST /api/login",
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseMultipartForm(2048)

			username := r.Form.Get("username")
			password := r.Form.Get("password")
			log.Println(username, password)

			cookie := http.Cookie{
				Name:     COOKIE_ACCESS_TOKEN,
				Value:    username,
				Expires:  time.Now().Add(time.Minute * 15),
				Secure:   false,
				HttpOnly: true,
				Path:     "/",
			}

			http.SetCookie(w, &cookie)

			http.Redirect(w, r, "/collections", http.StatusTemporaryRedirect)
		})
}
