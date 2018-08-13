package generics

import (
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/CanDIG/go-model-service/variant-service/transformations"
	"github.com/CanDIG/go-model-service/variant-service/errors"
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/handlers/utilities"
	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
)

// PostResource processes a Resource posted by the API request and creates it into the database.
// It then retrieves the newly created Resource from the database and returns it, along with its URL location.
func PostResource(params operations.PostResourceParams) middleware.Responder {
	funcName := "handlers.Post"

	tx, errPayload := utilities.ConnectDevelopment(funcName)
	if errPayload != nil {
		return operations.NewPostResourceInternalServerError().WithPayload(errPayload)
	}

	_, err := getResourceByID(params.Resource.ID.String(), tx)
	if err == nil { // TODO this is not a great check
		message := "This Resource already exists in the database. " +
			"It cannot be overwritten with POST; please use PUT instead."
		errors.Log(nil, 405, funcName, message)
		errPayload := &apimodels.Error{Code: 405001, Message: &message}
		return operations.NewPostResourceMethodNotAllowed().WithPayload(errPayload)
	}

	newResource, errPayload := transformations.ResourceAPIToDataModel(*params.Resource, tx)
	if errPayload != nil {
		return operations.NewPostResourceInternalServerError().WithPayload(errPayload)
	}

	err = tx.Create(newResource)
	if err != nil {
		errors.Log(err, 500, funcName,
			"Create into database failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostResourceInternalServerError().WithPayload(errPayload)
	}

	// TODO if errors occur from this point on, the resource may have already been created,
	// so it should be deleted prior to return
	retrievedDataResource, err := getResourceByID(newResource.ID.String(), tx)
	if err != nil {
		errors.Log(err, 500, funcName,
			"Failed to get Resource by ID from database immediately following its creation")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostResourceInternalServerError().WithPayload(errPayload)
	}

	retrievedAPIResource, errPayload := transformations.ResourceDataToAPIModel(*retrievedDataResource)
	if err != nil {
		return operations.NewPostResourceInternalServerError().WithPayload(errPayload)
	}

	location := params.HTTPRequest.URL.Host + params.HTTPRequest.URL.EscapedPath() +
		"/" + retrievedAPIResource.ID.String()
	return operations.NewPostResourceCreated().WithPayload(retrievedAPIResource).WithLocation(location)
}