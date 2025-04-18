// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRepositoriesNugetGroupByRepositoryNameReader is a Reader for the PutRepositoriesNugetGroupByRepositoryName structure.
type PutRepositoriesNugetGroupByRepositoryNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesNugetGroupByRepositoryNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesNugetGroupByRepositoryNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRepositoriesNugetGroupByRepositoryNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRepositoriesNugetGroupByRepositoryNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesNugetGroupByRepositoryNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/nuget/group/{repositoryName}] PutRepositoriesNugetGroupByRepositoryName", response, response.Code())
	}
}

// NewPutRepositoriesNugetGroupByRepositoryNameNoContent creates a PutRepositoriesNugetGroupByRepositoryNameNoContent with default headers values
func NewPutRepositoriesNugetGroupByRepositoryNameNoContent() *PutRepositoriesNugetGroupByRepositoryNameNoContent {
	return &PutRepositoriesNugetGroupByRepositoryNameNoContent{}
}

/*
PutRepositoriesNugetGroupByRepositoryNameNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutRepositoriesNugetGroupByRepositoryNameNoContent struct {
}

// IsSuccess returns true when this put repositories nuget group by repository name no content response has a 2xx status code
func (o *PutRepositoriesNugetGroupByRepositoryNameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put repositories nuget group by repository name no content response has a 3xx status code
func (o *PutRepositoriesNugetGroupByRepositoryNameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories nuget group by repository name no content response has a 4xx status code
func (o *PutRepositoriesNugetGroupByRepositoryNameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put repositories nuget group by repository name no content response has a 5xx status code
func (o *PutRepositoriesNugetGroupByRepositoryNameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories nuget group by repository name no content response a status code equal to that given
func (o *PutRepositoriesNugetGroupByRepositoryNameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put repositories nuget group by repository name no content response
func (o *PutRepositoriesNugetGroupByRepositoryNameNoContent) Code() int {
	return 204
}

func (o *PutRepositoriesNugetGroupByRepositoryNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/group/{repositoryName}][%d] putRepositoriesNugetGroupByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesNugetGroupByRepositoryNameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/group/{repositoryName}][%d] putRepositoriesNugetGroupByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesNugetGroupByRepositoryNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesNugetGroupByRepositoryNameBadRequest creates a PutRepositoriesNugetGroupByRepositoryNameBadRequest with default headers values
func NewPutRepositoriesNugetGroupByRepositoryNameBadRequest() *PutRepositoriesNugetGroupByRepositoryNameBadRequest {
	return &PutRepositoriesNugetGroupByRepositoryNameBadRequest{}
}

/*
PutRepositoriesNugetGroupByRepositoryNameBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutRepositoriesNugetGroupByRepositoryNameBadRequest struct {
}

// IsSuccess returns true when this put repositories nuget group by repository name bad request response has a 2xx status code
func (o *PutRepositoriesNugetGroupByRepositoryNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories nuget group by repository name bad request response has a 3xx status code
func (o *PutRepositoriesNugetGroupByRepositoryNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories nuget group by repository name bad request response has a 4xx status code
func (o *PutRepositoriesNugetGroupByRepositoryNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories nuget group by repository name bad request response has a 5xx status code
func (o *PutRepositoriesNugetGroupByRepositoryNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories nuget group by repository name bad request response a status code equal to that given
func (o *PutRepositoriesNugetGroupByRepositoryNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put repositories nuget group by repository name bad request response
func (o *PutRepositoriesNugetGroupByRepositoryNameBadRequest) Code() int {
	return 400
}

func (o *PutRepositoriesNugetGroupByRepositoryNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/group/{repositoryName}][%d] putRepositoriesNugetGroupByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesNugetGroupByRepositoryNameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/group/{repositoryName}][%d] putRepositoriesNugetGroupByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesNugetGroupByRepositoryNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesNugetGroupByRepositoryNameUnauthorized creates a PutRepositoriesNugetGroupByRepositoryNameUnauthorized with default headers values
func NewPutRepositoriesNugetGroupByRepositoryNameUnauthorized() *PutRepositoriesNugetGroupByRepositoryNameUnauthorized {
	return &PutRepositoriesNugetGroupByRepositoryNameUnauthorized{}
}

/*
PutRepositoriesNugetGroupByRepositoryNameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutRepositoriesNugetGroupByRepositoryNameUnauthorized struct {
}

// IsSuccess returns true when this put repositories nuget group by repository name unauthorized response has a 2xx status code
func (o *PutRepositoriesNugetGroupByRepositoryNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories nuget group by repository name unauthorized response has a 3xx status code
func (o *PutRepositoriesNugetGroupByRepositoryNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories nuget group by repository name unauthorized response has a 4xx status code
func (o *PutRepositoriesNugetGroupByRepositoryNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories nuget group by repository name unauthorized response has a 5xx status code
func (o *PutRepositoriesNugetGroupByRepositoryNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories nuget group by repository name unauthorized response a status code equal to that given
func (o *PutRepositoriesNugetGroupByRepositoryNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put repositories nuget group by repository name unauthorized response
func (o *PutRepositoriesNugetGroupByRepositoryNameUnauthorized) Code() int {
	return 401
}

func (o *PutRepositoriesNugetGroupByRepositoryNameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/group/{repositoryName}][%d] putRepositoriesNugetGroupByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesNugetGroupByRepositoryNameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/group/{repositoryName}][%d] putRepositoriesNugetGroupByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesNugetGroupByRepositoryNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesNugetGroupByRepositoryNameForbidden creates a PutRepositoriesNugetGroupByRepositoryNameForbidden with default headers values
func NewPutRepositoriesNugetGroupByRepositoryNameForbidden() *PutRepositoriesNugetGroupByRepositoryNameForbidden {
	return &PutRepositoriesNugetGroupByRepositoryNameForbidden{}
}

/*
PutRepositoriesNugetGroupByRepositoryNameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutRepositoriesNugetGroupByRepositoryNameForbidden struct {
}

// IsSuccess returns true when this put repositories nuget group by repository name forbidden response has a 2xx status code
func (o *PutRepositoriesNugetGroupByRepositoryNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories nuget group by repository name forbidden response has a 3xx status code
func (o *PutRepositoriesNugetGroupByRepositoryNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories nuget group by repository name forbidden response has a 4xx status code
func (o *PutRepositoriesNugetGroupByRepositoryNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories nuget group by repository name forbidden response has a 5xx status code
func (o *PutRepositoriesNugetGroupByRepositoryNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories nuget group by repository name forbidden response a status code equal to that given
func (o *PutRepositoriesNugetGroupByRepositoryNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put repositories nuget group by repository name forbidden response
func (o *PutRepositoriesNugetGroupByRepositoryNameForbidden) Code() int {
	return 403
}

func (o *PutRepositoriesNugetGroupByRepositoryNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/group/{repositoryName}][%d] putRepositoriesNugetGroupByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesNugetGroupByRepositoryNameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/group/{repositoryName}][%d] putRepositoriesNugetGroupByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesNugetGroupByRepositoryNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
