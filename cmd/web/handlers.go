package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		// log.Print(err.Error())
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
	}
	// w.Write([]byte("Hello from main app"))
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	app.infoLog.Print(r.URL)
	fmt.Fprintf(w, "<h1>Display a specific snippet with id = %d...</h1>", id)
	// w.Write([]byte("<h1>Displaying a specific snippet</h1>"))
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Print(r.URL)
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte("METHOD NOT ALLOWED"))
		// This is the short version of the above
		app.clientError(w, http.StatusMethodNotAllowed)
		// http.NotA
		return
	}
	w.WriteHeader(201)
	w.Write([]byte("Creating a snippet"))
}
