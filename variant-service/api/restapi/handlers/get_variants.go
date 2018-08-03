/*
Package handlers implements the backend handling of requests to the server.
This set of functions has been exported into its own package mostly for the sake of legibility.
 */
package handlers

import (
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"fmt"
	"github.com/CanDIG/go-model-service/variant-service/transformations"
	"github.com/gobuffalo/pop"
	"github.com/CanDIG/go-model-service/variant-service/errors"
	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	datamodels "github.com/CanDIG/go-model-service/variant-service/data/models"
)

// GetVariants returns the set of variants in the database that meets a given set of query parameters.
// It rejects get-all requests, as such a request would, in a production service, return a prohibitively
// large amount of data and would likely only be entered in error or in malice.
func GetVariants(params operations.GetVariantsParams) middleware.Responder {
	funcName := "handlers.GetVariants"

	tx, err := pop.Connect("development")
	if err != nil {
		errors.Log(err, 500, funcName,"Failed to connect to database: development")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewGetVariantsInternalServerError().WithPayload(errPayload)
	}

	conditions := ""

	if params.Chromosome != nil {
		conditions = fmt.Sprintf(addAND(conditions) + "chromosome = '%s'", *params.Chromosome)
	}
	if params.Start != nil {
		conditions = fmt.Sprintf(addAND(conditions) + "start >= '%d'", *params.Start)
	}
	if params.End != nil {
		conditions = fmt.Sprintf(addAND(conditions) + "start <= '%d'", *params.End)
	}

	var dataVariants []datamodels.Variant
	if conditions != "" {
		query := tx.Where(conditions)
		err = query.All(&dataVariants)
	} else {
		message := "Forbidden to query for all variants. " +
			"Please provide parameters in the query string for 'chromosome', 'start', and/or 'end'."
		errors.Log(nil, 403, funcName, message)
		errPayload := &apimodels.Error{Code: 403001, Message: &message}
		return operations.NewGetVariantsForbidden().WithPayload(errPayload)
	}

	if err != nil {
		errors.Log(err, 500, funcName, "Problems getting variants from database")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewGetVariantsInternalServerError().WithPayload(errPayload)
	}

	var apiVariants []*apimodels.Variant
	for _, dataVariant := range dataVariants {
		apiVariant, errPayload := transformations.VariantDataToAPIModel(dataVariant)
		if errPayload != nil {
			return operations.NewGetVariantsInternalServerError().WithPayload(errPayload)
		}
		apiVariants = append(apiVariants, apiVariant)
	}

	return operations.NewGetVariantsOK().WithPayload(apiVariants)
}