//Filename: internal/data/school.go

package data

import (
	"database/sql"
	"time"

	"kriol.michaelgomez.net/internal/validator"
)

type School struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Level     string    `json:"level"`
	Contact   string    `json:"contact"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email,omitempty"`
	Website   string    `json:"website,omitempty"`
	Address   string    `json:"address"`
	Mode      []string  `json:"mode"`
	Version   int32     `json:"version"`
}

func ValidateSchool(v *validator.Validator, school *School) {
	//Use the Check() method to execute our validation checks
	v.Check(school.Name != "", "name", "must be provided")
	v.Check(len(school.Name) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(school.Level != "", "level", "must be provided")
	v.Check(len(school.Level) <= 200, "level", "must not be more than 200 bytes long")

	v.Check(school.Contact != "", "contact", "must be provided")
	v.Check(len(school.Contact) <= 200, "contact", "must not be more than 200 bytes long")

	v.Check(school.Phone != "", "phone", "must be provided")
	v.Check(validator.Matches(school.Phone, validator.PhoneRX), "phone", "must not be a valid phone number")

	v.Check(school.Email != "", "email", "must be provided")
	v.Check(validator.Matches(school.Email, validator.EmailRX), "email", "must not be a valid email address")

	v.Check(school.Website != "", "website", "must be provided")
	v.Check(validator.ValidWebsite(school.Website), "website", "must not be a valid url")

	v.Check(school.Address != "", "address", "must be provided")
	v.Check(len(school.Address) <= 500, "address", "must not be more than 200 bytes long")

	v.Check(school.Mode != nil, "mode", "must be provided")
	v.Check(len(school.Mode) >= 1, "mode", "must contain at least 1 entry")
	v.Check(len(school.Mode) >= 5, "mode", "must contain at most 5 entries")
	v.Check(validator.Unique(school.Mode), "mode", "must not contain duplicate entires")
}

// Define a SchoolModel which wraps a sql.DB connection pool
type SchoolModel struct {
	DB *sql.DB
}

// Insert() allows us to create a new school
func (m SchoolModel) Insert(school *School) error {
	return nil
}

// Get() allows us to retrieve a specific School
func (m SchoolModel) Get(id int64) (*School, error) {
	return nil, nil
}

// Update() allows us to edit/alter a specific school
func (m SchoolModel) Update(school *School) error {
	return nil
}

// Delete() removes a specific school
func (m SchoolModel) Delete(id int64) error {
	return nil
}
