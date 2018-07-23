// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"log"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"

	"github.com/candig/go-model-service/variant-service/api/restapi/operations"
	apimodels "github.com/candig/go-model-service/variant-service/api/models"
	datamodels "github.com/candig/go-model-service/variant-service/data/models"
)

//go:generate swagger generate server --target .. --name variant-service --spec ../swagger.yml

func getVariantByID(id string, tx *pop.Connection) (*datamodels.Variant, error) {
	variant := &datamodels.Variant{}
	err := tx.Find(variant, id)
	return variant, err
}

//TODO export to transformations package
func transformVariantToAPIModel(dataVariant datamodels.Variant) (*apimodels.Variant, *apimodels.Error) {
	startNonNullable, ok := dataVariant.Start.Interface().(int) // TODO assert as int64?
	if !ok {
		log.Println(
				"500: Transformation of non-nullable field Variant.Start from data to api model fails to yield valid int\n" +
				"In: transformVariantToAPIModel\n")
			errPayload := &apimodels.Error{Code: 50001, Message: swag.String("")} //TODO message
			return nil, errPayload
	}
	transformedStart := int64(startNonNullable)

	apiVariant := &apimodels.Variant{
		ID:			strfmt.UUID(dataVariant.ID.String()),
		Name:		&dataVariant.Name,
		Chromosome:	&dataVariant.Chromosome,
		Start:		&transformedStart,
		Ref:		&dataVariant.Ref,
		Alt:		&dataVariant.Alt}

	// TODO should this validation step be exported to transformations package as well?
	err := apiVariant.Validate(strfmt.NewFormats())
	if err != nil {
			// TODO generalize/modularize error logging in a method
			log.Println(
				"500: API Schema validation for API-model Variant failed upon transformation\n" +
				"In: transformVariantToAPIModel\n" +
				"Error message follows:")
			log.Println(err)
			errPayload := &apimodels.Error{Code: 50001, Message: swag.String("")} //TODO message
			return apiVariant, errPayload
	}

	return apiVariant, nil
}

//TODO export to transformations package
func transformVariantToDataModel(apiVariant apimodels.Variant) (*datamodels.Variant, *apimodels.Error) {
	dataVariant := &datamodels.Variant{
		Name:		*apiVariant.Name,
		Chromosome:	*apiVariant.Chromosome,
		Start:		nulls.NewInt(int(*apiVariant.Start)),
		Ref:		*apiVariant.Ref,
		Alt:		*apiVariant.Alt}

	return dataVariant, nil
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
			log.Panic(
				"500 ERROR: Failed to connect to database: development\n" +
				"In: api.MainGetVariantHandler\n" +
				"Error message follows:")
			log.Println(err)
			errPayload := &apimodels.Error{Code: 50001, Message: swag.String("")} //TODO message
	        return operations.NewMainGetVariantsInternalServerError().WithPayload(errPayload)
	    }

	    query := tx.Where("chromosome = '%s' AND start BETWEEN %d AND %d",
	    	*params.Chromosome, *params.Start, *params.End)
	    dataVariants := []datamodels.Variant{}
	    err = query.All(&dataVariants)
	    if err != nil {
			log.Println( // TODO does this need to be panic?
				"500 ERROR: Problems getting variants from database\n" +
				"In: api.MainGetVariantHandler\n" +
				"Error message follows:")
			log.Println(err)
			errPayload := &apimodels.Error{Code: 50001, Message: swag.String("")} //TODO message
	        return operations.NewMainGetVariantsInternalServerError().WithPayload(errPayload)
	    }

	    // TODO variants of datamodel, not model structure
	    // Iterate through all and convert via a conversion method
	    apiVariants := []*apimodels.Variant{}
	    for _, dataVariant := range dataVariants {
	    	apiVariant, errPayload := transformVariantToAPIModel(dataVariant)
	    	if errPayload != nil {
				return operations.NewMainGetVariantsInternalServerError().WithPayload(errPayload)
	    	}
	    	apiVariants = append(apiVariants, apiVariant)
	    }

	    return operations.NewMainGetVariantsOK().WithPayload(apiVariants)
	})
	api.MainPostVariantHandler = operations.MainPostVariantHandlerFunc(func(params operations.MainPostVariantParams) middleware.Responder {
		err := params.Variant.Validate(strfmt.NewFormats())
		if err != nil {
			log.Println(
				"400: API Schema validation for Variant param failed\n" +
				"In: api.MainPostVariantHandler\n" +
				"Error message follows:")
			log.Println(err)
			errPayload := &apimodels.Error{Code: 40001, Message: swag.String("")} //TODO message
			return operations.NewMainPostVariantBadRequest().WithPayload(errPayload)
		}

		tx, err := pop.Connect("development")
		if err != nil {
			log.Panic(
				"500 ERROR: Failed to connect to database: development\n" +
				"In: api.MainPostVariantHandler\n" +
				"Error message follows:")
			log.Println(err)
			errPayload := &apimodels.Error{Code: 50001, Message: swag.String("")} //TODO message
	        return operations.NewMainPostVariantInternalServerError().WithPayload(errPayload)
	    }

	    _, err = getVariantByID(params.Variant.ID.String(), tx)
	    if err == nil { // TODO not actually a great check
	    	log.Println(
	    		"405: Variant ID already exists in database; cannot overwrite with put\n" +
	    		"In: api.MainPostVariantHandler, getVariantByID(string)")
	    	errPayload := &apimodels.Error{Code: 40501, Message: swag.String("")} //TODO message
	        return operations.NewMainPostVariantMethodNotAllowed().WithPayload(errPayload)
	    }

		newVariant, errPayload := transformVariantToDataModel(*params.Variant)
		if err != nil {
	        return operations.NewMainPostVariantInternalServerError().WithPayload(errPayload)
	    }

		_, err = tx.ValidateAndCreate(newVariant)
		if err != nil {
			log.Println(
				"500 ERROR: ValidateAndCreate into database failed\n" +
				"In: api.MainPostVariantHandler\n" +
				"Error message follows:")
			log.Println(err)
			errPayload := &apimodels.Error{Code: 50001, Message: swag.String("")} //TODO message
	        return operations.NewMainPostVariantInternalServerError().WithPayload(errPayload)
	    }

	    dataVariant, err := getVariantByID(newVariant.ID.String(), tx) //TODO ID
	    if err != nil {
			log.Println(
				"500 ERROR: Failed to get variant by ID from database immediately following its creation\n" +
				"In: api.MainPostVariantHandler, getVariantByID(string)\n" +
				"Error message follows:")
			log.Println(err)
			errPayload := &apimodels.Error{Code: 50001, Message: swag.String("")} //TODO message
	        return operations.NewMainPostVariantInternalServerError().WithPayload(errPayload)
	    }

	    apiVariant, errPayload := transformVariantToAPIModel(*dataVariant)
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
// TODO I changed the Server from graceful.Server to http.Server to avoid
// mismatches with server.go. Must figure out why the auto-generated code was
// incompatible and fix.
func configureServer(s *http.Server, scheme, addr string) {
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
