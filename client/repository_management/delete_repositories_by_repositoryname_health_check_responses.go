// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteRepositoriesByRepositorynameHealthCheckReader is a Reader for the DeleteRepositoriesByRepositorynameHealthCheck structure.
type DeleteRepositoriesByRepositorynameHealthCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRepositoriesByRepositorynameHealthCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRepositoriesByRepositorynameHealthCheckNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteRepositoriesByRepositorynameHealthCheckUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRepositoriesByRepositorynameHealthCheckForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRepositoriesByRepositorynameHealthCheckNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/repositories/{repositoryName}/health-check] DeleteRepositoriesByRepositorynameHealthCheck", response, response.Code())
	}
}

// NewDeleteRepositoriesByRepositorynameHealthCheckNoContent creates a DeleteRepositoriesByRepositorynameHealthCheckNoContent with default headers values
func NewDeleteRepositoriesByRepositorynameHealthCheckNoContent() *DeleteRepositoriesByRepositorynameHealthCheckNoContent {
	return &DeleteRepositoriesByRepositorynameHealthCheckNoContent{}
}

/*
DeleteRepositoriesByRepositorynameHealthCheckNoContent describes a response with status code 204, with default header values.

Repository Health Check disabled
*/
type DeleteRepositoriesByRepositorynameHealthCheckNoContent struct {
}

