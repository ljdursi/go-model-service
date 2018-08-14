package handlers

import (
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/operations"
	"github.com/gobuffalo/pop"
	"fmt"
	"github.com/CanDIG/go-model-service/variant-service/errors"
	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	"github.com/go-openapi/runtime/middleware"
)

// addAND only adds an AND to the given conditions string if it already has contents.
func addAND(conditions string) string {
	if conditions == "" {
		return ""
	} else {
		return conditions + " AND "
	}
}

// getIndividualsQuery builds an Individuals-specific query out of the given parameters.
// Since there are presently no parameters expected for this request, it simply returns all Individuals.
func getIndividualsQuery(params operations.GetIndividualsParams, tx *pop.Connection) (*pop.Query, middleware.Responder) {
	return pop.Q(tx), nil
}

// getCallsQuery builds an Calls-specific query out of the given parameters.
// Since there are presently no parameters expected for this request, it simply returns all Calls.
func getCallsQuery(params operations.GetCallsParams, tx *pop.Connection) (*pop.Query, middleware.Responder) {
	return pop.Q(tx), nil
}

// getVariantsQuery builds an Individuals-specific query out of the given parameters.
// It rejects get-all requests, as such a request would, in a production service, return a prohibitively
// large amount of data and would likely only be entered in error or in malice.
// May return a 403: Forbidden response.
func getVariantsQuery(params operations.GetVariantsParams, tx *pop.Connection) (*pop.Query, middleware.Responder) {
	funcName := "handlers.getVariantsQuery"

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

	if conditions == "" {
		message := "Forbidden to query for all variants. " +
			"Please provide parameters in the query string for 'chromosome', 'start', and/or 'end'."
		errors.Log(nil, 403, funcName, message)
		errPayload := &apimodels.Error{Code: 403001, Message: &message}
		return nil, operations.NewGetVariantsForbidden().WithPayload(errPayload)
	}

	query := tx.Where(conditions)
	return query, nil
}

/*
func getIndividualsByVariantQuery(params operations.GetIndividualsByVariantParams, tx *pop.Connection) (*pop.Query, middleware.Responder) {
	// funcName := "handlers.getIndividualsByVariantQuery"

	// TODO rm
}
*/