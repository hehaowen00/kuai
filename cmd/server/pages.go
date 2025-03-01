package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

func BindPages(mux *http.ServeMux) {
	fp, _ := filepath.Abs("./templates/*.html")
	t := template.Must(template.ParseGlob(fp))
	debug := true

	sp, err := filepath.Abs("./public/")
	if err != nil {
		panic(err)
	}
	fs := http.FileServer(http.Dir(sp))

	mux.HandleFunc("GET /{$}",
		func(w http.ResponseWriter, r *http.Request) {
			reloadTemplates(t, fp, debug)
			err := t.ExecuteTemplate(w, "index.html", map[string]any{})
			if err != nil {
				panic(err)
			}
		})

	mux.HandleFunc("GET /register",
		func(w http.ResponseWriter, r *http.Request) {
			reloadTemplates(t, fp, debug)
			err := t.ExecuteTemplate(w, "register.html", map[string]any{})
			if err != nil {
				panic(err)
			}
		})

	mux.HandleFunc("GET /login",
		func(w http.ResponseWriter, r *http.Request) {
			reloadTemplates(t, fp, debug)

			cookie, err := r.Cookie(COOKIE_ACCESS_TOKEN)
			if err != http.ErrNoCookie && err != nil {
				panic(err)
			}

			if err != http.ErrNoCookie && cookie != nil {
				http.Redirect(w, r, "/collections", http.StatusTemporaryRedirect)
				return
			}

			err = t.ExecuteTemplate(w, "login.html", map[string]any{})
			if err != nil {
				panic(err)
			}
		})

	mux.HandleFunc("GET /logout",
		func(w http.ResponseWriter, r *http.Request) {
			reloadTemplates(t, fp, debug)

			cookie := http.Cookie{
				Name:    COOKIE_ACCESS_TOKEN,
				Value:   "",
				Expires: time.Now().Add(-time.Hour * 24),
			}

			http.SetCookie(w, &cookie)

			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		})

	mux.HandleFunc("GET /search",
		func(w http.ResponseWriter, r *http.Request) {
			reloadTemplates(t, fp, debug)
		})

	mux.HandleFunc("GET /collections",
		func(w http.ResponseWriter, r *http.Request) {
			reloadTemplates(t, fp, debug)

			cookie, err := r.Cookie(COOKIE_ACCESS_TOKEN)
			if err != http.ErrNoCookie && err != nil {
				panic(err)
			}

			if cookie == nil {
				http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
				return
			}

			err = t.ExecuteTemplate(w, "collections.html", map[string]any{})
			if err != nil {
				panic(err)
			}
		})

	mux.HandleFunc("GET /collections/{cid}",
		func(w http.ResponseWriter, r *http.Request) {
			reloadTemplates(t, fp, debug)
			err := t.ExecuteTemplate(w, "collections_cid.html", map[string]any{})
			if err != nil {
				panic(err)
			}
		})

	mux.HandleFunc("GET /collections/{cid}/edit",
		func(w http.ResponseWriter, r *http.Request) {
			reloadTemplates(t, fp, debug)
		})

	mux.HandleFunc("GET /collections/{cid}/add",
		func(w http.ResponseWriter, r *http.Request) {
			reloadTemplates(t, fp, debug)
		})

	mux.HandleFunc("GET /collections/{cid}/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			reloadTemplates(t, fp, debug)
		})

	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			reloadTemplates(t, fp, debug)
			w.WriteHeader(http.StatusNotFound)
			t.ExecuteTemplate(w, "not_found.html", nil)
		})
}
