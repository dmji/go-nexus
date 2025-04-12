// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetStatusReader is a Reader for the GetStatus structure.
type GetStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 503:
		result := NewGetStatusServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/status] GetStatus", response, response.Code())
	}
}

// NewGetStatusOK creates a GetStatusOK with default headers values
func NewGetStatusOK() *GetStatusOK {
	return &GetStatusOK{}
}

/*
GetStatusOK describes a response with status code 200, with default header values.

Available to service requests
*/
type GetStatusOK struct {
}

// IsSuccess returns true when this get status o k response has a 2xx status code
func (o *GetStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get status o k response has a 3xx status code
func (o *GetStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get status o k response has a 4xx status code
func (o *GetStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get status o k response has a 5xx status code
func (o *GetStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get status o k response a status code equal to that given
func (o *GetStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get status o k response
func (o *GetStatusOK) Code() int {
	return 200
}

func (o *GetStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/status][%d] getStatusOK", 200)
}

func (o *GetStatusOK) String() string {
	return fmt.Sprintf("[GET /v1/status][%d] getStatusOK", 200)
}

func (o *GetStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStatusServiceUnavailable creates a GetStatusServiceUnavailable with default headers values
func NewGetStatusServiceUnavailable() *GetStatusServiceUnavailable {
	return &GetStatusServiceUnavailable{}
}

/*
GetStatusServiceUnavailable describes a response with status code 503, with default header values.

Unavailable to service requests
*/
type GetStatusServiceUnavailable struct {
}

// IsSuccess returns true when this get status service unavailable response has a 2xx status code
func (o *GetStatusServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get status service unavailable response has a 3xx status code
func (o *GetStatusServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get status service unavailable response has a 4xx status code
func (o *GetStatusServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get status service unavailable response has a 5xx status code
func (o *GetStatusServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get status service unavailable response a status code equal to that given
func (o *GetStatusServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get status service unavailable response
func (o *GetStatusServiceUnavailable) Code() int {
	return 503
}

func (o *GetStatusServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/status][%d] getStatusServiceUnavailable", 503)
}

func (o *GetStatusServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/status][%d] getStatusServiceUnavailable", 503)
}

func (o *GetStatusServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
