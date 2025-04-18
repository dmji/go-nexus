// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutSecurityPrivilegesRepositoryViewByPrivilegeNameReader is a Reader for the PutSecurityPrivilegesRepositoryViewByPrivilegeName structure.
type PutSecurityPrivilegesRepositoryViewByPrivilegeNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/security/privileges/repository-view/{privilegeName}] PutSecurityPrivilegesRepositoryViewByPrivilegeName", response, response.Code())
	}
}

// NewPutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest creates a PutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest with default headers values
func NewPutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest() *PutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest {
	return &PutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest{}
}

/*
PutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest describes a response with status code 400, with default header values.

Privilege object not configured properly.
*/
type PutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest struct {
}

// IsSuccess returns true when this put security privileges repository view by privilege name bad request response has a 2xx status code
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put security privileges repository view by privilege name bad request response has a 3xx status code
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put security privileges repository view by privilege name bad request response has a 4xx status code
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put security privileges repository view by privilege name bad request response has a 5xx status code
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put security privileges repository view by privilege name bad request response a status code equal to that given
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put security privileges repository view by privilege name bad request response
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest) Code() int {
	return 400
}

func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-view/{privilegeName}][%d] putSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest", 400)
}

func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-view/{privilegeName}][%d] putSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest", 400)
}

func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden creates a PutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden with default headers values
func NewPutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden() *PutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden {
	return &PutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden{}
}

/*
PutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type PutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden struct {
}

// IsSuccess returns true when this put security privileges repository view by privilege name forbidden response has a 2xx status code
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put security privileges repository view by privilege name forbidden response has a 3xx status code
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put security privileges repository view by privilege name forbidden response has a 4xx status code
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put security privileges repository view by privilege name forbidden response has a 5xx status code
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put security privileges repository view by privilege name forbidden response a status code equal to that given
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put security privileges repository view by privilege name forbidden response
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden) Code() int {
	return 403
}

func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-view/{privilegeName}][%d] putSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden", 403)
}

func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-view/{privilegeName}][%d] putSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden", 403)
}

func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound creates a PutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound with default headers values
func NewPutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound() *PutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound {
	return &PutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound{}
}

/*
PutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound describes a response with status code 404, with default header values.

Privilege not found in the system.
*/
type PutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound struct {
}

// IsSuccess returns true when this put security privileges repository view by privilege name not found response has a 2xx status code
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put security privileges repository view by privilege name not found response has a 3xx status code
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put security privileges repository view by privilege name not found response has a 4xx status code
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put security privileges repository view by privilege name not found response has a 5xx status code
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put security privileges repository view by privilege name not found response a status code equal to that given
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put security privileges repository view by privilege name not found response
func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound) Code() int {
	return 404
}

func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-view/{privilegeName}][%d] putSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound", 404)
}

func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-view/{privilegeName}][%d] putSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound", 404)
}

func (o *PutSecurityPrivilegesRepositoryViewByPrivilegeNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
