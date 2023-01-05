// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kaz-as/zip/models"
)

// InitUploadArchiveOKCode is the HTTP code returned for type InitUploadArchiveOK
const InitUploadArchiveOKCode int = 200

/*
InitUploadArchiveOK new upload initialized

swagger:response initUploadArchiveOK
*/
type InitUploadArchiveOK struct {

	/*
	  In: Body
	*/
	Payload *models.InitUploadSuccess `json:"body,omitempty"`
}

// NewInitUploadArchiveOK creates InitUploadArchiveOK with default headers values
func NewInitUploadArchiveOK() *InitUploadArchiveOK {

	return &InitUploadArchiveOK{}
}

// WithPayload adds the payload to the init upload archive o k response
func (o *InitUploadArchiveOK) WithPayload(payload *models.InitUploadSuccess) *InitUploadArchiveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the init upload archive o k response
func (o *InitUploadArchiveOK) SetPayload(payload *models.InitUploadSuccess) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InitUploadArchiveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
InitUploadArchiveDefault generic error response

swagger:response initUploadArchiveDefault
*/
type InitUploadArchiveDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewInitUploadArchiveDefault creates InitUploadArchiveDefault with default headers values
func NewInitUploadArchiveDefault(code int) *InitUploadArchiveDefault {
	if code <= 0 {
		code = 500
	}

	return &InitUploadArchiveDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the init upload archive default response
func (o *InitUploadArchiveDefault) WithStatusCode(code int) *InitUploadArchiveDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the init upload archive default response
func (o *InitUploadArchiveDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the init upload archive default response
func (o *InitUploadArchiveDefault) WithPayload(payload *models.Error) *InitUploadArchiveDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the init upload archive default response
func (o *InitUploadArchiveDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InitUploadArchiveDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}