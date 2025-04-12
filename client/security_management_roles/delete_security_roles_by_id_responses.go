// Code generated by go-swagger; DO NOT EDIT.

package security_management_roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteSecurityRolesByIDReader is a Reader for the DeleteSecurityRolesByID structure.
type DeleteSecurityRolesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSecurityRolesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 403:
		result := NewDeleteSecurityRolesByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSecurityRolesByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/security/roles/{id}] DeleteSecurityRolesById", response, response.Code())
	}
}

// NewDeleteSecurityRolesByIDForbidden creates a DeleteSecurityRolesByIDForbidden with default headers values
func NewDeleteSecurityRolesByIDForbidden() *DeleteSecurityRolesByIDForbidden {
	return &DeleteSecurityRolesByIDForbidden{}
}

/*
DeleteSecurityRolesByIDForbidden describes a response with status code 403, with default header values.

Insufficient permissions to delete role
*/
type DeleteSecurityRolesByIDForbidden struct {
}

// IsSuccess returns true when this delete security roles by Id forbidden response has a 2xx status code
func (o *DeleteSecurityRolesByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete security roles by Id forbidden response has a 3xx status code
func (o *DeleteSecurityRolesByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete security roles by Id forbidden response has a 4xx status code
func (o *DeleteSecurityRolesByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete security roles by Id forbidden response has a 5xx status code
func (o *DeleteSecurityRolesByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete security roles by Id forbidden response a status code equal to that given
func (o *DeleteSecurityRolesByIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete security roles by Id forbidden response
func (o *DeleteSecurityRolesByIDForbidden) Code() int {
	return 403
}

func (o *DeleteSecurityRolesByIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/security/roles/{id}][%d] deleteSecurityRolesByIdForbidden", 403)
}

func (o *DeleteSecurityRolesByIDForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/security/roles/{id}][%d] deleteSecurityRolesByIdForbidden", 403)
}

func (o *DeleteSecurityRolesByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSecurityRolesByIDNotFound creates a DeleteSecurityRolesByIDNotFound with default headers values
func NewDeleteSecurityRolesByIDNotFound() *DeleteSecurityRolesByIDNotFound {
	return &DeleteSecurityRolesByIDNotFound{}
}

/*
DeleteSecurityRolesByIDNotFound describes a response with status code 404, with default header values.

Role not found
*/
type DeleteSecurityRolesByIDNotFound struct {
}

// IsSuccess returns true when this delete security roles by Id not found response has a 2xx status code
func (o *DeleteSecurityRolesByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete security roles by Id not found response has a 3xx status code
func (o *DeleteSecurityRolesByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete security roles by Id not found response has a 4xx status code
func (o *DeleteSecurityRolesByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete security roles by Id not found response has a 5xx status code
func (o *DeleteSecurityRolesByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete security roles by Id not found response a status code equal to that given
func (o *DeleteSecurityRolesByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete security roles by Id not found response
func (o *DeleteSecurityRolesByIDNotFound) Code() int {
	return 404
}

func (o *DeleteSecurityRolesByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/security/roles/{id}][%d] deleteSecurityRolesByIdNotFound", 404)
}

func (o *DeleteSecurityRolesByIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/security/roles/{id}][%d] deleteSecurityRolesByIdNotFound", 404)
}

func (o *DeleteSecurityRolesByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
