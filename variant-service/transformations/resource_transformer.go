package transformations

import (
	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	datamodels "github.com/CanDIG/go-model-service/variant-service/data/models"
	"github.com/gobuffalo/pop"
)

// ResourceDataToAPIModel is a generic placeholder function for a data-model-to-api-model transformer.
// This function is used only as scaffolding for the development of api generic handlers, and should never be called.
func ResourceDataToAPIModel(dataResource datamodels.Resource) (*apimodels.Resource, *apimodels.Error) {
	return &apimodels.Resource{}, nil
}

// ResourceDataToAPIModel is a generic placeholder function for a api-model-to-data-model transformer.
// This function is used only as scaffolding for the development of api generic handlers, and should never be called.
func ResourceAPIToDataModel(apiResource apimodels.Resource, tx *pop.Connection) (*datamodels.Resource, *apimodels.Error) {
	return &datamodels.Resource{}, nil
}