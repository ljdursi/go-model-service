package handlers

import (
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/CanDIG/go-model-service/variant-service/transformations"
	"github.com/gobuffalo/pop"
		"github.com/CanDIG/go-model-service/variant-service/errors"
	)

// PostVariant processes a variant posted by the API request and creates it into the database.
// It then retrieves the newly created variant from the database and returns it, along with its URL location.
func PostVariant(params operations.PostVariantParams) middleware.Responder {
	funcName := "handlers.PostVariant"

	tx, err := pop.Connect("development")
	if err != nil {
		errors.Log(err, 500, funcName, "Failed to connect to database: development")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostVariantInternalServerError().WithPayload(errPayload)
	}

	_, err = getVariantByID(params.Variant.ID.String(), tx)
	if err == nil { // TODO this is not a great check
		errPayload := errors.OverwriteForbiddenInPost("Variant", funcName)
		return operations.NewPostVariantMethodNotAllowed().WithPayload(errPayload)
	}

	newVariant, errPayload := transformations.VariantAPIToDataModel(*params.Variant, tx)
	if errPayload != nil {
		return operations.NewPostVariantInternalServerError().WithPayload(errPayload)
	}

	err = tx.Create(newVariant)
	if err != nil {
		errors.Log(err, 500, funcName,
			"Create into database failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostVariantInternalServerError().WithPayload(errPayload)
	}

	dataVariant, err := getVariantByID(newVariant.ID.String(), tx)
	if err != nil {
		errors.Log(err, 500, funcName,
			"Failed to get variant by ID from database immediately following its creation")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostVariantInternalServerError().WithPayload(errPayload)
	}

	apiVariant, errPayload := transformations.VariantDataToAPIModel(*dataVariant)
	if err != nil {
		return operations.NewPostVariantInternalServerError().WithPayload(errPayload)
	}

	location := params.HTTPRequest.URL.Host + params.HTTPRequest.URL.EscapedPath() +
		"/" + apiVariant.ID.String()
	return operations.NewPostVariantCreated().WithPayload(apiVariant).WithLocation(location)
}