package utilities

import (
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/operations"
	"github.com/gobuffalo/pop"
	"fmt"
	"github.com/CanDIG/go-model-service/model-vs/errors"
	apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
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

// GetIndividuals builds an Individuals-specific query out of the given parameters.
// Since there are presently no parameters expected for this request, it simply returns all Individuals.
func GetIndividuals(params operations.GetIndividualsParams, tx *pop.Connection) (*pop.Query, middleware.Responder) {
	return pop.Q(tx), nil
}

// GetCalls builds an Calls-specific query out of the given parameters.
// Since there are presently no parameters expected for this request, it simply returns all Calls.
func GetCalls(params operations.GetCallsParams, tx *pop.Connection) (*pop.Query, middleware.Responder) {
	return pop.Q(tx), nil
}

// GetVariants builds an Individuals-specific query out of the given parameters.
// It rejects get-all requests, as such a request would, in a production service, return a prohibitively
// large amount of data and would likely only be entered in error or in malice.
// May return a 403: Forbidden response.
func GetVariants(params operations.GetVariantsParams, tx *pop.Connection) (*pop.Query, middleware.Responder) {
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

// GetIndividualsByVariant builds a query for all associated Individuals for a given Variant, given Call junctions.
func GetIndividualsByVariant(params operations.GetIndividualsByVariantParams, tx *pop.Connection) (*pop.Query, middleware.Responder) {
	sql := "SELECT i.* FROM calls c " +
		"JOIN individuals i ON i.id = c.Individual " +
		"WHERE c.Variant=?"
	args := params.VariantID.String()
	return tx.RawQuery(sql, args), nil
}

// GetVariantsByIndividual builds a query for all associated Variants for a given Individual, given Call junctions.
func GetVariantsByIndividual(params operations.GetVariantsByIndividualParams, tx *pop.Connection) (*pop.Query, middleware.Responder) {
	sql := "SELECT v.* FROM calls c " +
		"JOIN variants v ON v.id = c.variant " +
		"WHERE c.individual=?"
	args := params.IndividualID.String()
	return tx.RawQuery(sql, args), nil
}