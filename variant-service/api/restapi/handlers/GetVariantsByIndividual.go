package handlers

import (
	datamodels "github.com/CanDIG/go-model-service/variant-service/data/models"
	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/handlers/utilities"
	"github.com/CanDIG/go-model-service/variant-service/errors"
	"github.com/CanDIG/go-model-service/variant-service/transformations"
)

// GetVariantsByIndividual returns all Individuals with a given Variant called.
// Since Individuals and Variants have a many-to-many relationship, Calls are used as the relation/junction between them.
func GetVariantsByIndividual(params operations.GetVariantsByIndividualParams) middleware.Responder {
	funcName := "handlers.GetVariantsByIndividual"

	tx, errPayload := utilities.ConnectDevelopment(funcName)
	if errPayload != nil {
		return operations.NewGetVariantsByIndividualInternalServerError().WithPayload(errPayload)
	}

	_, err := getIndividualByID(params.IndividualID.String(), tx)
	if err != nil {
		message := "The Variant by which you are trying to query by cannot be found."
		errors.Log(err, 404, funcName, message)
		errPayload := &apimodels.Error{Code: 404002, Message: &message}
		return operations.NewGetVariantsByIndividualNotFound().WithPayload(errPayload)
	}

	var dataVariants []datamodels.Variant

	sql := "SELECT v.* FROM calls c " +
		"JOIN variants v ON v.id = c.variant " +
		"WHERE c.individual=?"
	args := params.IndividualID.String()
	err = tx.RawQuery(sql, args).All(&dataVariants)

	if err != nil {
		errors.Log(err, 500, funcName, "Problems getting Variants from database")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewGetVariantsByIndividualInternalServerError().WithPayload(errPayload)
	}

	var apiVariants []*apimodels.Variant
	for _, dataVariant := range dataVariants {
		apiVariant, errPayload := transformations.VariantDataToAPIModel(dataVariant)
		if errPayload != nil {
			return operations.NewGetVariantsByIndividualInternalServerError().WithPayload(errPayload)
		}
		apiVariants = append(apiVariants, apiVariant)
	}

	return operations.NewGetVariantsByIndividualOK().WithPayload(apiVariants)
}
