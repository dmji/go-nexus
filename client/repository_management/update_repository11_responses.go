// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository11Reader is a Reader for the UpdateRepository11 structure.
type UpdateRepository11Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository11Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository11NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRepository11BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateRepository11Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository11Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/cocoapods/proxy/{repositoryName}] updateRepository_11", response, response.Code())
	}
}

// NewUpdateRepository11NoContent creates a UpdateRepository11NoContent with default headers values
func NewUpdateRepository11NoContent() *UpdateRepository11NoContent {
	return &UpdateRepository11NoContent{}
}

/*
UpdateRepository11NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository11NoContent struct {
}

// IsSuccess returns true when this update repository11 no content response has a 2xx status code
func (o *UpdateRepository11NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository11 no content response has a 3xx status code
func (o *UpdateRepository11NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository11 no content response has a 4xx status code
func (o *UpdateRepository11NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository11 no content response has a 5xx status code
func (o *UpdateRepository11NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository11 no content response a status code equal to that given
func (o *UpdateRepository11NoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update repository11 no content response
func (o *UpdateRepository11NoContent) Code() int {
	return 204
}

func (o *UpdateRepository11NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/cocoapods/proxy/{repositoryName}][%d] updateRepository11NoContent", 204)
}

func (o *UpdateRepository11NoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/cocoapods/proxy/{repositoryName}][%d] updateRepository11NoContent", 204)
}

func (o *UpdateRepository11NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository11BadRequest creates a UpdateRepository11BadRequest with default headers values
func NewUpdateRepository11BadRequest() *UpdateRepository11BadRequest {
	return &UpdateRepository11BadRequest{}
}

/*
UpdateRepository11BadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateRepository11BadRequest struct {
}

// IsSuccess returns true when this update repository11 bad request response has a 2xx status code
func (o *UpdateRepository11BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository11 bad request response has a 3xx status code
func (o *UpdateRepository11BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository11 bad request response has a 4xx status code
func (o *UpdateRepository11BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository11 bad request response has a 5xx status code
func (o *UpdateRepository11BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository11 bad request response a status code equal to that given
func (o *UpdateRepository11BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update repository11 bad request response
func (o *UpdateRepository11BadRequest) Code() int {
	return 400
}

func (o *UpdateRepository11BadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/cocoapods/proxy/{repositoryName}][%d] updateRepository11BadRequest", 400)
}

func (o *UpdateRepository11BadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/cocoapods/proxy/{repositoryName}][%d] updateRepository11BadRequest", 400)
}

func (o *UpdateRepository11BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository11Unauthorized creates a UpdateRepository11Unauthorized with default headers values
func NewUpdateRepository11Unauthorized() *UpdateRepository11Unauthorized {
	return &UpdateRepository11Unauthorized{}
}

/*
UpdateRepository11Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository11Unauthorized struct {
}

// IsSuccess returns true when this update repository11 unauthorized response has a 2xx status code
func (o *UpdateRepository11Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository11 unauthorized response has a 3xx status code
func (o *UpdateRepository11Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository11 unauthorized response has a 4xx status code
func (o *UpdateRepository11Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository11 unauthorized response has a 5xx status code
func (o *UpdateRepository11Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository11 unauthorized response a status code equal to that given
func (o *UpdateRepository11Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update repository11 unauthorized response
func (o *UpdateRepository11Unauthorized) Code() int {
	return 401
}

func (o *UpdateRepository11Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/cocoapods/proxy/{repositoryName}][%d] updateRepository11Unauthorized", 401)
}

func (o *UpdateRepository11Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/cocoapods/proxy/{repositoryName}][%d] updateRepository11Unauthorized", 401)
}

func (o *UpdateRepository11Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository11Forbidden creates a UpdateRepository11Forbidden with default headers values
func NewUpdateRepository11Forbidden() *UpdateRepository11Forbidden {
	return &UpdateRepository11Forbidden{}
}

/*
UpdateRepository11Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository11Forbidden struct {
}

// IsSuccess returns true when this update repository11 forbidden response has a 2xx status code
func (o *UpdateRepository11Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository11 forbidden response has a 3xx status code
func (o *UpdateRepository11Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository11 forbidden response has a 4xx status code
func (o *UpdateRepository11Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository11 forbidden response has a 5xx status code
func (o *UpdateRepository11Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository11 forbidden response a status code equal to that given
func (o *UpdateRepository11Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update repository11 forbidden response
func (o *UpdateRepository11Forbidden) Code() int {
	return 403
}

func (o *UpdateRepository11Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/cocoapods/proxy/{repositoryName}][%d] updateRepository11Forbidden", 403)
}

func (o *UpdateRepository11Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/cocoapods/proxy/{repositoryName}][%d] updateRepository11Forbidden", 403)
}

func (o *UpdateRepository11Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
