package main

import (
	"html/template"
	"net/http"
	"regexp"
)

var (
	addressValidateRegex = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
)

// Handles a Faucet POST request
func handleFaucetRequest(tpl *template.Template, w http.ResponseWriter, req *http.Request) error {
	if req.Method != http.MethodPost {
		http.Redirect(w, req, "/", http.StatusFound)
		return nil
	}

	addr := req.PostFormValue("address")

	// Check if its a valid address
	if !addressValidateRegex.MatchString(addr) {
		return tpl.ExecuteTemplate(w, "index.html", map[string]interface{}{
			"invalidAddress": true,
		})
	}

	tokensGiveawayMap.add(addr)

	return tpl.ExecuteTemplate(w, "index.html", map[string]interface{}{
		"success": true,
	})
}

// Shows the index page
func showFaucetPage(tpl *template.Template, w http.ResponseWriter, req *http.Request) error {
	return tpl.ExecuteTemplate(w, "index.html", nil)
}
