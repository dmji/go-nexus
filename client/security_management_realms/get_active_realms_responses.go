// Code generated by go-swagger; DO NOT EDIT.

package security_management_realms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetActiveRealmsReader is a Reader for the GetActiveRealms structure.
type GetActiveRealmsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActiveRealmsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetActiveRealmsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/security/realms/active] getActiveRealms", response, response.Code())
	}
}

// NewGetActiveRealmsOK creates a GetActiveRealmsOK with default headers values
func NewGetActiveRealmsOK() *GetActiveRealmsOK {
	return &GetActiveRealmsOK{}
}

/*
GetActiveRealmsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetActiveRealmsOK struct {
	Payload []string
}

// IsSuccess returns true when this get active realms o k response has a 2xx status code
func (o *GetActiveRealmsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get active realms o k response has a 3xx status code
func (o *GetActiveRealmsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get active realms o k response has a 4xx status code
func (o *GetActiveRealmsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get active realms o k response has a 5xx status code
func (o *GetActiveRealmsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get active realms o k response a status code equal to that given
func (o *GetActiveRealmsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get active realms o k response
func (o *GetActiveRealmsOK) Code() int {
	return 200
}

func (o *GetActiveRealmsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/security/realms/active][%d] getActiveRealmsOK %s", 200, payload)
}

func (o *GetActiveRealmsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/security/realms/active][%d] getActiveRealmsOK %s", 200, payload)
}

func (o *GetActiveRealmsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetActiveRealmsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
