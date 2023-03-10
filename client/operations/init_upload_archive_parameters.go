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

	"github.com/kaz-as/zip/models"
)

// NewInitUploadArchiveParams creates a new InitUploadArchiveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitUploadArchiveParams() *InitUploadArchiveParams {
	return &InitUploadArchiveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitUploadArchiveParamsWithTimeout creates a new InitUploadArchiveParams object
// with the ability to set a timeout on a request.
func NewInitUploadArchiveParamsWithTimeout(timeout time.Duration) *InitUploadArchiveParams {
	return &InitUploadArchiveParams{
		timeout: timeout,
	}
}

// NewInitUploadArchiveParamsWithContext creates a new InitUploadArchiveParams object
// with the ability to set a context for a request.
func NewInitUploadArchiveParamsWithContext(ctx context.Context) *InitUploadArchiveParams {
	return &InitUploadArchiveParams{
		Context: ctx,
	}
}

// NewInitUploadArchiveParamsWithHTTPClient creates a new InitUploadArchiveParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitUploadArchiveParamsWithHTTPClient(client *http.Client) *InitUploadArchiveParams {
	return &InitUploadArchiveParams{
		HTTPClient: client,
	}
}

/*
InitUploadArchiveParams contains all the parameters to send to the API endpoint

	for the init upload archive operation.

	Typically these are written to a http.Request.
*/
type InitUploadArchiveParams struct {

	// Archive.
	Archive *models.Archive

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the init upload archive params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitUploadArchiveParams) WithDefaults() *InitUploadArchiveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the init upload archive params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitUploadArchiveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the init upload archive params
func (o *InitUploadArchiveParams) WithTimeout(timeout time.Duration) *InitUploadArchiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the init upload archive params
func (o *InitUploadArchiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the init upload archive params
func (o *InitUploadArchiveParams) WithContext(ctx context.Context) *InitUploadArchiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the init upload archive params
func (o *InitUploadArchiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the init upload archive params
func (o *InitUploadArchiveParams) WithHTTPClient(client *http.Client) *InitUploadArchiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the init upload archive params
func (o *InitUploadArchiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArchive adds the archive to the init upload archive params
func (o *InitUploadArchiveParams) WithArchive(archive *models.Archive) *InitUploadArchiveParams {
	o.SetArchive(archive)
	return o
}

// SetArchive adds the archive to the init upload archive params
func (o *InitUploadArchiveParams) SetArchive(archive *models.Archive) {
	o.Archive = archive
}

// WriteToRequest writes these params to a swagger request
func (o *InitUploadArchiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Archive != nil {
		if err := r.SetBodyParam(o.Archive); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
