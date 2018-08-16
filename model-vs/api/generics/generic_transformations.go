package generics

import (
	"github.com/go-openapi/strfmt"
	
	apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
	datamodels "github.com/CanDIG/go-model-service/model-vs/data/models"
	"github.com/CanDIG/go-model-service/model-vs/errors"
	"github.com/gobuffalo/pop"
	"github.com/CanDIG/go-model-service/model-vs/transformers"
)

// individualDataToAPIModel transforms a data.models representation of the Individual from the pop ORM-like
// to an api.models representation of the Individual from the Go-Swagger-defined API.
// This allows for the movement of Individual data from the database to the server for GET requests.Individual
// An *apimodels.Error pointer is returned alongside the transformed Individual for ease of error
// response, as it can be used as the response payload immediately.
func individualDataToAPIModel(dataIndividual datamodels.Individual) (*apimodels.Individual, *apimodels.Error) {
	funcName := "handlers.individualDataToAPIModel"

	apiIndividual, err := transformers.IndividualDataToAPI(dataIndividual)
	if err != nil {
		errors.Log(err, 500, funcName + ", transformers.IndividualDataToAPI",
			"Failed transformation of Individual from data to api model")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	err = apiIndividual.Validate(strfmt.NewFormats())
	if err != nil {
		errors.Log(err, 500, funcName,
			"API schema validation for API-model Individual failed upon transformation")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	return apiIndividual, nil
}

//TODO is it really ok to have the validation occur here, with only a Save in configure_Individual_service following the Individual
// individualAPIToDataModel transforms an api.models representation of the Individual from the Go-Swagger-
// defined API to a data.models representation of the Individual from the pop ORM-like.
// This allows for the movement of Individual data from the server to the database for POST/PUT/DELETE
// requests.
// The transformed Individual is validated within this function prior to its return.
// An *apimodels.Error pointer is returned alongside the transformed Individual for ease of error
// response, as it can be used as the response payload immediately.
func individualAPIToDataModel(apiIndividual apimodels.Individual, tx *pop.Connection) (*datamodels.Individual, *apimodels.Error) {
	funcName := "handlers.individualAPIToDataModel"

	dataIndividual, err := transformers.IndividualAPIToData(apiIndividual)
	if err != nil {
		errors.Log(err, 500, funcName + ", transformers.IndividualAPIToData",
			"Failed transformation of Individual from api to data model")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	validationErrors, err := dataIndividual.Validate(tx)
	if err != nil {
		errors.Log(err, 500, funcName,
			"Data schema validation for data-model Variant failed upon transformation with the following validation errors:\n" +
			validationErrors.Error())
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	return dataIndividual, nil
}