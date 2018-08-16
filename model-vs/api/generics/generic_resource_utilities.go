package generics

import (
	"github.com/gobuffalo/pop"
		datamodels "github.com/CanDIG/go-model-service/model-vs/data/models"
)

// GetIndividualByID returns the Individual in the database corresponding to the given ID (or nil if no match is found)
func GetIndividualByID(id string, tx *pop.Connection) (*datamodels.Individual, error) {
	Individual := &datamodels.Individual{}
	err := tx.Find(Individual, id)
	return Individual, err
}