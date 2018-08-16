package generics

import (
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
		"github.com/CanDIG/go-model-service/model-vs/errors"
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/handlers/utilities"
	apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
)

// PostIndividual processes a Individual posted by the API request and creates it into the database.
// It then retrieves the newly created Individual from the database and returns it, along with its URL location.
func PostIndividual(params operations.PostIndividualParams) middleware.Responder {
	funcName := "handlers.Post"

	tx, errPayload := utilities.ConnectDevelopment(funcName)
	if errPayload != nil {
		return operations.NewPostIndividualInternalServerError().WithPayload(errPayload)
	}

	_, err := utilities.GetIndividualByID(params.Individual.ID.String(), tx)
	if err == nil { // TODO this is not a great check
		message := "This Individual already exists in the database. " +
			"It cannot be overwritten with POST; please use PUT instead."
		errors.Log(nil, 405, funcName, message)
		errPayload := &apimodels.Error{Code: 405001, Message: &message}
		return operations.NewPostIndividualMethodNotAllowed().WithPayload(errPayload)
	}

	newIndividual, errPayload := individualAPIToDataModel(*params.Individual, tx)
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

	// TODO if errors occur from this point on, the Individual may have already been created,
	// so it should be deleted prior to return
	retrievedDataIndividual, err := utilities.GetIndividualByID(newIndividual.ID.String(), tx)
	if err != nil {
		errors.Log(err, 500, funcName,
			"Failed to get Individual by ID from database immediately following its creation")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostIndividualInternalServerError().WithPayload(errPayload)
	}

	retrievedAPIIndividual, errPayload := individualDataToAPIModel(*retrievedDataIndividual)
	if err != nil {
		return operations.NewPostIndividualInternalServerError().WithPayload(errPayload)
	}

	location := params.HTTPRequest.URL.Host + params.HTTPRequest.URL.EscapedPath() +
		"/" + retrievedAPIIndividual.ID.String()
	return operations.NewPostIndividualCreated().WithPayload(retrievedAPIIndividual).WithLocation(location)
}