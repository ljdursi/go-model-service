package transformers

import (
	"github.com/go-openapi/strfmt"

	apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
	datamodels "github.com/CanDIG/go-model-service/model-vs/data/models"
	)

// IndividualDataToAPI contains the model-building step of the api-model-to-data-model transformer.
func IndividualDataToAPI(dataIndividual datamodels.Individual) (*apimodels.Individual, error) {
	return &apimodels.Individual{
		ID:         	strfmt.UUID(dataIndividual.ID.String()),
		Created:		strfmt.DateTime(dataIndividual.CreatedAt),
		Description:	&dataIndividual.Description}, nil
}

// IndividualAPIToData contains the model-building step of the data-model-to-api-model transformer.
func IndividualAPIToData(apiIndividual apimodels.Individual) (*datamodels.Individual, error) {
	return &datamodels.Individual{
		Description: *apiIndividual.Description}, nil
}