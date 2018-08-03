/*
Package transformations implements model-to-model transformation functions that enable the movement and validation
of data across different layers of the service stack.
 */
package transformations

import (
	"github.com/go-openapi/strfmt"

	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	datamodels "github.com/CanDIG/go-model-service/variant-service/data/models"
	"github.com/CanDIG/go-model-service/variant-service/errors"
	)

// VariantDataToAPIModel transforms a data.models representation of the Variant from the pop ORM-like
// to an api.models representation of the Variant from the Go-Swagger-defined API.
// This allows for the movement of variant data from the database to the server for GET requests.
// An *apimodels.Error pointer is returned alongside the transformed Variant for ease of error
// response, as it can be used as the response payload immediately.
func VariantDataToAPIModel(dataVariant datamodels.Variant) (*apimodels.Variant, *apimodels.Error) {
	startNonNullable, ok := dataVariant.Start.Interface().(int)
	if !ok {
		errors.Log(nil, 500, "transformVariantToAPIModel",
			"Transformation of non-nullable field Variant.Start from data to api model fails to yield valid int")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}
	transformedStart := int64(startNonNullable)

	apiVariant := &apimodels.Variant{
		ID:         strfmt.UUID(dataVariant.ID.String()),
		Name:       &dataVariant.Name,
		Chromosome: &dataVariant.Chromosome,
		Start:      &transformedStart,
		Ref:        &dataVariant.Ref,
		Alt:        &dataVariant.Alt}

	err := apiVariant.Validate(strfmt.NewFormats())
	if err != nil {
		errors.Log(err, 500, "transformations.VariantDataToAPIModel",
			"API schema validation for API-model Variant failed upon transformation")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	return apiVariant, nil
}