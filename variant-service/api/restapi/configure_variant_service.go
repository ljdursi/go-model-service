// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/tylerb/graceful"

	"github.com/CanDIG/go-model-service/variant-service/api/restapi/operations"
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/handlers"
)

//go:generate swagger generate server --target .. --name variant-service --spec ../swagger.yml

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

	api.GetCallsHandler = operations.GetCallsHandlerFunc(func(params operations.GetCallsParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetCalls has not yet been implemented")
	})
	api.GetIndividualsHandler = operations.GetIndividualsHandlerFunc(func(params operations.GetIndividualsParams) middleware.Responder {
		return handlers.GetIndividuals(params)
	})
	api.GetIndividualsByVariantHandler = operations.GetIndividualsByVariantHandlerFunc(func(params operations.GetIndividualsByVariantParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetIndividualsByVariant has not yet been implemented")
	})
	api.GetOneCallHandler = operations.GetOneCallHandlerFunc(func(params operations.GetOneCallParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetOneCall has not yet been implemented")
	})
	api.GetOneIndividualHandler = operations.GetOneIndividualHandlerFunc(func(params operations.GetOneIndividualParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetOneIndividual has not yet been implemented")
	})
	api.GetOneVariantHandler = operations.GetOneVariantHandlerFunc(func(params operations.GetOneVariantParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetOneVariant has not yet been implemented")
	})
	api.GetVariantsHandler = operations.GetVariantsHandlerFunc(func(params operations.GetVariantsParams) middleware.Responder {
		return handlers.GetVariants(params)
	})
	api.GetVariantsByIndividualHandler = operations.GetVariantsByIndividualHandlerFunc(func(params operations.GetVariantsByIndividualParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetVariantsByIndividual has not yet been implemented")
	})
	api.PostCallHandler = operations.PostCallHandlerFunc(func(params operations.PostCallParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostCall has not yet been implemented")
	})
	api.PostIndividualHandler = operations.PostIndividualHandlerFunc(func(params operations.PostIndividualParams) middleware.Responder {
		return handlers.PostIndividual(params)
	})
	api.PostVariantHandler = operations.PostVariantHandlerFunc(func(params operations.PostVariantParams) middleware.Responder {
		return handlers.PostVariant(params)
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
