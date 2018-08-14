package generics

import (
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/CanDIG/go-model-service/variant-service/errors"
	"github.com/CanDIG/go-model-service/variant-service/transformations"
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/handlers/utilities"
	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
)

// GetOneResource returns the Resource in the database that corresponds to a given UUID.
func GetOneResource(params operations.GetOneResourceParams) middleware.Responder {
	funcName := "handlers.GetOneResource"

	tx, errPayload := utilities.ConnectDevelopment(funcName)
	if errPayload != nil {
		return operations.NewGetOneResourceInternalServerError().WithPayload(errPayload)
	}

	dataResource, err := getResourceByID(params.ResourceID.String(), tx)
	if err != nil {
		message := "This Resource cannot be found."
		errors.Log(err, 404, funcName, message)
		errPayload := &apimodels.Error{Code: 404001, Message: &message}
		return operations.NewGetOneResourceNotFound().WithPayload(errPayload)
	}

	apiResource, errPayload := transformations.ResourceDataToAPIModel(*dataResource)
	if errPayload != nil {
		return operations.NewGetOneResourceInternalServerError().WithPayload(errPayload)
	}

	return operations.NewGetOneResourceOK().WithPayload(apiResource)
}