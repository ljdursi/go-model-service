/*
Package transformers implements model-to-model transformation functions that enable the movement and validation
of data across different layers of the service stack.
 */
package transformers

import (
	"errors"
	"github.com/go-openapi/strfmt"

	apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
	datamodels "github.com/CanDIG/go-model-service/model-vs/data/models"
	"github.com/gobuffalo/pop/nulls"
)

// VariantDataToAPI contains the model-building step of the api-model-to-data-model transformer.
func VariantDataToAPI(dataVariant datamodels.Variant) (*apimodels.Variant, error) {
	startNonNullable, ok := dataVariant.Start.Interface().(int)
	if !ok {
		err := errors.New("Transformation of non-nullable field Variant.Start from data to api model fails to yield valid int")
		return nil, err
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

// VariantAPIToData contains the model-building step of the data-model-to-api-model transformer.
func VariantAPIToData(apiVariant apimodels.Variant) (*datamodels.Variant, error) {
	return &datamodels.Variant{
		Name:       *apiVariant.Name,
		Chromosome: *apiVariant.Chromosome,
		Start:      nulls.NewInt(int(*apiVariant.Start)),
		Ref:        *apiVariant.Ref,
		Alt:        *apiVariant.Alt}, nil
}