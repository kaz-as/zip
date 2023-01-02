// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/kaz-as/zip/restapi/operations"
)

//go:generate swagger generate server --target .. --name Zip --spec ..\docs\swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.ZipAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ZipAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.BinConsumer = runtime.ByteStreamConsumer()
	api.JSONConsumer = runtime.JSONConsumer()

	api.BinProducer = runtime.ByteStreamProducer()
	api.JSONProducer = runtime.JSONProducer()

	if api.GetFilesHandler == nil {
		api.GetFilesHandler = operations.GetFilesHandlerFunc(func(params operations.GetFilesParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetFiles has not yet been implemented")
		})
	}
	if api.GetFilesUploadHandler == nil {
		api.GetFilesUploadHandler = operations.GetFilesUploadHandlerFunc(func(params operations.GetFilesUploadParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetFilesUpload has not yet been implemented")
		})
	}
	if api.PostFilesHandler == nil {
		api.PostFilesHandler = operations.PostFilesHandlerFunc(func(params operations.PostFilesParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostFiles has not yet been implemented")
		})
	}
	if api.PostFilesUploadHandler == nil {
		api.PostFilesUploadHandler = operations.PostFilesUploadHandlerFunc(func(params operations.PostFilesUploadParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostFilesUpload has not yet been implemented")
		})
	}
	if api.PostFilesZipHandler == nil {
		api.PostFilesZipHandler = operations.PostFilesZipHandlerFunc(func(params operations.PostFilesZipParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostFilesZip has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
