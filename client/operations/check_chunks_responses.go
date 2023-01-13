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

// CheckChunksReader is a Reader for the CheckChunks structure.
type CheckChunksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckChunksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckChunksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCheckChunksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCheckChunksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCheckChunksOK creates a CheckChunksOK with default headers values
func NewCheckChunksOK() *CheckChunksOK {
	return &CheckChunksOK{}
}

/*
CheckChunksOK describes a response with status code 200, with default header values.

list of chunks number
*/
type CheckChunksOK struct {
	Payload []models.ChunkNumber
}

// IsSuccess returns true when this check chunks o k response has a 2xx status code
func (o *CheckChunksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this check chunks o k response has a 3xx status code
func (o *CheckChunksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check chunks o k response has a 4xx status code
func (o *CheckChunksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this check chunks o k response has a 5xx status code
func (o *CheckChunksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this check chunks o k response a status code equal to that given
func (o *CheckChunksOK) IsCode(code int) bool {
	return code == 200
}

func (o *CheckChunksOK) Error() string {
	return fmt.Sprintf("[GET /files/upload][%d] checkChunksOK  %+v", 200, o.Payload)
}

func (o *CheckChunksOK) String() string {
	return fmt.Sprintf("[GET /files/upload][%d] checkChunksOK  %+v", 200, o.Payload)
}

func (o *CheckChunksOK) GetPayload() []models.ChunkNumber {
	return o.Payload
}

func (o *CheckChunksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckChunksNotFound creates a CheckChunksNotFound with default headers values
func NewCheckChunksNotFound() *CheckChunksNotFound {
	return &CheckChunksNotFound{}
}

/*
CheckChunksNotFound describes a response with status code 404, with default header values.

archive not found
*/
type CheckChunksNotFound struct {
}

// IsSuccess returns true when this check chunks not found response has a 2xx status code
func (o *CheckChunksNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this check chunks not found response has a 3xx status code
func (o *CheckChunksNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check chunks not found response has a 4xx status code
func (o *CheckChunksNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this check chunks not found response has a 5xx status code
func (o *CheckChunksNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this check chunks not found response a status code equal to that given
func (o *CheckChunksNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CheckChunksNotFound) Error() string {
	return fmt.Sprintf("[GET /files/upload][%d] checkChunksNotFound ", 404)
}

func (o *CheckChunksNotFound) String() string {
	return fmt.Sprintf("[GET /files/upload][%d] checkChunksNotFound ", 404)
}

func (o *CheckChunksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCheckChunksDefault creates a CheckChunksDefault with default headers values
func NewCheckChunksDefault(code int) *CheckChunksDefault {
	return &CheckChunksDefault{
		_statusCode: code,
	}
}

/*
CheckChunksDefault describes a response with status code -1, with default header values.

generic error response
*/
type CheckChunksDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the check chunks default response
func (o *CheckChunksDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this check chunks default response has a 2xx status code
func (o *CheckChunksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this check chunks default response has a 3xx status code
func (o *CheckChunksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this check chunks default response has a 4xx status code
func (o *CheckChunksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this check chunks default response has a 5xx status code
func (o *CheckChunksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this check chunks default response a status code equal to that given
func (o *CheckChunksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CheckChunksDefault) Error() string {
	return fmt.Sprintf("[GET /files/upload][%d] checkChunks default  %+v", o._statusCode, o.Payload)
}

func (o *CheckChunksDefault) String() string {
	return fmt.Sprintf("[GET /files/upload][%d] checkChunks default  %+v", o._statusCode, o.Payload)
}

func (o *CheckChunksDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CheckChunksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}