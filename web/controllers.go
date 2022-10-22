package main

import (
	"html/template"
	"net/http"
)

// Handles a Faucet POST request
func handleFaucetRequest(tpl *template.Template, w http.ResponseWriter, req *http.Request) error {
	w.Write([]byte("hola"))
	return nil
}

// Shows the index page
func showFaucetPage(tpl *template.Template, w http.ResponseWriter, req *http.Request) error {
	return tpl.ExecuteTemplate(w, "index.html", nil)
}
