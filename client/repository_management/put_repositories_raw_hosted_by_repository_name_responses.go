// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRepositoriesRawHostedByRepositoryNameReader is a Reader for the PutRepositoriesRawHostedByRepositoryName structure.
type PutRepositoriesRawHostedByRepositoryNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesRawHostedByRepositoryNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesRawHostedByRepositoryNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRepositoriesRawHostedByRepositoryNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRepositoriesRawHostedByRepositoryNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesRawHostedByRepositoryNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/raw/hosted/{repositoryName}] PutRepositoriesRawHostedByRepositoryName", response, response.Code())
	}
}

// NewPutRepositoriesRawHostedByRepositoryNameNoContent creates a PutRepositoriesRawHostedByRepositoryNameNoContent with default headers values
func NewPutRepositoriesRawHostedByRepositoryNameNoContent() *PutRepositoriesRawHostedByRepositoryNameNoContent {
	return &PutRepositoriesRawHostedByRepositoryNameNoContent{}
}

/*
PutRepositoriesRawHostedByRepositoryNameNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutRepositoriesRawHostedByRepositoryNameNoContent struct {
}

// IsSuccess returns true when this put repositories raw hosted by repository name no content response has a 2xx status code
func (o *PutRepositoriesRawHostedByRepositoryNameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put repositories raw hosted by repository name no content response has a 3xx status code
func (o *PutRepositoriesRawHostedByRepositoryNameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories raw hosted by repository name no content response has a 4xx status code
func (o *PutRepositoriesRawHostedByRepositoryNameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put repositories raw hosted by repository name no content response has a 5xx status code
func (o *PutRepositoriesRawHostedByRepositoryNameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories raw hosted by repository name no content response a status code equal to that given
func (o *PutRepositoriesRawHostedByRepositoryNameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put repositories raw hosted by repository name no content response
func (o *PutRepositoriesRawHostedByRepositoryNameNoContent) Code() int {
	return 204
}

func (o *PutRepositoriesRawHostedByRepositoryNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/hosted/{repositoryName}][%d] putRepositoriesRawHostedByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesRawHostedByRepositoryNameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/hosted/{repositoryName}][%d] putRepositoriesRawHostedByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesRawHostedByRepositoryNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesRawHostedByRepositoryNameBadRequest creates a PutRepositoriesRawHostedByRepositoryNameBadRequest with default headers values
func NewPutRepositoriesRawHostedByRepositoryNameBadRequest() *PutRepositoriesRawHostedByRepositoryNameBadRequest {
	return &PutRepositoriesRawHostedByRepositoryNameBadRequest{}
}

/*
PutRepositoriesRawHostedByRepositoryNameBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutRepositoriesRawHostedByRepositoryNameBadRequest struct {
}

// IsSuccess returns true when this put repositories raw hosted by repository name bad request response has a 2xx status code
func (o *PutRepositoriesRawHostedByRepositoryNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories raw hosted by repository name bad request response has a 3xx status code
func (o *PutRepositoriesRawHostedByRepositoryNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories raw hosted by repository name bad request response has a 4xx status code
func (o *PutRepositoriesRawHostedByRepositoryNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories raw hosted by repository name bad request response has a 5xx status code
func (o *PutRepositoriesRawHostedByRepositoryNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories raw hosted by repository name bad request response a status code equal to that given
func (o *PutRepositoriesRawHostedByRepositoryNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put repositories raw hosted by repository name bad request response
func (o *PutRepositoriesRawHostedByRepositoryNameBadRequest) Code() int {
	return 400
}

func (o *PutRepositoriesRawHostedByRepositoryNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/hosted/{repositoryName}][%d] putRepositoriesRawHostedByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesRawHostedByRepositoryNameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/hosted/{repositoryName}][%d] putRepositoriesRawHostedByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesRawHostedByRepositoryNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesRawHostedByRepositoryNameUnauthorized creates a PutRepositoriesRawHostedByRepositoryNameUnauthorized with default headers values
func NewPutRepositoriesRawHostedByRepositoryNameUnauthorized() *PutRepositoriesRawHostedByRepositoryNameUnauthorized {
	return &PutRepositoriesRawHostedByRepositoryNameUnauthorized{}
}

/*
PutRepositoriesRawHostedByRepositoryNameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutRepositoriesRawHostedByRepositoryNameUnauthorized struct {
}

// IsSuccess returns true when this put repositories raw hosted by repository name unauthorized response has a 2xx status code
func (o *PutRepositoriesRawHostedByRepositoryNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories raw hosted by repository name unauthorized response has a 3xx status code
func (o *PutRepositoriesRawHostedByRepositoryNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories raw hosted by repository name unauthorized response has a 4xx status code
func (o *PutRepositoriesRawHostedByRepositoryNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories raw hosted by repository name unauthorized response has a 5xx status code
func (o *PutRepositoriesRawHostedByRepositoryNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories raw hosted by repository name unauthorized response a status code equal to that given
func (o *PutRepositoriesRawHostedByRepositoryNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put repositories raw hosted by repository name unauthorized response
func (o *PutRepositoriesRawHostedByRepositoryNameUnauthorized) Code() int {
	return 401
}

func (o *PutRepositoriesRawHostedByRepositoryNameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/hosted/{repositoryName}][%d] putRepositoriesRawHostedByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesRawHostedByRepositoryNameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/hosted/{repositoryName}][%d] putRepositoriesRawHostedByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesRawHostedByRepositoryNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesRawHostedByRepositoryNameForbidden creates a PutRepositoriesRawHostedByRepositoryNameForbidden with default headers values
func NewPutRepositoriesRawHostedByRepositoryNameForbidden() *PutRepositoriesRawHostedByRepositoryNameForbidden {
	return &PutRepositoriesRawHostedByRepositoryNameForbidden{}
}

/*
PutRepositoriesRawHostedByRepositoryNameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutRepositoriesRawHostedByRepositoryNameForbidden struct {
}

// IsSuccess returns true when this put repositories raw hosted by repository name forbidden response has a 2xx status code
func (o *PutRepositoriesRawHostedByRepositoryNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories raw hosted by repository name forbidden response has a 3xx status code
func (o *PutRepositoriesRawHostedByRepositoryNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories raw hosted by repository name forbidden response has a 4xx status code
func (o *PutRepositoriesRawHostedByRepositoryNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories raw hosted by repository name forbidden response has a 5xx status code
func (o *PutRepositoriesRawHostedByRepositoryNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories raw hosted by repository name forbidden response a status code equal to that given
func (o *PutRepositoriesRawHostedByRepositoryNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put repositories raw hosted by repository name forbidden response
func (o *PutRepositoriesRawHostedByRepositoryNameForbidden) Code() int {
	return 403
}

func (o *PutRepositoriesRawHostedByRepositoryNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/hosted/{repositoryName}][%d] putRepositoriesRawHostedByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesRawHostedByRepositoryNameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/raw/hosted/{repositoryName}][%d] putRepositoriesRawHostedByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesRawHostedByRepositoryNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
