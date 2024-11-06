package controllers

import (
	"net/http"

	"github.com/unrolled/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// Initialize render with the correct layout file name
	renderer := render.New(render.Options{
		Layout:     "layout", // Correct the layout name if necessary
		Extensions: []string{".html", ".tmpl"},
	})

	// Render the HTML response
	_ = renderer.HTML(w, http.StatusOK, "home", map[string]interface{}{
		"Title": "Home Title", // Fixed typo
		"Body":  "Home Body",  // Correct key if needed
	})
}
