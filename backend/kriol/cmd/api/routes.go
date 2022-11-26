//filename: kriol/backend/kriol/cmd/api/main.go

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	// httprouter instance and paths for handler fucntions
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/schools", app.requirePermission("schools:read", app.listSchoolsHandler))
	router.HandlerFunc(http.MethodPost, "/v1/entries", app.requirePermission("schools:write", app.createEntryHandler))
	router.HandlerFunc(http.MethodGet, "/v1/entries/:id", app.requirePermission("schools:read", app.showEntryHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/entries/:id", app.requirePermission("schools:write", app.updateSchoolHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/entries/:id", app.requirePermission("schools:write", app.deleteSchoolHandler))
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activationUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	return app.recoverPanic(app.rateLimit(app.authenicate(router)))
}