// IsSuccess returns true when this delete repositories by repositoryname health check no content response has a 2xx status code
func (o *DeleteRepositoriesByRepositorynameHealthCheckNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete repositories by repositoryname health check no content response has a 3xx status code
func (o *DeleteRepositoriesByRepositorynameHealthCheckNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete repositories by repositoryname health check no content response has a 4xx status code
func (o *DeleteRepositoriesByRepositorynameHealthCheckNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete repositories by repositoryname health check no content response has a 5xx status code
func (o *DeleteRepositoriesByRepositorynameHealthCheckNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete repositories by repositoryname health check no content response a status code equal to that given
func (o *DeleteRepositoriesByRepositorynameHealthCheckNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete repositories by repositoryname health check no content response
func (o *DeleteRepositoriesByRepositorynameHealthCheckNoContent) Code() int {
	return 204
}

func (o *DeleteRepositoriesByRepositorynameHealthCheckNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/repositories/{repositoryName}/health-check][%d] deleteRepositoriesByRepositorynameHealthCheckNoContent", 204)
}

func (o *DeleteRepositoriesByRepositorynameHealthCheckNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/repositories/{repositoryName}/health-check][%d] deleteRepositoriesByRepositorynameHealthCheckNoContent", 204)
}

func (o *DeleteRepositoriesByRepositorynameHealthCheckNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepositoriesByRepositorynameHealthCheckUnauthorized creates a DeleteRepositoriesByRepositorynameHealthCheckUnauthorized with default headers values
func NewDeleteRepositoriesByRepositorynameHealthCheckUnauthorized() *DeleteRepositoriesByRepositorynameHealthCheckUnauthorized {
	return &DeleteRepositoriesByRepositorynameHealthCheckUnauthorized{}
}

/*
DeleteRepositoriesByRepositorynameHealthCheckUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type DeleteRepositoriesByRepositorynameHealthCheckUnauthorized struct {
}

// IsSuccess returns true when this delete repositories by repositoryname health check unauthorized response has a 2xx status code
func (o *DeleteRepositoriesByRepositorynameHealthCheckUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete repositories by repositoryname health check unauthorized response has a 3xx status code
func (o *DeleteRepositoriesByRepositorynameHealthCheckUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete repositories by repositoryname health check unauthorized response has a 4xx status code
func (o *DeleteRepositoriesByRepositorynameHealthCheckUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete repositories by repositoryname health check unauthorized response has a 5xx status code
func (o *DeleteRepositoriesByRepositorynameHealthCheckUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete repositories by repositoryname health check unauthorized response a status code equal to that given
func (o *DeleteRepositoriesByRepositorynameHealthCheckUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete repositories by repositoryname health check unauthorized response
func (o *DeleteRepositoriesByRepositorynameHealthCheckUnauthorized) Code() int {
	return 401
}

func (o *DeleteRepositoriesByRepositorynameHealthCheckUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/repositories/{repositoryName}/health-check][%d] deleteRepositoriesByRepositorynameHealthCheckUnauthorized", 401)
}

func (o *DeleteRepositoriesByRepositorynameHealthCheckUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/repositories/{repositoryName}/health-check][%d] deleteRepositoriesByRepositorynameHealthCheckUnauthorized", 401)
}

func (o *DeleteRepositoriesByRepositorynameHealthCheckUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepositoriesByRepositorynameHealthCheckForbidden creates a DeleteRepositoriesByRepositorynameHealthCheckForbidden with default headers values
func NewDeleteRepositoriesByRepositorynameHealthCheckForbidden() *DeleteRepositoriesByRepositorynameHealthCheckForbidden {
	return &DeleteRepositoriesByRepositorynameHealthCheckForbidden{}
}

/*
DeleteRepositoriesByRepositorynameHealthCheckForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type DeleteRepositoriesByRepositorynameHealthCheckForbidden struct {
}

// IsSuccess returns true when this delete repositories by repositoryname health check forbidden response has a 2xx status code
func (o *DeleteRepositoriesByRepositorynameHealthCheckForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete repositories by repositoryname health check forbidden response has a 3xx status code
func (o *DeleteRepositoriesByRepositorynameHealthCheckForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete repositories by repositoryname health check forbidden response has a 4xx status code
func (o *DeleteRepositoriesByRepositorynameHealthCheckForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete repositories by repositoryname health check forbidden response has a 5xx status code
func (o *DeleteRepositoriesByRepositorynameHealthCheckForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete repositories by repositoryname health check forbidden response a status code equal to that given
func (o *DeleteRepositoriesByRepositorynameHealthCheckForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete repositories by repositoryname health check forbidden response
func (o *DeleteRepositoriesByRepositorynameHealthCheckForbidden) Code() int {
	return 403
}

func (o *DeleteRepositoriesByRepositorynameHealthCheckForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/repositories/{repositoryName}/health-check][%d] deleteRepositoriesByRepositorynameHealthCheckForbidden", 403)
}

func (o *DeleteRepositoriesByRepositorynameHealthCheckForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/repositories/{repositoryName}/health-check][%d] deleteRepositoriesByRepositorynameHealthCheckForbidden", 403)
}

func (o *DeleteRepositoriesByRepositorynameHealthCheckForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepositoriesByRepositorynameHealthCheckNotFound creates a DeleteRepositoriesByRepositorynameHealthCheckNotFound with default headers values
func NewDeleteRepositoriesByRepositorynameHealthCheckNotFound() *DeleteRepositoriesByRepositorynameHealthCheckNotFound {
	return &DeleteRepositoriesByRepositorynameHealthCheckNotFound{}
}

/*
DeleteRepositoriesByRepositorynameHealthCheckNotFound describes a response with status code 404, with default header values.

Repository not found
*/
type DeleteRepositoriesByRepositorynameHealthCheckNotFound struct {
}

// IsSuccess returns true when this delete repositories by repositoryname health check not found response has a 2xx status code
func (o *DeleteRepositoriesByRepositorynameHealthCheckNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete repositories by repositoryname health check not found response has a 3xx status code
func (o *DeleteRepositoriesByRepositorynameHealthCheckNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete repositories by repositoryname health check not found response has a 4xx status code
func (o *DeleteRepositoriesByRepositorynameHealthCheckNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete repositories by repositoryname health check not found response has a 5xx status code
func (o *DeleteRepositoriesByRepositorynameHealthCheckNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete repositories by repositoryname health check not found response a status code equal to that given
func (o *DeleteRepositoriesByRepositorynameHealthCheckNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete repositories by repositoryname health check not found response
func (o *DeleteRepositoriesByRepositorynameHealthCheckNotFound) Code() int {
	return 404
}

func (o *DeleteRepositoriesByRepositorynameHealthCheckNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/repositories/{repositoryName}/health-check][%d] deleteRepositoriesByRepositorynameHealthCheckNotFound", 404)
}

func (o *DeleteRepositoriesByRepositorynameHealthCheckNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/repositories/{repositoryName}/health-check][%d] deleteRepositoriesByRepositorynameHealthCheckNotFound", 404)
}

func (o *DeleteRepositoriesByRepositorynameHealthCheckNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
