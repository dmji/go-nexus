// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteSecurityPrivilegesByPrivilegeNameReader is a Reader for the DeleteSecurityPrivilegesByPrivilegeName structure.
type DeleteSecurityPrivilegesByPrivilegeNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSecurityPrivilegesByPrivilegeNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewDeleteSecurityPrivilegesByPrivilegeNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSecurityPrivilegesByPrivilegeNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSecurityPrivilegesByPrivilegeNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/security/privileges/{privilegeName}] DeleteSecurityPrivilegesByPrivilegeName", response, response.Code())
	}
}

// NewDeleteSecurityPrivilegesByPrivilegeNameBadRequest creates a DeleteSecurityPrivilegesByPrivilegeNameBadRequest with default headers values
func NewDeleteSecurityPrivilegesByPrivilegeNameBadRequest() *DeleteSecurityPrivilegesByPrivilegeNameBadRequest {
	return &DeleteSecurityPrivilegesByPrivilegeNameBadRequest{}
}

/*
DeleteSecurityPrivilegesByPrivilegeNameBadRequest describes a response with status code 400, with default header values.

The privilege is internal and may not be altered.
*/
type DeleteSecurityPrivilegesByPrivilegeNameBadRequest struct {
}

// IsSuccess returns true when this delete security privileges by privilege name bad request response has a 2xx status code
func (o *DeleteSecurityPrivilegesByPrivilegeNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete security privileges by privilege name bad request response has a 3xx status code
func (o *DeleteSecurityPrivilegesByPrivilegeNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete security privileges by privilege name bad request response has a 4xx status code
func (o *DeleteSecurityPrivilegesByPrivilegeNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete security privileges by privilege name bad request response has a 5xx status code
func (o *DeleteSecurityPrivilegesByPrivilegeNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete security privileges by privilege name bad request response a status code equal to that given
func (o *DeleteSecurityPrivilegesByPrivilegeNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete security privileges by privilege name bad request response
func (o *DeleteSecurityPrivilegesByPrivilegeNameBadRequest) Code() int {
	return 400
}

func (o *DeleteSecurityPrivilegesByPrivilegeNameBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/security/privileges/{privilegeName}][%d] deleteSecurityPrivilegesByPrivilegeNameBadRequest", 400)
}

func (o *DeleteSecurityPrivilegesByPrivilegeNameBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/security/privileges/{privilegeName}][%d] deleteSecurityPrivilegesByPrivilegeNameBadRequest", 400)
}

func (o *DeleteSecurityPrivilegesByPrivilegeNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSecurityPrivilegesByPrivilegeNameForbidden creates a DeleteSecurityPrivilegesByPrivilegeNameForbidden with default headers values
func NewDeleteSecurityPrivilegesByPrivilegeNameForbidden() *DeleteSecurityPrivilegesByPrivilegeNameForbidden {
	return &DeleteSecurityPrivilegesByPrivilegeNameForbidden{}
}

/*
DeleteSecurityPrivilegesByPrivilegeNameForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type DeleteSecurityPrivilegesByPrivilegeNameForbidden struct {
}

// IsSuccess returns true when this delete security privileges by privilege name forbidden response has a 2xx status code
func (o *DeleteSecurityPrivilegesByPrivilegeNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete security privileges by privilege name forbidden response has a 3xx status code
func (o *DeleteSecurityPrivilegesByPrivilegeNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete security privileges by privilege name forbidden response has a 4xx status code
func (o *DeleteSecurityPrivilegesByPrivilegeNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete security privileges by privilege name forbidden response has a 5xx status code
func (o *DeleteSecurityPrivilegesByPrivilegeNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete security privileges by privilege name forbidden response a status code equal to that given
func (o *DeleteSecurityPrivilegesByPrivilegeNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete security privileges by privilege name forbidden response
func (o *DeleteSecurityPrivilegesByPrivilegeNameForbidden) Code() int {
	return 403
}

func (o *DeleteSecurityPrivilegesByPrivilegeNameForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/security/privileges/{privilegeName}][%d] deleteSecurityPrivilegesByPrivilegeNameForbidden", 403)
}

func (o *DeleteSecurityPrivilegesByPrivilegeNameForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/security/privileges/{privilegeName}][%d] deleteSecurityPrivilegesByPrivilegeNameForbidden", 403)
}

func (o *DeleteSecurityPrivilegesByPrivilegeNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSecurityPrivilegesByPrivilegeNameNotFound creates a DeleteSecurityPrivilegesByPrivilegeNameNotFound with default headers values
func NewDeleteSecurityPrivilegesByPrivilegeNameNotFound() *DeleteSecurityPrivilegesByPrivilegeNameNotFound {
	return &DeleteSecurityPrivilegesByPrivilegeNameNotFound{}
}

/*
DeleteSecurityPrivilegesByPrivilegeNameNotFound describes a response with status code 404, with default header values.

Privilege not found in the system.
*/
type DeleteSecurityPrivilegesByPrivilegeNameNotFound struct {
}

// IsSuccess returns true when this delete security privileges by privilege name not found response has a 2xx status code
func (o *DeleteSecurityPrivilegesByPrivilegeNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete security privileges by privilege name not found response has a 3xx status code
func (o *DeleteSecurityPrivilegesByPrivilegeNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete security privileges by privilege name not found response has a 4xx status code
func (o *DeleteSecurityPrivilegesByPrivilegeNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete security privileges by privilege name not found response has a 5xx status code
func (o *DeleteSecurityPrivilegesByPrivilegeNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete security privileges by privilege name not found response a status code equal to that given
func (o *DeleteSecurityPrivilegesByPrivilegeNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete security privileges by privilege name not found response
func (o *DeleteSecurityPrivilegesByPrivilegeNameNotFound) Code() int {
	return 404
}

func (o *DeleteSecurityPrivilegesByPrivilegeNameNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/security/privileges/{privilegeName}][%d] deleteSecurityPrivilegesByPrivilegeNameNotFound", 404)
}

func (o *DeleteSecurityPrivilegesByPrivilegeNameNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/security/privileges/{privilegeName}][%d] deleteSecurityPrivilegesByPrivilegeNameNotFound", 404)
}

func (o *DeleteSecurityPrivilegesByPrivilegeNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
