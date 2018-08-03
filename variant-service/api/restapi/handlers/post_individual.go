package handlers

import (
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/CanDIG/go-model-service/variant-service/transformations"
	"github.com/gobuffalo/pop"
		"github.com/CanDIG/go-model-service/variant-service/errors"
)

// TODO clearly these POST functions can be standardized into a single function or factory
// TODO on that note, GetOne functions can also be standardized.

// PostIndividual processes a Individual posted by the API request and creates it into the database.
// It then retrieves the newly created Individual from the database and returns it, along with its URL location.
func PostIndividual(params operations.PostIndividualParams) middleware.Responder {
	funcName := "handlers.PostIndividual"

	tx, err := pop.Connect("development")
	if err != nil {
		errors.Log(err, 500, funcName, "Failed to connect to database: development")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostIndividualInternalServerError().WithPayload(errPayload)
	}

	_, err = getIndividualByID(params.Individual.ID.String(), tx)
	if err == nil { // TODO this is not a great check
		errPayload := errors.OverwriteForbiddenInPost("Individual", funcName)
		return operations.NewPostIndividualMethodNotAllowed().WithPayload(errPayload)
	}

	newIndividual, errPayload := transformations.IndividualAPIToDataModel(*params.Individual, tx)
	if errPayload != nil {
		return operations.NewPostIndividualInternalServerError().WithPayload(errPayload)
	}

	err = tx.Create(newIndividual)
	if err != nil {
		errors.Log(err, 500, funcName,
			"Create into database failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostIndividualInternalServerError().WithPayload(errPayload)
	}

	dataIndividual, err := getIndividualByID(newIndividual.ID.String(), tx)
	if err != nil {
		errors.Log(err, 500, funcName,
			"Failed to get Individual by ID from database immediately following its creation")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostIndividualInternalServerError().WithPayload(errPayload)
	}

	apiIndividual, errPayload := transformations.IndividualDataToAPIModel(*dataIndividual)
	if err != nil {
		return operations.NewPostIndividualInternalServerError().WithPayload(errPayload)
	}

	location := params.HTTPRequest.URL.Host + params.HTTPRequest.URL.EscapedPath() +
		"/" + apiIndividual.ID.String()
	return operations.NewPostIndividualCreated().WithPayload(apiIndividual).WithLocation(location)
}