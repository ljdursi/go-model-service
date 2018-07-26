package validators

import (
	"fmt"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

//TODO generalize to nulls of any type or REMOVE

// IsNotNull is a structure describing a named nullable
type IsNotNull struct {
	Name  string
	Field nulls.Nulls
}

// IsValid checks if nullible field is null; if so returns error
func (v *IsNotNull) IsValid(errors *validate.Errors) {
	if v.Field.Interface() == nil {
		errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("%s can not be null.", v.Name))
	}
}
