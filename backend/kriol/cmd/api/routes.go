//filename: kriol/backend/kriol/cmd/api/main.go

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	//creating httprouter instance
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/createEntryHandler", app.createEntryHandler)
	router.HandlerFunc(http.MethodGet, "/v1/showEntryHandler", app.showEntryHandler)

	return router
}
