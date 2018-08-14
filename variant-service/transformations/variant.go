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
		"github.com/gobuffalo/pop/nulls"
)

// variantToAPI contains the model-building step of the api-model-to-data-model transformer.
func variantToAPI(dataVariant datamodels.Variant) (*apimodels.Variant, *apimodels.Error) {
	startNonNullable, ok := dataVariant.Start.Interface().(int)
	if !ok {
		errors.Log(nil, 500, "transformations.variantToAPI",
			"Transformation of non-nullable field Variant.Start from data to api model fails to yield valid int")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}
	transformedStart := int64(startNonNullable)

	return &apimodels.Variant{
		ID:         strfmt.UUID(dataVariant.ID.String()),
		Name:       &dataVariant.Name,
		Chromosome: &dataVariant.Chromosome,
		Start:      &transformedStart,
		Ref:        &dataVariant.Ref,
		Alt:        &dataVariant.Alt}, nil
}

// variantToData contains the model-building step of the data-model-to-api-model transformer.
func variantToData(apiVariant apimodels.Variant) (*datamodels.Variant, *apimodels.Error) {
	return &datamodels.Variant{
		Name:       *apiVariant.Name,
		Chromosome: *apiVariant.Chromosome,
		Start:      nulls.NewInt(int(*apiVariant.Start)),
		Ref:        *apiVariant.Ref,
		Alt:        *apiVariant.Alt}, nil
}