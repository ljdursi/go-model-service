package transformations

import (
	"github.com/go-openapi/strfmt"

	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	datamodels "github.com/CanDIG/go-model-service/variant-service/data/models"
			)

// individualToAPI contains the model-building step of the api-model-to-data-model transformer.
func individualToAPI(dataIndividual datamodels.Individual) (*apimodels.Individual, *apimodels.Error) {
	return &apimodels.Individual{
		ID:         	strfmt.UUID(dataIndividual.ID.String()),
		Created:		strfmt.DateTime(dataIndividual.CreatedAt),
		Description:	&dataIndividual.Description}, nil
}

// individualToData contains the model-building step of the data-model-to-api-model transformer.
func individualToData(apiIndividual apimodels.Individual) (*datamodels.Individual, *apimodels.Error) {
	return &datamodels.Individual{
		Description: *apiIndividual.Description}, nil
}