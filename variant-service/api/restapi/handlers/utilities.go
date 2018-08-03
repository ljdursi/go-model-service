package handlers

import (
	"github.com/gobuffalo/pop"
	datamodels "github.com/CanDIG/go-model-service/variant-service/data/models"
)

// getIndividualByID returns the individual in the database corresponding to the given ID (or nil if no match is found)
func getIndividualByID(id string, tx *pop.Connection) (*datamodels.Individual, error) {
	individual := &datamodels.Individual{}
	err := tx.Find(individual, id)
	return individual, err
}

// getVariantByID returns the variant in the database corresponding to the given ID (or nil if no match is found)
func getVariantByID(id string, tx *pop.Connection) (*datamodels.Variant, error) {
	variant := &datamodels.Variant{}
	err := tx.Find(variant, id)
	return variant, err
}

// addAND only adds an AND to the given conditions string if it already has contents.
func addAND(conditions string) string {
	if conditions == "" {
		return ""
	} else {
		return conditions + " AND "
	}
}