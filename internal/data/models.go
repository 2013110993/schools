// Filename: internal/data/models.go

package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

// A wrapper for our data models
type Models struct {
	Schools SchoolModel
}

// NewModels() allow us to create a new models
func NewModels(db *sql.DB) Models {
	return Models{
		Schools: SchoolModel{DB: db},
	}
}
