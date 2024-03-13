package main

import (
	"fmt"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the // application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintln(w, "status: available")
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, "environment: %s\n", app.config.env)
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, "version: %s\n", version)
	if err != nil {
		return
	}
}
