// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// IsCompletedHandlerFunc turns a function with the right signature into a is completed handler
type IsCompletedHandlerFunc func(IsCompletedParams) middleware.Responder

// Handle executing the request and returning a response
func (fn IsCompletedHandlerFunc) Handle(params IsCompletedParams) middleware.Responder {
	return fn(params)
}

// IsCompletedHandler interface for that can handle valid is completed params
type IsCompletedHandler interface {
	Handle(IsCompletedParams) middleware.Responder
}

// NewIsCompleted creates a new http.Handler for the is completed operation
func NewIsCompleted(ctx *middleware.Context, handler IsCompletedHandler) *IsCompleted {
	return &IsCompleted{Context: ctx, Handler: handler}
}

/*
	IsCompleted swagger:route HEAD /files/upload isCompleted

check if archive is completed (unarchived)
*/
type IsCompleted struct {
	Context *middleware.Context
	Handler IsCompletedHandler
}

func (o *IsCompleted) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewIsCompletedParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
