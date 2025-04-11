// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InvalidateCacheReader is a Reader for the InvalidateCache structure.
type InvalidateCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvalidateCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewInvalidateCacheNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInvalidateCacheBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInvalidateCacheUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInvalidateCacheForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInvalidateCacheNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/{repositoryName}/invalidate-cache] invalidateCache", response, response.Code())
	}
}

// NewInvalidateCacheNoContent creates a InvalidateCacheNoContent with default headers values
func NewInvalidateCacheNoContent() *InvalidateCacheNoContent {
	return &InvalidateCacheNoContent{}
}

/*
InvalidateCacheNoContent describes a response with status code 204, with default header values.

Repository cache invalidated
*/
type InvalidateCacheNoContent struct {
}

// IsSuccess returns true when this invalidate cache no content response has a 2xx status code
func (o *InvalidateCacheNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this invalidate cache no content response has a 3xx status code
func (o *InvalidateCacheNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invalidate cache no content response has a 4xx status code
func (o *InvalidateCacheNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this invalidate cache no content response has a 5xx status code
func (o *InvalidateCacheNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this invalidate cache no content response a status code equal to that given
func (o *InvalidateCacheNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the invalidate cache no content response
func (o *InvalidateCacheNoContent) Code() int {
	return 204
}

func (o *InvalidateCacheNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] invalidateCacheNoContent", 204)
}

func (o *InvalidateCacheNoContent) String() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] invalidateCacheNoContent", 204)
}

func (o *InvalidateCacheNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInvalidateCacheBadRequest creates a InvalidateCacheBadRequest with default headers values
func NewInvalidateCacheBadRequest() *InvalidateCacheBadRequest {
	return &InvalidateCacheBadRequest{}
}

/*
InvalidateCacheBadRequest describes a response with status code 400, with default header values.

Repository is not of proxy or group type
*/
type InvalidateCacheBadRequest struct {
}

// IsSuccess returns true when this invalidate cache bad request response has a 2xx status code
func (o *InvalidateCacheBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invalidate cache bad request response has a 3xx status code
func (o *InvalidateCacheBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invalidate cache bad request response has a 4xx status code
func (o *InvalidateCacheBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this invalidate cache bad request response has a 5xx status code
func (o *InvalidateCacheBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this invalidate cache bad request response a status code equal to that given
func (o *InvalidateCacheBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the invalidate cache bad request response
func (o *InvalidateCacheBadRequest) Code() int {
	return 400
}

func (o *InvalidateCacheBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] invalidateCacheBadRequest", 400)
}

func (o *InvalidateCacheBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] invalidateCacheBadRequest", 400)
}

func (o *InvalidateCacheBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInvalidateCacheUnauthorized creates a InvalidateCacheUnauthorized with default headers values
func NewInvalidateCacheUnauthorized() *InvalidateCacheUnauthorized {
	return &InvalidateCacheUnauthorized{}
}

/*
InvalidateCacheUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type InvalidateCacheUnauthorized struct {
}

// IsSuccess returns true when this invalidate cache unauthorized response has a 2xx status code
func (o *InvalidateCacheUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invalidate cache unauthorized response has a 3xx status code
func (o *InvalidateCacheUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invalidate cache unauthorized response has a 4xx status code
func (o *InvalidateCacheUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this invalidate cache unauthorized response has a 5xx status code
func (o *InvalidateCacheUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this invalidate cache unauthorized response a status code equal to that given
func (o *InvalidateCacheUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the invalidate cache unauthorized response
func (o *InvalidateCacheUnauthorized) Code() int {
	return 401
}

func (o *InvalidateCacheUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] invalidateCacheUnauthorized", 401)
}

func (o *InvalidateCacheUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] invalidateCacheUnauthorized", 401)
}

func (o *InvalidateCacheUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInvalidateCacheForbidden creates a InvalidateCacheForbidden with default headers values
func NewInvalidateCacheForbidden() *InvalidateCacheForbidden {
	return &InvalidateCacheForbidden{}
}

/*
InvalidateCacheForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type InvalidateCacheForbidden struct {
}

// IsSuccess returns true when this invalidate cache forbidden response has a 2xx status code
func (o *InvalidateCacheForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invalidate cache forbidden response has a 3xx status code
func (o *InvalidateCacheForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invalidate cache forbidden response has a 4xx status code
func (o *InvalidateCacheForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this invalidate cache forbidden response has a 5xx status code
func (o *InvalidateCacheForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this invalidate cache forbidden response a status code equal to that given
func (o *InvalidateCacheForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the invalidate cache forbidden response
func (o *InvalidateCacheForbidden) Code() int {
	return 403
}

func (o *InvalidateCacheForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] invalidateCacheForbidden", 403)
}

func (o *InvalidateCacheForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] invalidateCacheForbidden", 403)
}

func (o *InvalidateCacheForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInvalidateCacheNotFound creates a InvalidateCacheNotFound with default headers values
func NewInvalidateCacheNotFound() *InvalidateCacheNotFound {
	return &InvalidateCacheNotFound{}
}

/*
InvalidateCacheNotFound describes a response with status code 404, with default header values.

Repository not found
*/
type InvalidateCacheNotFound struct {
}

// IsSuccess returns true when this invalidate cache not found response has a 2xx status code
func (o *InvalidateCacheNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invalidate cache not found response has a 3xx status code
func (o *InvalidateCacheNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invalidate cache not found response has a 4xx status code
func (o *InvalidateCacheNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this invalidate cache not found response has a 5xx status code
func (o *InvalidateCacheNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this invalidate cache not found response a status code equal to that given
func (o *InvalidateCacheNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the invalidate cache not found response
func (o *InvalidateCacheNotFound) Code() int {
	return 404
}

func (o *InvalidateCacheNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] invalidateCacheNotFound", 404)
}

func (o *InvalidateCacheNotFound) String() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] invalidateCacheNotFound", 404)
}

func (o *InvalidateCacheNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
