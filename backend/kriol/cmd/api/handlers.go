package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//create a map that'll hold healtcheck data
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)

	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
}

func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new entry..")
}

func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	//Displaying the id, for now
	fmt.Fprintf(w, "show the details for school %d\n", id)
}
