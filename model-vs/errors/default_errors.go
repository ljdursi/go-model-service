package errors

import (
	"github.com/go-openapi/swag"
	apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
	)

// DefaultInternalServerError returns an error-non-specific payload for a generic 500 server response.
// For the sake of both simplicity and security, there should not be any further detail returned to the
// client in a 500: Internal Server Error response.
func DefaultInternalServerError() *apimodels.Error {
	return &apimodels.Error{Code: 500000, Message: swag.String("An internal server error has occurred")}
}