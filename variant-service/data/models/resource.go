package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"time"
)

// This is a generic representation of a pop model.
// It is used only as scaffolding for the development of generic api handlers, and should never be called.
type Resource struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
// This function is used only as scaffolding for the development of api generic handlers, and should never be called.
func (i Resource) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// This is used only as scaffolding for the development of api generic handlers, and should never be called.
type Resources []Resource

// String is not required by pop and may be deleted
// This function is used only as scaffolding for the development of api generic handlers, and should never be called.
func (i Resources) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
// This function is used only as scaffolding for the development of api generic handlers, and should never be called.
func (i *Resource) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
// This function is used only as scaffolding for the development of api generic handlers, and should never be called.
func (i *Resource) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
// This function is used only as scaffolding for the development of api generic handlers, and should never be called.
func (i *Resource) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
