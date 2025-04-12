// Code generated by go-swagger; DO NOT EDIT.

package security_management_anonymous_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dmji/go-nexus/models"
)

// PutSecurityAnonymousReader is a Reader for the PutSecurityAnonymous structure.
type PutSecurityAnonymousReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSecurityAnonymousReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutSecurityAnonymousOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPutSecurityAnonymousForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/security/anonymous] PutSecurityAnonymous", response, response.Code())
	}
}

// NewPutSecurityAnonymousOK creates a PutSecurityAnonymousOK with default headers values
func NewPutSecurityAnonymousOK() *PutSecurityAnonymousOK {
	return &PutSecurityAnonymousOK{}
}

/*
PutSecurityAnonymousOK describes a response with status code 200, with default header values.

successful operation
*/
type PutSecurityAnonymousOK struct {
	Payload *models.AnonymousAccessSettingsXO
}

// IsSuccess returns true when this put security anonymous o k response has a 2xx status code
func (o *PutSecurityAnonymousOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put security anonymous o k response has a 3xx status code
func (o *PutSecurityAnonymousOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put security anonymous o k response has a 4xx status code
func (o *PutSecurityAnonymousOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put security anonymous o k response has a 5xx status code
func (o *PutSecurityAnonymousOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put security anonymous o k response a status code equal to that given
func (o *PutSecurityAnonymousOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put security anonymous o k response
func (o *PutSecurityAnonymousOK) Code() int {
	return 200
}

func (o *PutSecurityAnonymousOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/security/anonymous][%d] putSecurityAnonymousOK %s", 200, payload)
}

func (o *PutSecurityAnonymousOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/security/anonymous][%d] putSecurityAnonymousOK %s", 200, payload)
}

func (o *PutSecurityAnonymousOK) GetPayload() *models.AnonymousAccessSettingsXO {
	return o.Payload
}

func (o *PutSecurityAnonymousOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnonymousAccessSettingsXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSecurityAnonymousForbidden creates a PutSecurityAnonymousForbidden with default headers values
func NewPutSecurityAnonymousForbidden() *PutSecurityAnonymousForbidden {
	return &PutSecurityAnonymousForbidden{}
}

/*
PutSecurityAnonymousForbidden describes a response with status code 403, with default header values.

Insufficient permissions to update settings
*/
type PutSecurityAnonymousForbidden struct {
}

// IsSuccess returns true when this put security anonymous forbidden response has a 2xx status code
func (o *PutSecurityAnonymousForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put security anonymous forbidden response has a 3xx status code
func (o *PutSecurityAnonymousForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put security anonymous forbidden response has a 4xx status code
func (o *PutSecurityAnonymousForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put security anonymous forbidden response has a 5xx status code
func (o *PutSecurityAnonymousForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put security anonymous forbidden response a status code equal to that given
func (o *PutSecurityAnonymousForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put security anonymous forbidden response
func (o *PutSecurityAnonymousForbidden) Code() int {
	return 403
}

func (o *PutSecurityAnonymousForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/security/anonymous][%d] putSecurityAnonymousForbidden", 403)
}

func (o *PutSecurityAnonymousForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/security/anonymous][%d] putSecurityAnonymousForbidden", 403)
}

func (o *PutSecurityAnonymousForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
