// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewCheckChunksParams creates a new CheckChunksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCheckChunksParams() *CheckChunksParams {
	return &CheckChunksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCheckChunksParamsWithTimeout creates a new CheckChunksParams object
// with the ability to set a timeout on a request.
func NewCheckChunksParamsWithTimeout(timeout time.Duration) *CheckChunksParams {
	return &CheckChunksParams{
		timeout: timeout,
	}
}

// NewCheckChunksParamsWithContext creates a new CheckChunksParams object
// with the ability to set a context for a request.
func NewCheckChunksParamsWithContext(ctx context.Context) *CheckChunksParams {
	return &CheckChunksParams{
		Context: ctx,
	}
}

// NewCheckChunksParamsWithHTTPClient creates a new CheckChunksParams object
// with the ability to set a custom HTTPClient for a request.
func NewCheckChunksParamsWithHTTPClient(client *http.Client) *CheckChunksParams {
	return &CheckChunksParams{
		HTTPClient: client,
	}
}

/*
CheckChunksParams contains all the parameters to send to the API endpoint

	for the check chunks operation.

	Typically these are written to a http.Request.
*/
type CheckChunksParams struct {

	/* ID.

	   archive uid

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the check chunks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckChunksParams) WithDefaults() *CheckChunksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the check chunks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckChunksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the check chunks params
func (o *CheckChunksParams) WithTimeout(timeout time.Duration) *CheckChunksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check chunks params
func (o *CheckChunksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check chunks params
func (o *CheckChunksParams) WithContext(ctx context.Context) *CheckChunksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check chunks params
func (o *CheckChunksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check chunks params
func (o *CheckChunksParams) WithHTTPClient(client *http.Client) *CheckChunksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check chunks params
func (o *CheckChunksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the check chunks params
func (o *CheckChunksParams) WithID(id int64) *CheckChunksParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the check chunks params
func (o *CheckChunksParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CheckChunksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param id
	qrID := o.ID
	qID := swag.FormatInt64(qrID)
	if qID != "" {

		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
