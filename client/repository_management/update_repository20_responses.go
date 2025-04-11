// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository20Reader is a Reader for the UpdateRepository20 structure.
type UpdateRepository20Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository20Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository20NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRepository20BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateRepository20Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository20Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}] updateRepository_20", response, response.Code())
	}
}

// NewUpdateRepository20NoContent creates a UpdateRepository20NoContent with default headers values
func NewUpdateRepository20NoContent() *UpdateRepository20NoContent {
	return &UpdateRepository20NoContent{}
}

/*
UpdateRepository20NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository20NoContent struct {
}

// IsSuccess returns true when this update repository20 no content response has a 2xx status code
func (o *UpdateRepository20NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository20 no content response has a 3xx status code
func (o *UpdateRepository20NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository20 no content response has a 4xx status code
func (o *UpdateRepository20NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository20 no content response has a 5xx status code
func (o *UpdateRepository20NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository20 no content response a status code equal to that given
func (o *UpdateRepository20NoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update repository20 no content response
func (o *UpdateRepository20NoContent) Code() int {
	return 204
}

func (o *UpdateRepository20NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] updateRepository20NoContent", 204)
}

func (o *UpdateRepository20NoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] updateRepository20NoContent", 204)
}

func (o *UpdateRepository20NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository20BadRequest creates a UpdateRepository20BadRequest with default headers values
func NewUpdateRepository20BadRequest() *UpdateRepository20BadRequest {
	return &UpdateRepository20BadRequest{}
}

/*
UpdateRepository20BadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateRepository20BadRequest struct {
}

// IsSuccess returns true when this update repository20 bad request response has a 2xx status code
func (o *UpdateRepository20BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository20 bad request response has a 3xx status code
func (o *UpdateRepository20BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository20 bad request response has a 4xx status code
func (o *UpdateRepository20BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository20 bad request response has a 5xx status code
func (o *UpdateRepository20BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository20 bad request response a status code equal to that given
func (o *UpdateRepository20BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update repository20 bad request response
func (o *UpdateRepository20BadRequest) Code() int {
	return 400
}

func (o *UpdateRepository20BadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] updateRepository20BadRequest", 400)
}

func (o *UpdateRepository20BadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] updateRepository20BadRequest", 400)
}

func (o *UpdateRepository20BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository20Unauthorized creates a UpdateRepository20Unauthorized with default headers values
func NewUpdateRepository20Unauthorized() *UpdateRepository20Unauthorized {
	return &UpdateRepository20Unauthorized{}
}

/*
UpdateRepository20Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository20Unauthorized struct {
}

// IsSuccess returns true when this update repository20 unauthorized response has a 2xx status code
func (o *UpdateRepository20Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository20 unauthorized response has a 3xx status code
func (o *UpdateRepository20Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository20 unauthorized response has a 4xx status code
func (o *UpdateRepository20Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository20 unauthorized response has a 5xx status code
func (o *UpdateRepository20Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository20 unauthorized response a status code equal to that given
func (o *UpdateRepository20Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update repository20 unauthorized response
func (o *UpdateRepository20Unauthorized) Code() int {
	return 401
}

func (o *UpdateRepository20Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] updateRepository20Unauthorized", 401)
}

func (o *UpdateRepository20Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] updateRepository20Unauthorized", 401)
}

func (o *UpdateRepository20Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository20Forbidden creates a UpdateRepository20Forbidden with default headers values
func NewUpdateRepository20Forbidden() *UpdateRepository20Forbidden {
	return &UpdateRepository20Forbidden{}
}

/*
UpdateRepository20Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository20Forbidden struct {
}

// IsSuccess returns true when this update repository20 forbidden response has a 2xx status code
func (o *UpdateRepository20Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository20 forbidden response has a 3xx status code
func (o *UpdateRepository20Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository20 forbidden response has a 4xx status code
func (o *UpdateRepository20Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository20 forbidden response has a 5xx status code
func (o *UpdateRepository20Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository20 forbidden response a status code equal to that given
func (o *UpdateRepository20Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update repository20 forbidden response
func (o *UpdateRepository20Forbidden) Code() int {
	return 403
}

func (o *UpdateRepository20Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] updateRepository20Forbidden", 403)
}

func (o *UpdateRepository20Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] updateRepository20Forbidden", 403)
}

func (o *UpdateRepository20Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
