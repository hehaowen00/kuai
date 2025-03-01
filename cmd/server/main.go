package main

import (
	"html/template"
	"net/http"
)

func reloadTemplates(t *template.Template, fp string, debug bool) {
	if debug {
		t = template.Must(template.ParseGlob(fp))
	}
}

func main() {
	mux := http.NewServeMux()
	BindPages(mux)
	BindEndpoints(mux)

	http.ListenAndServe(":8080", mux)
}
