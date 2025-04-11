// Code generated by go-swagger; DO NOT EDIT.

package manage_sonatype_repository_firewall_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VerifyConnection1Reader is a Reader for the VerifyConnection1 structure.
type VerifyConnection1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerifyConnection1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVerifyConnection1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/iq/verify-connection] verifyConnection_1", response, response.Code())
	}
}

// NewVerifyConnection1OK creates a VerifyConnection1OK with default headers values
func NewVerifyConnection1OK() *VerifyConnection1OK {
	return &VerifyConnection1OK{}
}

/*
VerifyConnection1OK describes a response with status code 200, with default header values.

Connection verification complete, check response body for result
*/
type VerifyConnection1OK struct {
}

// IsSuccess returns true when this verify connection1 o k response has a 2xx status code
func (o *VerifyConnection1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this verify connection1 o k response has a 3xx status code
func (o *VerifyConnection1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this verify connection1 o k response has a 4xx status code
func (o *VerifyConnection1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this verify connection1 o k response has a 5xx status code
func (o *VerifyConnection1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this verify connection1 o k response a status code equal to that given
func (o *VerifyConnection1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the verify connection1 o k response
func (o *VerifyConnection1OK) Code() int {
	return 200
}

func (o *VerifyConnection1OK) Error() string {
	return fmt.Sprintf("[POST /v1/iq/verify-connection][%d] verifyConnection1OK", 200)
}

func (o *VerifyConnection1OK) String() string {
	return fmt.Sprintf("[POST /v1/iq/verify-connection][%d] verifyConnection1OK", 200)
}

func (o *VerifyConnection1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
