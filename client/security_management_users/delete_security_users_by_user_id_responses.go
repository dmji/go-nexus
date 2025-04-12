// Code generated by go-swagger; DO NOT EDIT.

package security_management_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteSecurityUsersByUserIDReader is a Reader for the DeleteSecurityUsersByUserID structure.
type DeleteSecurityUsersByUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSecurityUsersByUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewDeleteSecurityUsersByUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSecurityUsersByUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSecurityUsersByUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/security/users/{userId}] DeleteSecurityUsersByUserId", response, response.Code())
	}
}

// NewDeleteSecurityUsersByUserIDBadRequest creates a DeleteSecurityUsersByUserIDBadRequest with default headers values
func NewDeleteSecurityUsersByUserIDBadRequest() *DeleteSecurityUsersByUserIDBadRequest {
	return &DeleteSecurityUsersByUserIDBadRequest{}
}

/*
DeleteSecurityUsersByUserIDBadRequest describes a response with status code 400, with default header values.

There was problem deleting a user. Consult the response body for more details
*/
type DeleteSecurityUsersByUserIDBadRequest struct {
}

// IsSuccess returns true when this delete security users by user Id bad request response has a 2xx status code
func (o *DeleteSecurityUsersByUserIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete security users by user Id bad request response has a 3xx status code
func (o *DeleteSecurityUsersByUserIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete security users by user Id bad request response has a 4xx status code
func (o *DeleteSecurityUsersByUserIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete security users by user Id bad request response has a 5xx status code
func (o *DeleteSecurityUsersByUserIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete security users by user Id bad request response a status code equal to that given
func (o *DeleteSecurityUsersByUserIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete security users by user Id bad request response
func (o *DeleteSecurityUsersByUserIDBadRequest) Code() int {
	return 400
}

func (o *DeleteSecurityUsersByUserIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/security/users/{userId}][%d] deleteSecurityUsersByUserIdBadRequest", 400)
}

func (o *DeleteSecurityUsersByUserIDBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/security/users/{userId}][%d] deleteSecurityUsersByUserIdBadRequest", 400)
}

func (o *DeleteSecurityUsersByUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSecurityUsersByUserIDForbidden creates a DeleteSecurityUsersByUserIDForbidden with default headers values
func NewDeleteSecurityUsersByUserIDForbidden() *DeleteSecurityUsersByUserIDForbidden {
	return &DeleteSecurityUsersByUserIDForbidden{}
}

/*
DeleteSecurityUsersByUserIDForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type DeleteSecurityUsersByUserIDForbidden struct {
}

// IsSuccess returns true when this delete security users by user Id forbidden response has a 2xx status code
func (o *DeleteSecurityUsersByUserIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete security users by user Id forbidden response has a 3xx status code
func (o *DeleteSecurityUsersByUserIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete security users by user Id forbidden response has a 4xx status code
func (o *DeleteSecurityUsersByUserIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete security users by user Id forbidden response has a 5xx status code
func (o *DeleteSecurityUsersByUserIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete security users by user Id forbidden response a status code equal to that given
func (o *DeleteSecurityUsersByUserIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete security users by user Id forbidden response
func (o *DeleteSecurityUsersByUserIDForbidden) Code() int {
	return 403
}

func (o *DeleteSecurityUsersByUserIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/security/users/{userId}][%d] deleteSecurityUsersByUserIdForbidden", 403)
}

func (o *DeleteSecurityUsersByUserIDForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/security/users/{userId}][%d] deleteSecurityUsersByUserIdForbidden", 403)
}

func (o *DeleteSecurityUsersByUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSecurityUsersByUserIDNotFound creates a DeleteSecurityUsersByUserIDNotFound with default headers values
func NewDeleteSecurityUsersByUserIDNotFound() *DeleteSecurityUsersByUserIDNotFound {
	return &DeleteSecurityUsersByUserIDNotFound{}
}

/*
DeleteSecurityUsersByUserIDNotFound describes a response with status code 404, with default header values.

User or user source not found in the system.
*/
type DeleteSecurityUsersByUserIDNotFound struct {
}

// IsSuccess returns true when this delete security users by user Id not found response has a 2xx status code
func (o *DeleteSecurityUsersByUserIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete security users by user Id not found response has a 3xx status code
func (o *DeleteSecurityUsersByUserIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete security users by user Id not found response has a 4xx status code
func (o *DeleteSecurityUsersByUserIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete security users by user Id not found response has a 5xx status code
func (o *DeleteSecurityUsersByUserIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete security users by user Id not found response a status code equal to that given
func (o *DeleteSecurityUsersByUserIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete security users by user Id not found response
func (o *DeleteSecurityUsersByUserIDNotFound) Code() int {
	return 404
}

func (o *DeleteSecurityUsersByUserIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/security/users/{userId}][%d] deleteSecurityUsersByUserIdNotFound", 404)
}

func (o *DeleteSecurityUsersByUserIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/security/users/{userId}][%d] deleteSecurityUsersByUserIdNotFound", 404)
}

func (o *DeleteSecurityUsersByUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
