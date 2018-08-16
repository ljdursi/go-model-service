package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	customValidators "github.com/CanDIG/go-model-service/tools/validators"
	"time"
)

type Variant struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	Name       string    `json:"name" db:"name"`
	Chromosome string    `json:"chromosome" db:"chromosome"`
	Start      nulls.Int `json:"start" db:"start"`
	Ref        string    `json:"ref" db:"ref"`
	Alt        string    `json:"alt" db:"alt"`
}

// String is not required by pop and may be deleted
func (v Variant) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Variants is not required by pop and may be deleted
type Variants []Variant

// String is not required by pop and may be deleted
func (v Variants) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (v *Variant) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: v.Chromosome, Name: "Chromosome"},
		&customValidators.IntIsNotNull{Field: v.Start, Name: "Start"},
		&validators.StringIsPresent{Field: v.Ref, Name: "Ref"},
		&validators.StringIsPresent{Field: v.Alt, Name: "Alt"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (v *Variant) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (v *Variant) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
