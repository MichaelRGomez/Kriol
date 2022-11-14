// Filename: internal/data/models.go
package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// A wrapper for our data models
type Models struct {
	Schools SchoolModel
	Users   UserModel
	Tokens  TokenModel
}

// NewModels() allows us to create a new model
func NewModels(db *sql.DB) Models {
	return Models{
		Schools: SchoolModel{DB: db},
		Users:   UserModel{DB: db},
		Tokens:  TokenModel{DB: db},
	}
}
