//Filename: internal/data/school.so

package data

import (
	"time"
)

type School struct {
	ID        int64
	CreatedAt time.Time
	Name      string
	Level     string
	Contact   string
	Phone     string
	Email     string
	Website   string
	Address   string
	Mode      []string
	Version   int32
}
