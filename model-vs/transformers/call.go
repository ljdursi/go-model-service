package transformers

import (
		"github.com/go-openapi/strfmt"

	apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
	datamodels "github.com/CanDIG/go-model-service/model-vs/data/models"
	)

// CallDataToAPI contains the model-building step of the api-model-to-data-model transformer.
func CallDataToAPI(dataCall datamodels.Call) (*apimodels.Call, error) {
	return &apimodels.Call{
		ID:         	strfmt.UUID(dataCall.ID.String()),
		Created:		strfmt.DateTime(dataCall.CreatedAt),
		IndividualID:	strfmt.UUID(dataCall.Individual.String()),
		VariantID:		strfmt.UUID(dataCall.Variant.String()),
		Genotype:		&dataCall.Genotype,
		Format:			&dataCall.Format}, nil
}

// CallAPIToData contains the model-building step of the data-model-to-api-model transformer.
func CallAPIToData(apiCall apimodels.Call) (*datamodels.Call, error) {
	dataIndividualID, err := uuidAPIToData(apiCall.IndividualID, "IndividualID")
	if err != nil {
		return nil, err
	}
	dataVariantID, err := uuidAPIToData(apiCall.VariantID, "VariantID")
	if err != nil {
		return nil, err
	}

	return &datamodels.Call{
		Individual:		*dataIndividualID,
		Variant:		*dataVariantID,
		Genotype:		*apiCall.Genotype,
		Format:			*apiCall.Format}, nil
}