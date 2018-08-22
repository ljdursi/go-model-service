/*
Package utilities implements general-purpose utility functions for use by the restapi handlers.
 */

package utilities

import (
	"github.com/gobuffalo/pop"
	apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
	"github.com/CanDIG/go-model-service/model-vs/errors"
)

// ConnectDevelopment connects to the development database and returns the connection and/or error message
func ConnectDevelopment(funcName string) (*pop.Connection, *apimodels.Error) {
	tx, err := pop.Connect("development")
	if err != nil {
		errors.Log(err, 500, funcName + ", utilities.ConnectDevelopment",
			"Failed to connect to database: development")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}
	return tx, nil
}