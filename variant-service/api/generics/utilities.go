package generics

import (
	datamodels "github.com/CanDIG/go-model-service/variant-service/data/models"
	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/operations"
	"github.com/gobuffalo/pop"
	"github.com/go-openapi/runtime/middleware"
)

// Utilities are used by the generics package to make generic template files compilable, but it is not used in code
// generation by genny.

// resourceToAPI is a generic placeholder function for the model-building step of the api-model-to-data-model transformer.
// This function is used only as scaffolding for the development of api generic handlers, and should never be called.
func resourceToAPI(dataResource datamodels.Resource) (*apimodels.Resource, *apimodels.Error) {
	return &apimodels.Resource{}, nil
}

// resourceToData is a generic placeholder function for the model-building step of the data-model-to-api-model transformer.
// This function is used only as scaffolding for the development of api generic handlers, and should never be called.
func resourceToData(apiResource apimodels.Resource) (*datamodels.Resource, *apimodels.Error) {
	return &datamodels.Resource{}, nil
}

// getResourcesQuery is a generic placeholder function for a query-builder.
// This function is used only as scaffolding for the development of api generic handlers, and should never be called.
func getResourcesQuery(params operations.GetResourcesParams, tx *pop.Connection) (*pop.Query, middleware.Responder) {
	return pop.Q(tx), nil
}