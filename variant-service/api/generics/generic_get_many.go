package generics

import (
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/CanDIG/go-model-service/variant-service/errors"
	"github.com/CanDIG/go-model-service/variant-service/transformations"
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/handlers/utilities"
	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	datamodels "github.com/CanDIG/go-model-service/variant-service/data/models"
)

// GetResources returns all Resources in the database given zero or more query parameters.
// The query parameters are handled separately in getResourcesQuery.
func GetResources(params operations.GetResourcesParams) middleware.Responder {
	funcName := "handlers.GetResources"

	tx, errPayload := utilities.ConnectDevelopment(funcName)
	if errPayload != nil {
		return operations.NewPostResourceInternalServerError().WithPayload(errPayload)
	}

	// the full error response is handled here rather than the payload because a variety of http codes may occur
	query, errResponse := getResourcesQuery(params, tx)
	if errResponse != nil {
		return errResponse
	}

	var dataResources []datamodels.Resource
	err := query.All(&dataResources)
	if err != nil {
		errors.Log(err, 500, funcName, "Problems getting Resources from database")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewGetResourcesInternalServerError().WithPayload(errPayload)
	}

	var apiResources []*apimodels.Resource
	for _, dataResource := range dataResources {
		apiResource, errPayload := transformations.ResourceDataToAPIModel(dataResource)
		if errPayload != nil {
			return operations.NewGetResourcesInternalServerError().WithPayload(errPayload)
		}
		apiResources = append(apiResources, apiResource)
	}

	return operations.NewGetResourcesOK().WithPayload(apiResources)
}