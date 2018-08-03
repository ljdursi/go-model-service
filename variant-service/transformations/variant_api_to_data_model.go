package transformations

import (
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"

	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	datamodels "github.com/CanDIG/go-model-service/variant-service/data/models"
		"github.com/CanDIG/go-model-service/variant-service/errors"
)

//TODO is it really ok to have the validation occur here, with only a Save in configure_variant_service following the call
// VariantAPIToDataModel transforms an api.models representation of the Variant from the Go-Swagger-
// defined API to a data.models representation of the Variant from the pop ORM-like.
// This allows for the movement of variant data from the server to the database for POST/PUT/DELETE
// requests.
// The transformed Variant is validated within this function prior to its return.
// An *apimodels.Error pointer is returned alongside the transformed Variant for ease of error
// response, as it can be used as the response payload immediately.
func VariantAPIToDataModel(apiVariant apimodels.Variant, tx *pop.Connection) (*datamodels.Variant, *apimodels.Error) {
	dataVariant := &datamodels.Variant{
		Name:       *apiVariant.Name,
		Chromosome: *apiVariant.Chromosome,
		Start:      nulls.NewInt(int(*apiVariant.Start)),
		Ref:        *apiVariant.Ref,
		Alt:        *apiVariant.Alt}

	validationErrors, err := dataVariant.Validate(tx)
	if err != nil {
		errors.Log(err, 500,"transformations.VariantAPIToDataModel",
			"Data schema validation for data-model Variant failed upon transformation with the following validation errors:\n" +
			validationErrors.Error()) // Print validation error messages into logged message string
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	return dataVariant, nil
}