//filename: kriol/backend/kriol/cmd/api/main.go

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// httprouter instance and paths for handler fucntions
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/entries", app.createEntryHandler)
	router.HandlerFunc(http.MethodGet, "/v1/entries/:id", app.showEntryHandler)

	return router
}
