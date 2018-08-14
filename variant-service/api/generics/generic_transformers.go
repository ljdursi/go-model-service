package generics

import (
	"github.com/go-openapi/strfmt"
	
	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	datamodels "github.com/CanDIG/go-model-service/variant-service/data/models"
	"github.com/CanDIG/go-model-service/variant-service/errors"
	"github.com/gobuffalo/pop"
)

// ResourceDataToAPIModel transforms a data.models representation of the Resource from the pop ORM-like
// to an api.models representation of the Resource from the Go-Swagger-defined API.
// This allows for the movement of Resource data from the database to the server for GET requests.Resource
// An *apimodels.Error pointer is returned alongside the transformed Resource for ease of error
// response, as it can be used as the response payload immediately.
func ResourceDataToAPIModel(dataResource datamodels.Resource) (*apimodels.Resource, *apimodels.Error) {
	funcName := "transformations.ResourceDataToAPIModel"

	apiResource, errPayload := resourceToAPI(dataResource)
	if errPayload != nil {
		return nil, errPayload
	}

	err := apiResource.Validate(strfmt.NewFormats())
	if err != nil {
		errors.Log(err, 500, funcName,
			"API schema validation for API-model Resource failed upon transformation")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	return apiResource, nil
}


//TODO is it really ok to have the validation occur here, with only a Save in configure_Resource_service following the Resource
// ResourceAPIToDataModel transforms an api.models representation of the Resource from the Go-Swagger-
// defined API to a data.models representation of the Resource from the pop ORM-like.
// This allows for the movement of Resource data from the server to the database for POST/PUT/DELETE
// requests.
// The transformed Resource is validated within this function prior to its return.
// An *apimodels.Error pointer is returned alongside the transformed Resource for ease of error
// response, as it can be used as the response payload immediately.
func ResourceAPIToDataModel(apiResource apimodels.Resource, tx *pop.Connection) (*datamodels.Resource, *apimodels.Error) {
	funcName := "transformations.ResourceAPIToDataModel"

	dataResource, errPayload := resourceToData(apiResource)
	if errPayload != nil {
		return nil, errPayload
	}

	validationErrors, err := dataResource.Validate(tx)
	if err != nil {
		errors.Log(err, 500, funcName,
			"Data schema validation for data-model Resource failed upon transformation with the following validation errors:\n" +
				validationErrors.Error()) // Print validation error messages into logged message string
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	return dataResource, nil
}