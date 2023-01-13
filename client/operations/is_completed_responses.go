// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaz-as/zip/models"
)

// IsCompletedReader is a Reader for the IsCompleted structure.
type IsCompletedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsCompletedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIsCompletedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewIsCompletedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIsCompletedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIsCompletedOK creates a IsCompletedOK with default headers values
func NewIsCompletedOK() *IsCompletedOK {
	return &IsCompletedOK{}
}

/*
IsCompletedOK describes a response with status code 200, with default header values.

is archive completed
*/
type IsCompletedOK struct {
	Payload bool
}

// IsSuccess returns true when this is completed o k response has a 2xx status code
func (o *IsCompletedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this is completed o k response has a 3xx status code
func (o *IsCompletedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is completed o k response has a 4xx status code
func (o *IsCompletedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this is completed o k response has a 5xx status code
func (o *IsCompletedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this is completed o k response a status code equal to that given
func (o *IsCompletedOK) IsCode(code int) bool {
	return code == 200
}

func (o *IsCompletedOK) Error() string {
	return fmt.Sprintf("[HEAD /files/upload][%d] isCompletedOK  %+v", 200, o.Payload)
}

func (o *IsCompletedOK) String() string {
	return fmt.Sprintf("[HEAD /files/upload][%d] isCompletedOK  %+v", 200, o.Payload)
}

func (o *IsCompletedOK) GetPayload() bool {
	return o.Payload
}

func (o *IsCompletedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsCompletedNotFound creates a IsCompletedNotFound with default headers values
func NewIsCompletedNotFound() *IsCompletedNotFound {
	return &IsCompletedNotFound{}
}

/*
IsCompletedNotFound describes a response with status code 404, with default header values.

archive not found
*/
type IsCompletedNotFound struct {
}

// IsSuccess returns true when this is completed not found response has a 2xx status code
func (o *IsCompletedNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this is completed not found response has a 3xx status code
func (o *IsCompletedNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is completed not found response has a 4xx status code
func (o *IsCompletedNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this is completed not found response has a 5xx status code
func (o *IsCompletedNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this is completed not found response a status code equal to that given
func (o *IsCompletedNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *IsCompletedNotFound) Error() string {
	return fmt.Sprintf("[HEAD /files/upload][%d] isCompletedNotFound ", 404)
}

func (o *IsCompletedNotFound) String() string {
	return fmt.Sprintf("[HEAD /files/upload][%d] isCompletedNotFound ", 404)
}

func (o *IsCompletedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsCompletedDefault creates a IsCompletedDefault with default headers values
func NewIsCompletedDefault(code int) *IsCompletedDefault {
	return &IsCompletedDefault{
		_statusCode: code,
	}
}

/*
IsCompletedDefault describes a response with status code -1, with default header values.

generic error response
*/
type IsCompletedDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the is completed default response
func (o *IsCompletedDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this is completed default response has a 2xx status code
func (o *IsCompletedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this is completed default response has a 3xx status code
func (o *IsCompletedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this is completed default response has a 4xx status code
func (o *IsCompletedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this is completed default response has a 5xx status code
func (o *IsCompletedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this is completed default response a status code equal to that given
func (o *IsCompletedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IsCompletedDefault) Error() string {
	return fmt.Sprintf("[HEAD /files/upload][%d] isCompleted default  %+v", o._statusCode, o.Payload)
}

func (o *IsCompletedDefault) String() string {
	return fmt.Sprintf("[HEAD /files/upload][%d] isCompleted default  %+v", o._statusCode, o.Payload)
}

func (o *IsCompletedDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *IsCompletedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
