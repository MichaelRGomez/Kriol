//Filename: kriol/backend/kriol/cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
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
	//Displaying the id, for now
	fmt.Fprintf(w, "show the details for school %d\n", id)
}
