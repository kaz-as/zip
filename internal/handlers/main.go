package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/kaz-as/zip/internal/middlewares"
	"github.com/kaz-as/zip/pkg/archive"
	"github.com/kaz-as/zip/pkg/logger"
	"github.com/kaz-as/zip/restapi"
	"github.com/kaz-as/zip/restapi/operations"
)

// New creates handler with middlewares running after routing
func New(l logger.Interface, archiveService archive.Archiver, mws []middlewares.Middleware) (http.Handler, error) {
	swaggerDoc, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, fmt.Errorf("swagger loading: %s", err)
	}

	api := operations.NewZipAPI(swaggerDoc)
	api.Logger = l.Info
	api.UseSwaggerUI()

	// todo add exact handlers, using repositories and archive service

	return api.Serve(middleware.Builder(middlewares.Chain(mws))), nil
}
