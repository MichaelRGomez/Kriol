// Filename: kriol/backend/kriol/backend/cmd/api/users.go
package main

import (
	"errors"
	"net/http"
	"time"

	"kriol.michaelgomez.net/internal/data"
	"kriol.michaelgomez.net/internal/validator"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	//Hold data form the request body
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//Parse the request body into the anoymous struct
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	//Copy the data to a new struct
	user := &data.User{
		Name:      input.Name,
		Email:     input.Email,
		Activated: false,
	}

	//Generate a password hash
	err = user.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	//Perform validation
	v := validator.New()
	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	//Insert the data in the database
	err = app.models.Users.Insert(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	//Add permissions for the newly inserted user
	err = app.models.Permissions.AddForUser(user.ID, "schools:read")
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	//Generate a token for the newly-created user
	token, err := app.models.Tokens.New(user.ID, 1*24*time.Hour, data.ScopeActivation)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.background(func() {
		data := map[string]interface{}{
			"activationToken": token.Plaintext,
			"userID":          user.ID,
		}
		//Send the email to the new user
		err = app.mailer.Send(user.Email, "user_welcome.tmpl", data)
		if err != nil {
			app.logger.PrintError(err, nil)
		}
	})

	//Write a 202 accepted status
	err = app.writeJSON(w, http.StatusAccepted, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) activationUserHandler(w http.ResponseWriter, r *http.Request) {
	//Parse the plaintext activation token
	var input struct {
		TokenPlaintext string `json:"token"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	//Perform validation
	v := validator.New()
	if data.ValidateTokenPlaintext(v, input.TokenPlaintext); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	//Get the user details of the provided token or give the
	//Client feedback about an invalid token
	user, err := app.models.Users.GetForToken(data.ScopeActivation, input.TokenPlaintext)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			v.AddError("token", "invalid or expired activation token")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	//Update the user status
	user.Activated = true

	//Save the updated user's record in our database
	err = app.models.Users.Update(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	//Delete the user's token that was used for activation
	err = app.models.Tokens.DeleteAllForUsers(data.ScopeActivation, user.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	//Send a JSON response with the updated details
	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
