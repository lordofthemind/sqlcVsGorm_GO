// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlcgen

import (
	"database/sql"
)

type Author struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	Bio         sql.NullString `json:"bio"`
	Email       string         `json:"email"`
	DateOfBirth sql.NullTime   `json:"date_of_birth"`
}
