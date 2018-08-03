// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/tylerb/graceful"
	"github.com/gobuffalo/pop"

	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	datamodels "github.com/CanDIG/go-model-service/variant-service/data/models"
	customErrors "github.com/CanDIG/go-model-service/variant-service/errors"
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/operations"
	"github.com/CanDIG/go-model-service/variant-service/transformations"
)

//go:generate swagger generate server --target .. --name variant-service --spec ../swagger.yml

func getVariantByID(id string, tx *pop.Connection) (*datamodels.Variant, error) {
	variant := &datamodels.Variant{}
	err := tx.Find(variant, id)
	return variant, err
}

// Only add an AND to the conditions string if it already has contents.
 func addAND(conditions string) string {
 	if conditions == "" {
 		return ""
	} else {
		return conditions + " AND "
	}
 }

func configureFlags(api *operations.VariantServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.VariantServiceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.MainGetVariantsHandler = operations.MainGetVariantsHandlerFunc(func(params operations.MainGetVariantsParams) middleware.Responder {
		tx, err := pop.Connect("development")
		if err != nil {
			customErrors.Log(err, 500,"restapi.api.MainGetVariantHandler",
				"Failed to connect to database: development")
			errPayload := customErrors.DefaultInternalServerError()
			return operations.NewMainGetVariantsInternalServerError().WithPayload(errPayload)
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
			customErrors.Log(nil, 403,"api.MainGetVariantsHandler", message)
			errPayload := &apimodels.Error{Code: 403001, Message: &message}
			return operations.NewMainGetVariantsForbidden().WithPayload(errPayload)
		}

		if err != nil {
			// TODO does this need to be panic?
			customErrors.Log(err, 500,"restapi.api.MainGetVariantHandler",
				"Problems getting variants from database")
			errPayload := customErrors.DefaultInternalServerError()
			return operations.NewMainGetVariantsInternalServerError().WithPayload(errPayload)
		}

		var apiVariants []*apimodels.Variant
		for _, dataVariant := range dataVariants {
			apiVariant, errPayload := transformations.VariantDataToAPIModel(dataVariant)
			if errPayload != nil {
				return operations.NewMainGetVariantsInternalServerError().WithPayload(errPayload)
			}
			apiVariants = append(apiVariants, apiVariant)
		}

		return operations.NewMainGetVariantsOK().WithPayload(apiVariants)
	})
	api.MainPostVariantHandler = operations.MainPostVariantHandlerFunc(func(params operations.MainPostVariantParams) middleware.Responder {
		err := params.Variant.Validate(strfmt.NewFormats())

		tx, err := pop.Connect("development")
		if err != nil {
			customErrors.Log(err, 500,"restapi.api.MainPostVariantHandler",
				"Failed to connect to database: development")
			errPayload := customErrors.DefaultInternalServerError()
			return operations.NewMainPostVariantInternalServerError().WithPayload(errPayload)
		}

		_, err = getVariantByID(params.Variant.ID.String(), tx)
		if err == nil { // TODO this is not a great check
			message := "This variant already exists in the database. " +
				"It cannot be overwritten with POST; please use PUT instead."
			customErrors.Log(nil, 405,"restapi.api.MainPostVariantHandler", message)
			errPayload := &apimodels.Error{Code: 405001, Message: &message}
			return operations.NewMainPostVariantMethodNotAllowed().WithPayload(errPayload)
		}

		newVariant, errPayload := transformations.VariantAPIToDataModel(*params.Variant, tx)
		if errPayload != nil {
			return operations.NewMainPostVariantInternalServerError().WithPayload(errPayload)
		}

		_, err = tx.ValidateAndCreate(newVariant)
		if err != nil {
			customErrors.Log(err, 500,"restapi.api.MainPostVariantHandler",
				"ValidateAndCreate into database failed")
			errPayload := customErrors.DefaultInternalServerError()
			return operations.NewMainPostVariantInternalServerError().WithPayload(errPayload)
		}

		dataVariant, err := getVariantByID(newVariant.ID.String(), tx)
		if err != nil {
			customErrors.Log(err, 500,"restapi.api.MainPostVariantHandler, restapi.getVariantByID(string)",
				"Failed to get variant by ID from database immediately following its creation")
			errPayload := customErrors.DefaultInternalServerError()
			return operations.NewMainPostVariantInternalServerError().WithPayload(errPayload)
		}

		apiVariant, errPayload := transformations.VariantDataToAPIModel(*dataVariant)
		if err != nil {
			return operations.NewMainPostVariantInternalServerError().WithPayload(errPayload)
		}

		// TODO check and fix the construction of this URL
		location := params.HTTPRequest.URL.Host + params.HTTPRequest.URL.EscapedPath() +
			"/" + apiVariant.ID.String()
		return operations.NewMainPostVariantCreated().WithPayload(apiVariant).WithLocation(location)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
