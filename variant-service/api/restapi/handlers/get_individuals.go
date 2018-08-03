package handlers

import (
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gobuffalo/pop"
	"github.com/CanDIG/go-model-service/variant-service/errors"
	"github.com/CanDIG/go-model-service/variant-service/transformations"
	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	datamodels "github.com/CanDIG/go-model-service/variant-service/data/models"
)

// GetIndividuals returns all individuals in the database.
func GetIndividuals(params operations.GetIndividualsParams) middleware.Responder {
	funcName := "handlers.GetIndividuals"

	tx, err := pop.Connect("development")
	if err != nil {
		errors.Log(err, 500, funcName, "Failed to connect to database: development")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewGetIndividualsInternalServerError().WithPayload(errPayload)
	}

	var dataIndividuals []datamodels.Individual
	err = tx.All(&dataIndividuals)
	if err != nil {
		errors.Log(err, 500, funcName, "Problems getting Individuals from database")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewGetIndividualsInternalServerError().WithPayload(errPayload)
	}

	var apiIndividuals []*apimodels.Individual
	for _, dataIndividual := range dataIndividuals {
		apiIndividual, errPayload := transformations.IndividualDataToAPIModel(dataIndividual)
		if errPayload != nil {
			return operations.NewGetIndividualsInternalServerError().WithPayload(errPayload)
		}
		apiIndividuals = append(apiIndividuals, apiIndividual)
	}

	return operations.NewGetIndividualsOK().WithPayload(apiIndividuals)
}