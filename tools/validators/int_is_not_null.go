package validators

import (
	"fmt"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gobuffalo/pop/nulls"
)

//TODO generalize to nulls of any type

type IntIsNotNull struct {
	Name string
	Field nulls.Int
}

func (v *IntIsNotNull) IsValid(errors *validate.Errors) {
	if v.Field.Interface() == nil {
		errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("%s can not be null.", v.Name))
	}
}
