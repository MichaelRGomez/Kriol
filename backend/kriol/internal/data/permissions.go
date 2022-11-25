// kriol/backend/kriol/intenral/data/permissions.go
package data

import (
	"context"
	"database/sql"
	"time"
)

// Define a slice to hold the permissions codes
type Permissions []string

// Checks the slice for a specific permission code
func (p Permissions) Include(code string) bool {
	for i := range p {
		if code == p[i] {
			return true
		}
	}
	return false
}

type PermissionModel struct {
	DB *sql.DB
}

func (m PermissionModel) GetAllForUser(userID int64) (Permissions, error) {
	query := `
		select permissions.code 
		from permissions
		inner join users_permissions
		on users_permissions.permission_id = permissions.id
		inner join users
		on users_permissions.id = users.id
		where users.id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := m.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions Permissions
	for rows.Next() {
		var permisison string
		err := rows.Scan(&permisison)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permisison)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return permissions, nil
}
