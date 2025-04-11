// Code generated by go-swagger; DO NOT EDIT.

package product_licensing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"nexus/models"
)

// SetLicenseReader is a Reader for the SetLicense structure.
type SetLicenseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetLicenseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetLicenseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/system/license] setLicense", response, response.Code())
	}
}

// NewSetLicenseOK creates a SetLicenseOK with default headers values
func NewSetLicenseOK() *SetLicenseOK {
	return &SetLicenseOK{}
}

/*
SetLicenseOK describes a response with status code 200, with default header values.

successful operation
*/
type SetLicenseOK struct {
	Payload *models.APILicenseDetailsXO
}

// IsSuccess returns true when this set license o k response has a 2xx status code
func (o *SetLicenseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set license o k response has a 3xx status code
func (o *SetLicenseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set license o k response has a 4xx status code
func (o *SetLicenseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set license o k response has a 5xx status code
func (o *SetLicenseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set license o k response a status code equal to that given
func (o *SetLicenseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set license o k response
func (o *SetLicenseOK) Code() int {
	return 200
}

func (o *SetLicenseOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/system/license][%d] setLicenseOK %s", 200, payload)
}

func (o *SetLicenseOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/system/license][%d] setLicenseOK %s", 200, payload)
}

func (o *SetLicenseOK) GetPayload() *models.APILicenseDetailsXO {
	return o.Payload
}

func (o *SetLicenseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APILicenseDetailsXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
