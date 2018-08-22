package generics

import (
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/utilities"
	"github.com/CanDIG/go-model-service/model-vs/errors"
		apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
)

// GetOneIndividual returns the Individual in the database that corresponds to a given UUID.
func GetOneIndividual(params operations.GetOneIndividualParams) middleware.Responder {
	funcName := "handlers.GetOneIndividual"

	tx, errPayload := utilities.ConnectDevelopment(funcName)
	if errPayload != nil {
		return operations.NewGetOneIndividualInternalServerError().WithPayload(errPayload)
	}

	dataIndividual, err := utilities.GetIndividualByID(params.IndividualID.String(), tx)
	if err != nil {
		message := "This Individual cannot be found."
		errors.Log(err, 404, funcName, message)
		errPayload := &apimodels.Error{Code: 404001, Message: &message}
		return operations.NewGetOneIndividualNotFound().WithPayload(errPayload)
	}

	apiIndividual, errPayload := individualDataToAPIModel(*dataIndividual)
	if errPayload != nil {
		return operations.NewGetOneIndividualInternalServerError().WithPayload(errPayload)
	}

	return operations.NewGetOneIndividualOK().WithPayload(apiIndividual)
}