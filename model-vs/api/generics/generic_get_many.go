package generics

import (
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/CanDIG/go-model-service/model-vs/errors"
		"github.com/CanDIG/go-model-service/model-vs/api/restapi/handlers/utilities"
	apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
	datamodels "github.com/CanDIG/go-model-service/model-vs/data/models"
)

// GetIndividuals returns all Individuals in the database given zero or more query parameters.
// The query parameters are handled separately in getIndividualsQuery.
func GetIndividuals(params operations.GetIndividualsParams) middleware.Responder {
	funcName := "handlers.GetIndividuals"

	tx, errPayload := utilities.ConnectDevelopment(funcName)
	if errPayload != nil {
		return operations.NewPostIndividualInternalServerError().WithPayload(errPayload)
	}

	// the full error response is handled here rather than the payload because a variety of http codes may occur
	query, errResponse := utilities.GetIndividuals(params, tx)
	if errResponse != nil {
		return errResponse
	}

	var dataIndividuals []datamodels.Individual
	err := query.All(&dataIndividuals)
	if err != nil {
		errors.Log(err, 500, funcName, "Problems getting Individuals from database")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewGetIndividualsInternalServerError().WithPayload(errPayload)
	}

	var apiIndividuals []*apimodels.Individual
	for _, dataIndividual := range dataIndividuals {
		apiIndividual, errPayload := individualDataToAPIModel(dataIndividual)
		if errPayload != nil {
			return operations.NewGetIndividualsInternalServerError().WithPayload(errPayload)
		}
		apiIndividuals = append(apiIndividuals, apiIndividual)
	}

	return operations.NewGetIndividualsOK().WithPayload(apiIndividuals)
}