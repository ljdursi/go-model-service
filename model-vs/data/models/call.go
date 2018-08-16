package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"time"
)

type Call struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	Individual uuid.UUID `json:"individual" db:"individual"`
	Variant    uuid.UUID `json:"variant" db:"variant"`
	Genotype   string    `json:"genotype" db:"genotype"`
	Format     string    `json:"format" db:"format"`
}

// String is not required by pop and may be deleted
func (c Call) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Calls is not required by pop and may be deleted
type Calls []Call

// String is not required by pop and may be deleted
func (c Calls) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Call) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: c.Individual, Name: "Individual"},
		&validators.UUIDIsPresent{Field: c.Variant, Name: "Variant"},
		&validators.StringIsPresent{Field: c.Genotype, Name: "Genotype"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Call) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Call) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
