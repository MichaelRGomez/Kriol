//Filename: kriol/backend/kriol/cmd/api/helpers.go

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Define a new type named envelope
type envelope map[string]interface {
}

func (app *application) readIDParam(r *http.Request) (int64, error) {
	//getting request from slice
	params := httprouter.ParamsFromContext(r.Context())

	//geting the id
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	//converting map into a JSON object
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	//Decode the request body into the target destination
	err := json.NewDecoder(r.Body).Decode(dst)

	//Check for a bad request
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		//Switch to check for the errors
		switch {
		//check for syntax
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON(at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		//check for wrong types passed by the client
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field  %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		//Empty body
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		//Pass non-nil pointer error
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		//default
		default:
			return err
		}
	}
	return nil
}
