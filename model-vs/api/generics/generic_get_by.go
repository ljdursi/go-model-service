package generics

import (
	datamodels "github.com/CanDIG/go-model-service/model-vs/data/models"
	apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/handlers/utilities"
	"github.com/CanDIG/go-model-service/model-vs/errors"
		)

// GetIndividualsByVariant returns all Individuals with a given Variant called.
// Since Individuals and Variants have a many-to-many relationship, Calls are used as the relation/junction between them.
func GetIndividualsByVariant(params operations.GetIndividualsByVariantParams) middleware.Responder {
	funcName := "handlers.GetIndividualsByVariant"

	tx, errPayload := utilities.ConnectDevelopment(funcName)
	if errPayload != nil {
		return operations.NewGetIndividualsByVariantInternalServerError().WithPayload(errPayload)
	}

	_, err := utilities.GetVariantByID(params.VariantID.String(), tx)
	if err != nil {
		message := "The Variant by which you are trying to query by cannot be found."
		errors.Log(err, 404, funcName, message)
		errPayload := &apimodels.Error{Code: 404002, Message: &message}
		return operations.NewGetIndividualsByVariantNotFound().WithPayload(errPayload)
	}

	// the full error response is handled here rather than the payload because a variety of http codes may occur
	query, errResponse := utilities.GetIndividualsByVariant(params, tx)
	if errResponse != nil {
		return errResponse
	}

	var dataIndividuals []datamodels.Individual
	err = query.All(&dataIndividuals)
	if err != nil {
		errors.Log(err, 500, funcName, "Problems getting Individuals from database")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewGetIndividualsByVariantInternalServerError().WithPayload(errPayload)
	}

	var apiIndividuals []*apimodels.Individual
	for _, dataIndividual := range dataIndividuals {
		apiIndividual, errPayload := individualDataToAPIModel(dataIndividual)
		if errPayload != nil {
			return operations.NewGetIndividualsByVariantInternalServerError().WithPayload(errPayload)
		}
		apiIndividuals = append(apiIndividuals, apiIndividual)
	}

	return operations.NewGetIndividualsByVariantOK().WithPayload(apiIndividuals)
}