package handlers

import (
	"log"
	"net/http"

	"github.com/atotto/clipboard"
	"github.com/fajritsaniy/lpse-screening/controller"
	"github.com/fajritsaniy/lpse-screening/model"
	"github.com/fajritsaniy/lpse-screening/utils"
)

// HomeHandler handles requests to the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Render HTML template
	model.RenderTemplate(w, "index", "templates/index.html", model.Page{Title: "InsightForge"})
}

// HomeHandler handles requests to the home page
func TokenGenerator(w http.ResponseWriter, r *http.Request) {
	// Render HTML template
	model.RenderTemplate(w, "tokenGenerator", "templates/tokenGenerator.html", model.Page{Title: "InsightToken"})
}

// APIHandler handles API calls
func FindProjectAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get parameters from the form
	sessionID := r.FormValue("sessionID")
	searchInput := r.FormValue("searchInput")

	key := "0123456789abcdef0123456789abcdef"
	sessionIDDecrypted, _ := utils.Decrypt(key, sessionID)

	controller.FindProjectParticipant(w, sessionIDDecrypted, searchInput)

}

// APIHandler handles API calls
func GenerateTokenAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get parameters from the form
	key := "0123456789abcdef0123456789abcdef" // 32 bytes for AES-256
	sessionID := r.FormValue("sessionID")

	sessionIDEncrypted, _ := utils.Encrypt(key, sessionID)
	err = clipboard.WriteAll(sessionIDEncrypted)
	if err != nil {
		log.Fatal(err)
	}
	// Send a success response to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Token generated and copied to clipboard successfully"))
}
