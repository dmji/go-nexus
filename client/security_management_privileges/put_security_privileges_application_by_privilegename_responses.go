// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutSecurityPrivilegesApplicationByPrivilegenameReader is a Reader for the PutSecurityPrivilegesApplicationByPrivilegename structure.
type PutSecurityPrivilegesApplicationByPrivilegenameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSecurityPrivilegesApplicationByPrivilegenameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPutSecurityPrivilegesApplicationByPrivilegenameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutSecurityPrivilegesApplicationByPrivilegenameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutSecurityPrivilegesApplicationByPrivilegenameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/security/privileges/application/{privilegeName}] PutSecurityPrivilegesApplicationByPrivilegename", response, response.Code())
	}
}

// NewPutSecurityPrivilegesApplicationByPrivilegenameBadRequest creates a PutSecurityPrivilegesApplicationByPrivilegenameBadRequest with default headers values
func NewPutSecurityPrivilegesApplicationByPrivilegenameBadRequest() *PutSecurityPrivilegesApplicationByPrivilegenameBadRequest {
	return &PutSecurityPrivilegesApplicationByPrivilegenameBadRequest{}
}

/*
PutSecurityPrivilegesApplicationByPrivilegenameBadRequest describes a response with status code 400, with default header values.

Privilege object not configured properly.
*/
type PutSecurityPrivilegesApplicationByPrivilegenameBadRequest struct {
}

// IsSuccess returns true when this put security privileges application by privilegename bad request response has a 2xx status code
func (o *PutSecurityPrivilegesApplicationByPrivilegenameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put security privileges application by privilegename bad request response has a 3xx status code
func (o *PutSecurityPrivilegesApplicationByPrivilegenameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put security privileges application by privilegename bad request response has a 4xx status code
func (o *PutSecurityPrivilegesApplicationByPrivilegenameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put security privileges application by privilegename bad request response has a 5xx status code
func (o *PutSecurityPrivilegesApplicationByPrivilegenameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put security privileges application by privilegename bad request response a status code equal to that given
func (o *PutSecurityPrivilegesApplicationByPrivilegenameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put security privileges application by privilegename bad request response
func (o *PutSecurityPrivilegesApplicationByPrivilegenameBadRequest) Code() int {
	return 400
}

func (o *PutSecurityPrivilegesApplicationByPrivilegenameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/application/{privilegeName}][%d] putSecurityPrivilegesApplicationByPrivilegenameBadRequest", 400)
}

func (o *PutSecurityPrivilegesApplicationByPrivilegenameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/application/{privilegeName}][%d] putSecurityPrivilegesApplicationByPrivilegenameBadRequest", 400)
}

func (o *PutSecurityPrivilegesApplicationByPrivilegenameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSecurityPrivilegesApplicationByPrivilegenameForbidden creates a PutSecurityPrivilegesApplicationByPrivilegenameForbidden with default headers values
func NewPutSecurityPrivilegesApplicationByPrivilegenameForbidden() *PutSecurityPrivilegesApplicationByPrivilegenameForbidden {
	return &PutSecurityPrivilegesApplicationByPrivilegenameForbidden{}
}

/*
PutSecurityPrivilegesApplicationByPrivilegenameForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type PutSecurityPrivilegesApplicationByPrivilegenameForbidden struct {
}

// IsSuccess returns true when this put security privileges application by privilegename forbidden response has a 2xx status code
func (o *PutSecurityPrivilegesApplicationByPrivilegenameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put security privileges application by privilegename forbidden response has a 3xx status code
func (o *PutSecurityPrivilegesApplicationByPrivilegenameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put security privileges application by privilegename forbidden response has a 4xx status code
func (o *PutSecurityPrivilegesApplicationByPrivilegenameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put security privileges application by privilegename forbidden response has a 5xx status code
func (o *PutSecurityPrivilegesApplicationByPrivilegenameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put security privileges application by privilegename forbidden response a status code equal to that given
func (o *PutSecurityPrivilegesApplicationByPrivilegenameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put security privileges application by privilegename forbidden response
func (o *PutSecurityPrivilegesApplicationByPrivilegenameForbidden) Code() int {
	return 403
}

func (o *PutSecurityPrivilegesApplicationByPrivilegenameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/application/{privilegeName}][%d] putSecurityPrivilegesApplicationByPrivilegenameForbidden", 403)
}

func (o *PutSecurityPrivilegesApplicationByPrivilegenameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/application/{privilegeName}][%d] putSecurityPrivilegesApplicationByPrivilegenameForbidden", 403)
}

func (o *PutSecurityPrivilegesApplicationByPrivilegenameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSecurityPrivilegesApplicationByPrivilegenameNotFound creates a PutSecurityPrivilegesApplicationByPrivilegenameNotFound with default headers values
func NewPutSecurityPrivilegesApplicationByPrivilegenameNotFound() *PutSecurityPrivilegesApplicationByPrivilegenameNotFound {
	return &PutSecurityPrivilegesApplicationByPrivilegenameNotFound{}
}

/*
PutSecurityPrivilegesApplicationByPrivilegenameNotFound describes a response with status code 404, with default header values.

Privilege not found in the system.
*/
type PutSecurityPrivilegesApplicationByPrivilegenameNotFound struct {
}

// IsSuccess returns true when this put security privileges application by privilegename not found response has a 2xx status code
func (o *PutSecurityPrivilegesApplicationByPrivilegenameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put security privileges application by privilegename not found response has a 3xx status code
func (o *PutSecurityPrivilegesApplicationByPrivilegenameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put security privileges application by privilegename not found response has a 4xx status code
func (o *PutSecurityPrivilegesApplicationByPrivilegenameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put security privileges application by privilegename not found response has a 5xx status code
func (o *PutSecurityPrivilegesApplicationByPrivilegenameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put security privileges application by privilegename not found response a status code equal to that given
func (o *PutSecurityPrivilegesApplicationByPrivilegenameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put security privileges application by privilegename not found response
func (o *PutSecurityPrivilegesApplicationByPrivilegenameNotFound) Code() int {
	return 404
}

func (o *PutSecurityPrivilegesApplicationByPrivilegenameNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/application/{privilegeName}][%d] putSecurityPrivilegesApplicationByPrivilegenameNotFound", 404)
}

func (o *PutSecurityPrivilegesApplicationByPrivilegenameNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/application/{privilegeName}][%d] putSecurityPrivilegesApplicationByPrivilegenameNotFound", 404)
}

func (o *PutSecurityPrivilegesApplicationByPrivilegenameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
