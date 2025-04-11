// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SupportzippathReader is a Reader for the Supportzippath structure.
type SupportzippathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SupportzippathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSupportzippathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSupportzippathForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/support/supportzippath] supportzippath", response, response.Code())
	}
}

// NewSupportzippathOK creates a SupportzippathOK with default headers values
func NewSupportzippathOK() *SupportzippathOK {
	return &SupportzippathOK{}
}

/*
SupportzippathOK describes a response with status code 200, with default header values.

successful operation
*/
type SupportzippathOK struct {
}

// IsSuccess returns true when this supportzippath o k response has a 2xx status code
func (o *SupportzippathOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this supportzippath o k response has a 3xx status code
func (o *SupportzippathOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this supportzippath o k response has a 4xx status code
func (o *SupportzippathOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this supportzippath o k response has a 5xx status code
func (o *SupportzippathOK) IsServerError() bool {
	return false
}

// IsCode returns true when this supportzippath o k response a status code equal to that given
func (o *SupportzippathOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the supportzippath o k response
func (o *SupportzippathOK) Code() int {
	return 200
}

func (o *SupportzippathOK) Error() string {
	return fmt.Sprintf("[POST /v1/support/supportzippath][%d] supportzippathOK", 200)
}

func (o *SupportzippathOK) String() string {
	return fmt.Sprintf("[POST /v1/support/supportzippath][%d] supportzippathOK", 200)
}

func (o *SupportzippathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSupportzippathForbidden creates a SupportzippathForbidden with default headers values
func NewSupportzippathForbidden() *SupportzippathForbidden {
	return &SupportzippathForbidden{}
}

/*
SupportzippathForbidden describes a response with status code 403, with default header values.

Insufficient permissions to generate support zip
*/
type SupportzippathForbidden struct {
}

// IsSuccess returns true when this supportzippath forbidden response has a 2xx status code
func (o *SupportzippathForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this supportzippath forbidden response has a 3xx status code
func (o *SupportzippathForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this supportzippath forbidden response has a 4xx status code
func (o *SupportzippathForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this supportzippath forbidden response has a 5xx status code
func (o *SupportzippathForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this supportzippath forbidden response a status code equal to that given
func (o *SupportzippathForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the supportzippath forbidden response
func (o *SupportzippathForbidden) Code() int {
	return 403
}

func (o *SupportzippathForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/support/supportzippath][%d] supportzippathForbidden", 403)
}

func (o *SupportzippathForbidden) String() string {
	return fmt.Sprintf("[POST /v1/support/supportzippath][%d] supportzippathForbidden", 403)
}

func (o *SupportzippathForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
