package validators

import (
	"fmt"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

//TODO generalize to nulls of any type


// IntIsNotNull is a structure describing a named nullable int
type IntIsNotNull struct {
	Name  string
	Field nulls.Int
}

// IsValid checks if nullible int is null; if so returns error
func (v *IntIsNotNull) IsValid(errors *validate.Errors) {
	if v.Field.Interface() == nil {
		errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("%s can not be null.", v.Name))
	}
}
