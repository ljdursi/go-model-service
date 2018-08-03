package errors

import (
	"github.com/go-openapi/swag"
	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	)

// DefaultInternalServerError returns an error-non-specific payload for a generic 500 server response.
// For the sake of both simplicity and security, there should not be any further detail returned to the
// client in a 500: Internal Server Error response.
func DefaultInternalServerError() *apimodels.Error {
	return &apimodels.Error{Code: 500000, Message: swag.String("An internal server error has occurred")}
}

// OverwriteForbiddenInPost responds to the event that a PUT-like entity overwrite was attempted in a
// POST request, by logging the error and constructing the response payload.
func OverwriteForbiddenInPost(entity string, funcName string) *apimodels.Error {
	message := "This " + entity + " already exists in the database. " +
		"It cannot be overwritten with POST; please use PUT instead."
	Log(nil, 405, funcName, message)
	errPayload := &apimodels.Error{Code: 405001, Message: &message}
	return errPayload
}