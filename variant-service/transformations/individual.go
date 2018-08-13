package transformations

import (
	"github.com/go-openapi/strfmt"

	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	datamodels "github.com/CanDIG/go-model-service/variant-service/data/models"
	"github.com/CanDIG/go-model-service/variant-service/errors"
	"github.com/gobuffalo/pop"
	)

// IndividualDataToAPIModel transforms a data.models representation of the Individual from the pop ORM-like
// to an api.models representation of the Individual from the Go-Swagger-defined API.
// This allows for the movement of Individual data from the database to the server for GET requests.Individual
// An *apimodels.Error pointer is returned alongside the transformed Individual for ease of error
// response, as it can be used as the response payload immediately.
func IndividualDataToAPIModel(dataIndividual datamodels.Individual) (*apimodels.Individual, *apimodels.Error) {
	apiIndividual := &apimodels.Individual{
		ID:         	strfmt.UUID(dataIndividual.ID.String()),
		Created:		strfmt.DateTime(dataIndividual.CreatedAt),
		Description:	&dataIndividual.Description}

	err := apiIndividual.Validate(strfmt.NewFormats())
	if err != nil {
		errors.Log(err, 500, "transformations.IndividualDataToAPIModel",
			"API schema validation for API-model Individual failed upon transformation")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	return apiIndividual, nil
}


//TODO is it really ok to have the validation occur here, with only a Save in configure_Individual_service following the call
// IndividualAPIToDataModel transforms an api.models representation of the Individual from the Go-Swagger-
// defined API to a data.models representation of the Individual from the pop ORM-like.
// This allows for the movement of Individual data from the server to the database for POST/PUT/DELETE
// requests.
// The transformed Individual is validated within this function prior to its return.
// An *apimodels.Error pointer is returned alongside the transformed Individual for ease of error
func IndividualAPIToDataModel(apiIndividual apimodels.Individual, tx *pop.Connection) (*datamodels.Individual, *apimodels.Error) {
	dataIndividual := &datamodels.Individual{
		Description: *apiIndividual.Description}

	validationErrors, err := dataIndividual.Validate(tx)
	if err != nil {
		errors.Log(err, 500,"transformations.IndividualAPIToDataModel",
			"Data schema validation for data-model Individual failed upon transformation with the following validation errors:\n" +
				validationErrors.Error()) // Print validation error messages into logged message string
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	return dataIndividual, nil
}