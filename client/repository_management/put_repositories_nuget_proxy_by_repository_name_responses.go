// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRepositoriesNugetProxyByRepositoryNameReader is a Reader for the PutRepositoriesNugetProxyByRepositoryName structure.
type PutRepositoriesNugetProxyByRepositoryNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesNugetProxyByRepositoryNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesNugetProxyByRepositoryNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRepositoriesNugetProxyByRepositoryNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRepositoriesNugetProxyByRepositoryNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesNugetProxyByRepositoryNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/nuget/proxy/{repositoryName}] PutRepositoriesNugetProxyByRepositoryName", response, response.Code())
	}
}

// NewPutRepositoriesNugetProxyByRepositoryNameNoContent creates a PutRepositoriesNugetProxyByRepositoryNameNoContent with default headers values
func NewPutRepositoriesNugetProxyByRepositoryNameNoContent() *PutRepositoriesNugetProxyByRepositoryNameNoContent {
	return &PutRepositoriesNugetProxyByRepositoryNameNoContent{}
}

/*
PutRepositoriesNugetProxyByRepositoryNameNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutRepositoriesNugetProxyByRepositoryNameNoContent struct {
}

// IsSuccess returns true when this put repositories nuget proxy by repository name no content response has a 2xx status code
func (o *PutRepositoriesNugetProxyByRepositoryNameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put repositories nuget proxy by repository name no content response has a 3xx status code
func (o *PutRepositoriesNugetProxyByRepositoryNameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories nuget proxy by repository name no content response has a 4xx status code
func (o *PutRepositoriesNugetProxyByRepositoryNameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put repositories nuget proxy by repository name no content response has a 5xx status code
func (o *PutRepositoriesNugetProxyByRepositoryNameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories nuget proxy by repository name no content response a status code equal to that given
func (o *PutRepositoriesNugetProxyByRepositoryNameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put repositories nuget proxy by repository name no content response
func (o *PutRepositoriesNugetProxyByRepositoryNameNoContent) Code() int {
	return 204
}

func (o *PutRepositoriesNugetProxyByRepositoryNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/proxy/{repositoryName}][%d] putRepositoriesNugetProxyByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesNugetProxyByRepositoryNameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/proxy/{repositoryName}][%d] putRepositoriesNugetProxyByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesNugetProxyByRepositoryNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesNugetProxyByRepositoryNameBadRequest creates a PutRepositoriesNugetProxyByRepositoryNameBadRequest with default headers values
func NewPutRepositoriesNugetProxyByRepositoryNameBadRequest() *PutRepositoriesNugetProxyByRepositoryNameBadRequest {
	return &PutRepositoriesNugetProxyByRepositoryNameBadRequest{}
}

/*
PutRepositoriesNugetProxyByRepositoryNameBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutRepositoriesNugetProxyByRepositoryNameBadRequest struct {
}

// IsSuccess returns true when this put repositories nuget proxy by repository name bad request response has a 2xx status code
func (o *PutRepositoriesNugetProxyByRepositoryNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories nuget proxy by repository name bad request response has a 3xx status code
func (o *PutRepositoriesNugetProxyByRepositoryNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories nuget proxy by repository name bad request response has a 4xx status code
func (o *PutRepositoriesNugetProxyByRepositoryNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories nuget proxy by repository name bad request response has a 5xx status code
func (o *PutRepositoriesNugetProxyByRepositoryNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories nuget proxy by repository name bad request response a status code equal to that given
func (o *PutRepositoriesNugetProxyByRepositoryNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put repositories nuget proxy by repository name bad request response
func (o *PutRepositoriesNugetProxyByRepositoryNameBadRequest) Code() int {
	return 400
}

func (o *PutRepositoriesNugetProxyByRepositoryNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/proxy/{repositoryName}][%d] putRepositoriesNugetProxyByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesNugetProxyByRepositoryNameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/proxy/{repositoryName}][%d] putRepositoriesNugetProxyByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesNugetProxyByRepositoryNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesNugetProxyByRepositoryNameUnauthorized creates a PutRepositoriesNugetProxyByRepositoryNameUnauthorized with default headers values
func NewPutRepositoriesNugetProxyByRepositoryNameUnauthorized() *PutRepositoriesNugetProxyByRepositoryNameUnauthorized {
	return &PutRepositoriesNugetProxyByRepositoryNameUnauthorized{}
}

/*
PutRepositoriesNugetProxyByRepositoryNameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutRepositoriesNugetProxyByRepositoryNameUnauthorized struct {
}

// IsSuccess returns true when this put repositories nuget proxy by repository name unauthorized response has a 2xx status code
func (o *PutRepositoriesNugetProxyByRepositoryNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories nuget proxy by repository name unauthorized response has a 3xx status code
func (o *PutRepositoriesNugetProxyByRepositoryNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories nuget proxy by repository name unauthorized response has a 4xx status code
func (o *PutRepositoriesNugetProxyByRepositoryNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories nuget proxy by repository name unauthorized response has a 5xx status code
func (o *PutRepositoriesNugetProxyByRepositoryNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories nuget proxy by repository name unauthorized response a status code equal to that given
func (o *PutRepositoriesNugetProxyByRepositoryNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put repositories nuget proxy by repository name unauthorized response
func (o *PutRepositoriesNugetProxyByRepositoryNameUnauthorized) Code() int {
	return 401
}

func (o *PutRepositoriesNugetProxyByRepositoryNameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/proxy/{repositoryName}][%d] putRepositoriesNugetProxyByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesNugetProxyByRepositoryNameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/proxy/{repositoryName}][%d] putRepositoriesNugetProxyByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesNugetProxyByRepositoryNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesNugetProxyByRepositoryNameForbidden creates a PutRepositoriesNugetProxyByRepositoryNameForbidden with default headers values
func NewPutRepositoriesNugetProxyByRepositoryNameForbidden() *PutRepositoriesNugetProxyByRepositoryNameForbidden {
	return &PutRepositoriesNugetProxyByRepositoryNameForbidden{}
}

/*
PutRepositoriesNugetProxyByRepositoryNameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutRepositoriesNugetProxyByRepositoryNameForbidden struct {
}

// IsSuccess returns true when this put repositories nuget proxy by repository name forbidden response has a 2xx status code
func (o *PutRepositoriesNugetProxyByRepositoryNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories nuget proxy by repository name forbidden response has a 3xx status code
func (o *PutRepositoriesNugetProxyByRepositoryNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories nuget proxy by repository name forbidden response has a 4xx status code
func (o *PutRepositoriesNugetProxyByRepositoryNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories nuget proxy by repository name forbidden response has a 5xx status code
func (o *PutRepositoriesNugetProxyByRepositoryNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories nuget proxy by repository name forbidden response a status code equal to that given
func (o *PutRepositoriesNugetProxyByRepositoryNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put repositories nuget proxy by repository name forbidden response
func (o *PutRepositoriesNugetProxyByRepositoryNameForbidden) Code() int {
	return 403
}

func (o *PutRepositoriesNugetProxyByRepositoryNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/proxy/{repositoryName}][%d] putRepositoriesNugetProxyByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesNugetProxyByRepositoryNameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/proxy/{repositoryName}][%d] putRepositoriesNugetProxyByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesNugetProxyByRepositoryNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
