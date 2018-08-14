package transformations

import (
	"github.com/go-openapi/strfmt"

	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	datamodels "github.com/CanDIG/go-model-service/variant-service/data/models"
	)

// callToAPI contains the model-building step of the api-model-to-data-model transformer.
func callToAPI(dataCall datamodels.Call) (*apimodels.Call, *apimodels.Error) {
	return &apimodels.Call{
		ID:         	strfmt.UUID(dataCall.ID.String()),
		Created:		strfmt.DateTime(dataCall.CreatedAt),
		IndividualID:	strfmt.UUID(dataCall.Individual.String()),
		VariantID:		strfmt.UUID(dataCall.Variant.String()),
		Genotype:		&dataCall.Genotype,
		Format:			&dataCall.Format}, nil
}

// callToData contains the model-building step of the data-model-to-api-model transformer.
func callToData(apiCall apimodels.Call) (*datamodels.Call, *apimodels.Error) {
	dataIndividualID, errPayload := uuidAPIToData(apiCall.IndividualID, "IndividualID")
	if errPayload != nil {
		return nil, errPayload
	}
	dataVariantID, errPayload := uuidAPIToData(apiCall.VariantID, "VariantID")
	if errPayload != nil {
		return nil, errPayload
	}

	return &datamodels.Call{
		Individual:		*dataIndividualID,
		Variant:		*dataVariantID,
		Genotype:		*apiCall.Genotype,
		Format:			*apiCall.Format}, nil
}