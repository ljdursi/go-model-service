package transformers

import (
	"errors"
	"github.com/go-openapi/strfmt"
	"github.com/gobuffalo/uuid"
	)

// uuidAPIToData safely transforms a api-model UUID to a data-model UUID.
// This never needs to be done for entity ID fields (primary keys), such as Resource.ID,
// as the primary keys are generated automatically be the ORM.
// On the other hand, foreign keys can be transformed with this function.
func uuidAPIToData(apiUUID strfmt.UUID, fieldName string) (*uuid.UUID, error) {
	dataUUID, err := uuid.FromString(apiUUID.String())
	if err != nil {
		message := "Transformation of " + fieldName + " from api to data model fails to yield valid UUID with the following errors:\n"
		return nil, errors.New(message + err.Error())
	}

	return &dataUUID, nil
}