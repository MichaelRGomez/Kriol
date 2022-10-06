//Filename: kriol/backend/kriol/cmd/api/entries.go

package main

import (
	"errors"
	"fmt"
	"net/http"

	"kriol.michaelgomez.net/internal/data"
	"kriol.michaelgomez.net/internal/validator"
)

func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	//Our target decode desitnation
	var input struct {
		Name    string   `json:"name"`
		Level   string   `json:"level"`
		Contact string   `json:"contact"`
		Phone   string   `json:"phone"`
		Email   string   `json:"email"`
		Website string   `json:"website"`
		Address string   `json:"address"`
		Mode    []string `json:"mode"`
	}

	//Initialize a new json.Decoder instance
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	//Copy the valus from the input struct to a new school struct
	school := &data.School{
		Name:    input.Name,
		Level:   input.Level,
		Contact: input.Contact,
		Phone:   input.Phone,
		Email:   input.Email,
		Website: input.Website,
		Address: input.Address,
		Mode:    input.Mode,
	}

	// Initialize a new Validator Instance
	v := validator.New()

	//Check the map to determine if there were any validation errors
	if data.ValidateSchool(v, school); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	//Create a school
	err = app.models.Schools.Insert(school)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	//Create a location header for the newly created resource/School
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/schools/%d", school.ID))

	//Write the JSON response with 201 - Created status code with the body
	//being the School data and the header being the headers map
	err = app.writeJSON(w, http.StatusCreated, envelope{"school": school}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	//getting request data from param function in helpers.go
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	//Fetch the specitfic school
	school, err := app.models.Schools.Get(id)

	//Handle errors
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	//Wrte the data returned by Get()
	err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateSchoolHandler(w http.ResponseWriter, r *http.Request) {
	//This method does a complete replacement
	//Get the id for the school that needs updating
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	//Fetch the original record from the database
	school, err := app.models.Schools.Get(id)

	//Handle errors
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	//Create an input struct to hold data read in from the client
	var input struct {
		Name    string   `json:"name"`
		Level   string   `json:"level"`
		Contact string   `json:"contact"`
		Phone   string   `json:"phone"`
		Email   string   `json:"email"`
		Website string   `json:"website"`
		Address string   `json:"address"`
		Mode    []string `json:"mode"`
	}

	//Initialize a new json.Decoder instance
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	//Copy / Udate the fields/values in the school variable using the fields in the input struct
	school.Name = input.Name
	school.Level = input.Level
	school.Contact = input.Contact
	school.Phone = input.Phone
	school.Email = input.Email
	school.Website = input.Website
	school.Address = input.Address
	school.Mode = input.Mode

	//Perform validation on the updated School. If validation fails, then we send a 422 - unprocessable enitiy response to the client
	// Initialize a new Validator Instance
	v := validator.New()

	//Check the map to determine if there were any validation errors
	if data.ValidateSchool(v, school); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	//Pass the updated school record to the update() method
	err = app.models.Schools.Update(school)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	//Wrte the data returned by Get()
	err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteSchoolHandler(w http.ResponseWriter, r *http.Request) {
	//Get the id for the school that needs updating
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	//Delete the School from the database. Send a 404 Not found status code to the client if there is no mathcing record
	err = app.models.Schools.Delete(id)

	//Handle errors
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	//Return 200 Status Ok to the client with a success message
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "school successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
