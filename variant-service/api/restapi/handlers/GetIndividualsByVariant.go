package handlers

import (
	datamodels "github.com/CanDIG/go-model-service/variant-service/data/models"
	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/handlers/utilities"
	"github.com/CanDIG/go-model-service/variant-service/errors"
	"github.com/CanDIG/go-model-service/variant-service/transformations"
	)

// GetIndividualsByVariant returns all Individuals with a given Variant called.
// Since Individuals and Variants have a many-to-many relationship, Calls are used as the relation/junction between them.
func GetIndividualsByVariant(params operations.GetIndividualsByVariantParams) middleware.Responder {
	funcName := "handlers.GetIndividualsByVariant"

	tx, errPayload := utilities.ConnectDevelopment(funcName)
	if errPayload != nil {
		return operations.NewGetIndividualsByVariantInternalServerError().WithPayload(errPayload)
	}

	_, err := getVariantByID(params.VariantID.String(), tx)
	if err != nil {
		message := "The Variant by which you are trying to query by cannot be found."
		errors.Log(err, 404, funcName, message)
		errPayload := &apimodels.Error{Code: 404002, Message: &message}
		return operations.NewGetIndividualsByVariantNotFound().WithPayload(errPayload)
	}

	var dataIndividuals []datamodels.Individual

	sql := "SELECT i.* FROM calls c " +
		"JOIN individuals i ON i.id = c.individual " +
		"WHERE c.variant=?"
	args := params.VariantID.String()
	err = tx.RawQuery(sql, args).All(&dataIndividuals)

	if err != nil {
		errors.Log(err, 500, funcName, "Problems getting Individuals from database")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewGetIndividualsByVariantInternalServerError().WithPayload(errPayload)
	}

	var apiIndividuals []*apimodels.Individual
	for _, dataIndividual := range dataIndividuals {
		apiIndividual, errPayload := transformations.IndividualDataToAPIModel(dataIndividual)
		if errPayload != nil {
			return operations.NewGetIndividualsByVariantInternalServerError().WithPayload(errPayload)
		}
		apiIndividuals = append(apiIndividuals, apiIndividual)
	}

	return operations.NewGetIndividualsByVariantOK().WithPayload(apiIndividuals)
}
