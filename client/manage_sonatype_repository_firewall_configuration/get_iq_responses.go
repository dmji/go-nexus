// Code generated by go-swagger; DO NOT EDIT.

package manage_sonatype_repository_firewall_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetIqReader is a Reader for the GetIq structure.
type GetIqReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIqReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIqOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/iq] GetIq", response, response.Code())
	}
}

// NewGetIqOK creates a GetIqOK with default headers values
func NewGetIqOK() *GetIqOK {
	return &GetIqOK{}
}

/*
GetIqOK describes a response with status code 200, with default header values.

Sonatype Repository Firewall configuration returned
*/
type GetIqOK struct {
}

// IsSuccess returns true when this get iq o k response has a 2xx status code
func (o *GetIqOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get iq o k response has a 3xx status code
func (o *GetIqOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get iq o k response has a 4xx status code
func (o *GetIqOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get iq o k response has a 5xx status code
func (o *GetIqOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get iq o k response a status code equal to that given
func (o *GetIqOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get iq o k response
func (o *GetIqOK) Code() int {
	return 200
}

func (o *GetIqOK) Error() string {
	return fmt.Sprintf("[GET /v1/iq][%d] getIqOK", 200)
}

func (o *GetIqOK) String() string {
	return fmt.Sprintf("[GET /v1/iq][%d] getIqOK", 200)
}

func (o *GetIqOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
