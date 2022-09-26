//Filename: kriol/backend/kriol/cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
	"time"

	"kriol.michaelgomez.net/internal/data"
)

func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new entry..")
}

func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	//getting request data from param function in helpers.go
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Create a new instance of School struct containing the ID we extracted from out URL and some sample data
	school := data.School{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "Apple Tree",
		Level:     "High School",
		Contact:   "Anna Smith",
		Phone:     "601-4111",
		Address:   "14 Apple street",
		Mode:      []string{"blended", "online"},
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, school, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encounterd a problem and could not process your request", http.StatusInternalServerError)
	}
}
