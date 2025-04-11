// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository25Reader is a Reader for the UpdateRepository25 structure.
type UpdateRepository25Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository25Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository25NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRepository25BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateRepository25Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository25Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/huggingface/proxy/{repositoryName}] updateRepository_25", response, response.Code())
	}
}

// NewUpdateRepository25NoContent creates a UpdateRepository25NoContent with default headers values
func NewUpdateRepository25NoContent() *UpdateRepository25NoContent {
	return &UpdateRepository25NoContent{}
}

/*
UpdateRepository25NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository25NoContent struct {
}

// IsSuccess returns true when this update repository25 no content response has a 2xx status code
func (o *UpdateRepository25NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository25 no content response has a 3xx status code
func (o *UpdateRepository25NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository25 no content response has a 4xx status code
func (o *UpdateRepository25NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository25 no content response has a 5xx status code
func (o *UpdateRepository25NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository25 no content response a status code equal to that given
func (o *UpdateRepository25NoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update repository25 no content response
func (o *UpdateRepository25NoContent) Code() int {
	return 204
}

func (o *UpdateRepository25NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/huggingface/proxy/{repositoryName}][%d] updateRepository25NoContent", 204)
}

func (o *UpdateRepository25NoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/huggingface/proxy/{repositoryName}][%d] updateRepository25NoContent", 204)
}

func (o *UpdateRepository25NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository25BadRequest creates a UpdateRepository25BadRequest with default headers values
func NewUpdateRepository25BadRequest() *UpdateRepository25BadRequest {
	return &UpdateRepository25BadRequest{}
}

/*
UpdateRepository25BadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateRepository25BadRequest struct {
}

// IsSuccess returns true when this update repository25 bad request response has a 2xx status code
func (o *UpdateRepository25BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository25 bad request response has a 3xx status code
func (o *UpdateRepository25BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository25 bad request response has a 4xx status code
func (o *UpdateRepository25BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository25 bad request response has a 5xx status code
func (o *UpdateRepository25BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository25 bad request response a status code equal to that given
func (o *UpdateRepository25BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update repository25 bad request response
func (o *UpdateRepository25BadRequest) Code() int {
	return 400
}

func (o *UpdateRepository25BadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/huggingface/proxy/{repositoryName}][%d] updateRepository25BadRequest", 400)
}

func (o *UpdateRepository25BadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/huggingface/proxy/{repositoryName}][%d] updateRepository25BadRequest", 400)
}

func (o *UpdateRepository25BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository25Unauthorized creates a UpdateRepository25Unauthorized with default headers values
func NewUpdateRepository25Unauthorized() *UpdateRepository25Unauthorized {
	return &UpdateRepository25Unauthorized{}
}

/*
UpdateRepository25Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository25Unauthorized struct {
}

// IsSuccess returns true when this update repository25 unauthorized response has a 2xx status code
func (o *UpdateRepository25Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository25 unauthorized response has a 3xx status code
func (o *UpdateRepository25Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository25 unauthorized response has a 4xx status code
func (o *UpdateRepository25Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository25 unauthorized response has a 5xx status code
func (o *UpdateRepository25Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository25 unauthorized response a status code equal to that given
func (o *UpdateRepository25Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update repository25 unauthorized response
func (o *UpdateRepository25Unauthorized) Code() int {
	return 401
}

func (o *UpdateRepository25Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/huggingface/proxy/{repositoryName}][%d] updateRepository25Unauthorized", 401)
}

func (o *UpdateRepository25Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/huggingface/proxy/{repositoryName}][%d] updateRepository25Unauthorized", 401)
}

func (o *UpdateRepository25Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository25Forbidden creates a UpdateRepository25Forbidden with default headers values
func NewUpdateRepository25Forbidden() *UpdateRepository25Forbidden {
	return &UpdateRepository25Forbidden{}
}

/*
UpdateRepository25Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository25Forbidden struct {
}

// IsSuccess returns true when this update repository25 forbidden response has a 2xx status code
func (o *UpdateRepository25Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository25 forbidden response has a 3xx status code
func (o *UpdateRepository25Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository25 forbidden response has a 4xx status code
func (o *UpdateRepository25Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository25 forbidden response has a 5xx status code
func (o *UpdateRepository25Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository25 forbidden response a status code equal to that given
func (o *UpdateRepository25Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update repository25 forbidden response
func (o *UpdateRepository25Forbidden) Code() int {
	return 403
}

func (o *UpdateRepository25Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/huggingface/proxy/{repositoryName}][%d] updateRepository25Forbidden", 403)
}

func (o *UpdateRepository25Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/huggingface/proxy/{repositoryName}][%d] updateRepository25Forbidden", 403)
}

func (o *UpdateRepository25Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
