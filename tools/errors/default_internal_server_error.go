package errors

import (
	"github.com/go-openapi/swag"
	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
)

func DefaultInternalServerError() *apimodels.Error {
	return &apimodels.Error{Code: 500000, Message: swag.String("An internal server error has occurred")}
}
