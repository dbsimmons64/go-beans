package main

import "net/http"

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/about", app.About)
	mux.HandleFunc("/contact", app.Contact)

	// Serve static files
	fileserver := http.FileServer(http.Dir("./assets/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

	return mux
}
