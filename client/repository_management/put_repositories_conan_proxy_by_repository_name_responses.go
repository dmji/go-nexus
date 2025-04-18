// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRepositoriesConanProxyByRepositoryNameReader is a Reader for the PutRepositoriesConanProxyByRepositoryName structure.
type PutRepositoriesConanProxyByRepositoryNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesConanProxyByRepositoryNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesConanProxyByRepositoryNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRepositoriesConanProxyByRepositoryNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRepositoriesConanProxyByRepositoryNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesConanProxyByRepositoryNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutRepositoriesConanProxyByRepositoryNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/conan/proxy/{repositoryName}] PutRepositoriesConanProxyByRepositoryName", response, response.Code())
	}
}

// NewPutRepositoriesConanProxyByRepositoryNameNoContent creates a PutRepositoriesConanProxyByRepositoryNameNoContent with default headers values
func NewPutRepositoriesConanProxyByRepositoryNameNoContent() *PutRepositoriesConanProxyByRepositoryNameNoContent {
	return &PutRepositoriesConanProxyByRepositoryNameNoContent{}
}

/*
PutRepositoriesConanProxyByRepositoryNameNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutRepositoriesConanProxyByRepositoryNameNoContent struct {
}

// IsSuccess returns true when this put repositories conan proxy by repository name no content response has a 2xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put repositories conan proxy by repository name no content response has a 3xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories conan proxy by repository name no content response has a 4xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put repositories conan proxy by repository name no content response has a 5xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories conan proxy by repository name no content response a status code equal to that given
func (o *PutRepositoriesConanProxyByRepositoryNameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put repositories conan proxy by repository name no content response
func (o *PutRepositoriesConanProxyByRepositoryNameNoContent) Code() int {
	return 204
}

func (o *PutRepositoriesConanProxyByRepositoryNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/conan/proxy/{repositoryName}][%d] putRepositoriesConanProxyByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesConanProxyByRepositoryNameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/conan/proxy/{repositoryName}][%d] putRepositoriesConanProxyByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesConanProxyByRepositoryNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesConanProxyByRepositoryNameBadRequest creates a PutRepositoriesConanProxyByRepositoryNameBadRequest with default headers values
func NewPutRepositoriesConanProxyByRepositoryNameBadRequest() *PutRepositoriesConanProxyByRepositoryNameBadRequest {
	return &PutRepositoriesConanProxyByRepositoryNameBadRequest{}
}

/*
PutRepositoriesConanProxyByRepositoryNameBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutRepositoriesConanProxyByRepositoryNameBadRequest struct {
}

// IsSuccess returns true when this put repositories conan proxy by repository name bad request response has a 2xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories conan proxy by repository name bad request response has a 3xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories conan proxy by repository name bad request response has a 4xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories conan proxy by repository name bad request response has a 5xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories conan proxy by repository name bad request response a status code equal to that given
func (o *PutRepositoriesConanProxyByRepositoryNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put repositories conan proxy by repository name bad request response
func (o *PutRepositoriesConanProxyByRepositoryNameBadRequest) Code() int {
	return 400
}

func (o *PutRepositoriesConanProxyByRepositoryNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/conan/proxy/{repositoryName}][%d] putRepositoriesConanProxyByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesConanProxyByRepositoryNameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/conan/proxy/{repositoryName}][%d] putRepositoriesConanProxyByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesConanProxyByRepositoryNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesConanProxyByRepositoryNameUnauthorized creates a PutRepositoriesConanProxyByRepositoryNameUnauthorized with default headers values
func NewPutRepositoriesConanProxyByRepositoryNameUnauthorized() *PutRepositoriesConanProxyByRepositoryNameUnauthorized {
	return &PutRepositoriesConanProxyByRepositoryNameUnauthorized{}
}

/*
PutRepositoriesConanProxyByRepositoryNameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutRepositoriesConanProxyByRepositoryNameUnauthorized struct {
}

// IsSuccess returns true when this put repositories conan proxy by repository name unauthorized response has a 2xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories conan proxy by repository name unauthorized response has a 3xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories conan proxy by repository name unauthorized response has a 4xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories conan proxy by repository name unauthorized response has a 5xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories conan proxy by repository name unauthorized response a status code equal to that given
func (o *PutRepositoriesConanProxyByRepositoryNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put repositories conan proxy by repository name unauthorized response
func (o *PutRepositoriesConanProxyByRepositoryNameUnauthorized) Code() int {
	return 401
}

func (o *PutRepositoriesConanProxyByRepositoryNameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/conan/proxy/{repositoryName}][%d] putRepositoriesConanProxyByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesConanProxyByRepositoryNameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/conan/proxy/{repositoryName}][%d] putRepositoriesConanProxyByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesConanProxyByRepositoryNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesConanProxyByRepositoryNameForbidden creates a PutRepositoriesConanProxyByRepositoryNameForbidden with default headers values
func NewPutRepositoriesConanProxyByRepositoryNameForbidden() *PutRepositoriesConanProxyByRepositoryNameForbidden {
	return &PutRepositoriesConanProxyByRepositoryNameForbidden{}
}

/*
PutRepositoriesConanProxyByRepositoryNameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutRepositoriesConanProxyByRepositoryNameForbidden struct {
}

// IsSuccess returns true when this put repositories conan proxy by repository name forbidden response has a 2xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories conan proxy by repository name forbidden response has a 3xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories conan proxy by repository name forbidden response has a 4xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories conan proxy by repository name forbidden response has a 5xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories conan proxy by repository name forbidden response a status code equal to that given
func (o *PutRepositoriesConanProxyByRepositoryNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put repositories conan proxy by repository name forbidden response
func (o *PutRepositoriesConanProxyByRepositoryNameForbidden) Code() int {
	return 403
}

func (o *PutRepositoriesConanProxyByRepositoryNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/conan/proxy/{repositoryName}][%d] putRepositoriesConanProxyByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesConanProxyByRepositoryNameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/conan/proxy/{repositoryName}][%d] putRepositoriesConanProxyByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesConanProxyByRepositoryNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesConanProxyByRepositoryNameNotFound creates a PutRepositoriesConanProxyByRepositoryNameNotFound with default headers values
func NewPutRepositoriesConanProxyByRepositoryNameNotFound() *PutRepositoriesConanProxyByRepositoryNameNotFound {
	return &PutRepositoriesConanProxyByRepositoryNameNotFound{}
}

/*
PutRepositoriesConanProxyByRepositoryNameNotFound describes a response with status code 404, with default header values.

Repository not found
*/
type PutRepositoriesConanProxyByRepositoryNameNotFound struct {
}

// IsSuccess returns true when this put repositories conan proxy by repository name not found response has a 2xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories conan proxy by repository name not found response has a 3xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories conan proxy by repository name not found response has a 4xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories conan proxy by repository name not found response has a 5xx status code
func (o *PutRepositoriesConanProxyByRepositoryNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories conan proxy by repository name not found response a status code equal to that given
func (o *PutRepositoriesConanProxyByRepositoryNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put repositories conan proxy by repository name not found response
func (o *PutRepositoriesConanProxyByRepositoryNameNotFound) Code() int {
	return 404
}

func (o *PutRepositoriesConanProxyByRepositoryNameNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/conan/proxy/{repositoryName}][%d] putRepositoriesConanProxyByRepositoryNameNotFound", 404)
}

func (o *PutRepositoriesConanProxyByRepositoryNameNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/conan/proxy/{repositoryName}][%d] putRepositoriesConanProxyByRepositoryNameNotFound", 404)
}

func (o *PutRepositoriesConanProxyByRepositoryNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
