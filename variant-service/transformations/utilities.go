package transformations

import (
	"github.com/go-openapi/strfmt"
	"github.com/gobuffalo/uuid"
	"github.com/CanDIG/go-model-service/variant-service/errors"
	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
)

// uuidAPIToData safely transforms a api-model UUID to a data-model UUID.
// This never needs to be done for entity ID fields (primary keys), such as Resource.ID,
// as the primary keys are generated automatically be the ORM.
// On the other hand, foreign keys can be transformed with this function.
func uuidAPIToData(apiUUID strfmt.UUID, fieldName string) (*uuid.UUID, *apimodels.Error) {
	dataUUID, err := uuid.FromString(apiUUID.String())
	if err != nil {
		errors.Log(err, 500, "transformations.uuidAPIToData",
			"Transformation of " + fieldName + " from api to data model fails to yield valid UUID")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	return &dataUUID, nil
}