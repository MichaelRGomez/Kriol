//Filename: kriol/backend/kriol/cmd/api/tokens.go

package main

import (
	"errors"
	"net/http"
	"time"

	"kriol.michaelgomez.net/internal/data"
	"kriol.michaelgomez.net/internal/validator"
)

func (app *application) createAuthenticationTokenHandler(w http.ResponseWriter, r *http.Request) {
	//Parse the email and the password from the request body
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	//Validate the email and password
	v := validator.New()
	data.ValidateEmail(v, input.Email)
	data.ValidatePasswordPlaintex(v, input.Password)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	//Get the user details based on the provided email
	user, err := app.models.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.invalidCredntialsResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	//Check if the password matches
	match, err := user.Password.Matches(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	//if passwords don't match, then return an invalid credentials response
	if !match {
		app.invalidCredntialsResponse(w, r)
		return
	}

	//Password is correct, so we will generate a authentication token
	token, err := app.models.Tokens.New(user.ID, 24*time.Hour, data.ScopeAuthentication)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	//return the authentifcation toklen to the client
	err = app.writeJSON(w, http.StatusCreated, envelope{"authentication_token": token}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
