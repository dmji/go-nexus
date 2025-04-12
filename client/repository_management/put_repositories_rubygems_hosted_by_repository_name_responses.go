// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRepositoriesRubygemsHostedByRepositoryNameReader is a Reader for the PutRepositoriesRubygemsHostedByRepositoryName structure.
type PutRepositoriesRubygemsHostedByRepositoryNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesRubygemsHostedByRepositoryNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesRubygemsHostedByRepositoryNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRepositoriesRubygemsHostedByRepositoryNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRepositoriesRubygemsHostedByRepositoryNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesRubygemsHostedByRepositoryNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/rubygems/hosted/{repositoryName}] PutRepositoriesRubygemsHostedByRepositoryName", response, response.Code())
	}
}

// NewPutRepositoriesRubygemsHostedByRepositoryNameNoContent creates a PutRepositoriesRubygemsHostedByRepositoryNameNoContent with default headers values
func NewPutRepositoriesRubygemsHostedByRepositoryNameNoContent() *PutRepositoriesRubygemsHostedByRepositoryNameNoContent {
	return &PutRepositoriesRubygemsHostedByRepositoryNameNoContent{}
}

/*
PutRepositoriesRubygemsHostedByRepositoryNameNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutRepositoriesRubygemsHostedByRepositoryNameNoContent struct {
}

// IsSuccess returns true when this put repositories rubygems hosted by repository name no content response has a 2xx status code
func (o *PutRepositoriesRubygemsHostedByRepositoryNameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put repositories rubygems hosted by repository name no content response has a 3xx status code
func (o *PutRepositoriesRubygemsHostedByRepositoryNameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories rubygems hosted by repository name no content response has a 4xx status code
func (o *PutRepositoriesRubygemsHostedByRepositoryNameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put repositories rubygems hosted by repository name no content response has a 5xx status code
func (o *PutRepositoriesRubygemsHostedByRepositoryNameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories rubygems hosted by repository name no content response a status code equal to that given
func (o *PutRepositoriesRubygemsHostedByRepositoryNameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put repositories rubygems hosted by repository name no content response
func (o *PutRepositoriesRubygemsHostedByRepositoryNameNoContent) Code() int {
	return 204
}

func (o *PutRepositoriesRubygemsHostedByRepositoryNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/rubygems/hosted/{repositoryName}][%d] putRepositoriesRubygemsHostedByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesRubygemsHostedByRepositoryNameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/rubygems/hosted/{repositoryName}][%d] putRepositoriesRubygemsHostedByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesRubygemsHostedByRepositoryNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesRubygemsHostedByRepositoryNameBadRequest creates a PutRepositoriesRubygemsHostedByRepositoryNameBadRequest with default headers values
func NewPutRepositoriesRubygemsHostedByRepositoryNameBadRequest() *PutRepositoriesRubygemsHostedByRepositoryNameBadRequest {
	return &PutRepositoriesRubygemsHostedByRepositoryNameBadRequest{}
}

/*
PutRepositoriesRubygemsHostedByRepositoryNameBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutRepositoriesRubygemsHostedByRepositoryNameBadRequest struct {
}

// IsSuccess returns true when this put repositories rubygems hosted by repository name bad request response has a 2xx status code
func (o *PutRepositoriesRubygemsHostedByRepositoryNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories rubygems hosted by repository name bad request response has a 3xx status code
func (o *PutRepositoriesRubygemsHostedByRepositoryNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories rubygems hosted by repository name bad request response has a 4xx status code
func (o *PutRepositoriesRubygemsHostedByRepositoryNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories rubygems hosted by repository name bad request response has a 5xx status code
func (o *PutRepositoriesRubygemsHostedByRepositoryNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories rubygems hosted by repository name bad request response a status code equal to that given
func (o *PutRepositoriesRubygemsHostedByRepositoryNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put repositories rubygems hosted by repository name bad request response
func (o *PutRepositoriesRubygemsHostedByRepositoryNameBadRequest) Code() int {
	return 400
}

func (o *PutRepositoriesRubygemsHostedByRepositoryNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/rubygems/hosted/{repositoryName}][%d] putRepositoriesRubygemsHostedByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesRubygemsHostedByRepositoryNameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/rubygems/hosted/{repositoryName}][%d] putRepositoriesRubygemsHostedByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesRubygemsHostedByRepositoryNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesRubygemsHostedByRepositoryNameUnauthorized creates a PutRepositoriesRubygemsHostedByRepositoryNameUnauthorized with default headers values
func NewPutRepositoriesRubygemsHostedByRepositoryNameUnauthorized() *PutRepositoriesRubygemsHostedByRepositoryNameUnauthorized {
	return &PutRepositoriesRubygemsHostedByRepositoryNameUnauthorized{}
}

/*
PutRepositoriesRubygemsHostedByRepositoryNameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutRepositoriesRubygemsHostedByRepositoryNameUnauthorized struct {
}

// IsSuccess returns true when this put repositories rubygems hosted by repository name unauthorized response has a 2xx status code
func (o *PutRepositoriesRubygemsHostedByRepositoryNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories rubygems hosted by repository name unauthorized response has a 3xx status code
func (o *PutRepositoriesRubygemsHostedByRepositoryNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories rubygems hosted by repository name unauthorized response has a 4xx status code
func (o *PutRepositoriesRubygemsHostedByRepositoryNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories rubygems hosted by repository name unauthorized response has a 5xx status code
func (o *PutRepositoriesRubygemsHostedByRepositoryNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories rubygems hosted by repository name unauthorized response a status code equal to that given
func (o *PutRepositoriesRubygemsHostedByRepositoryNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put repositories rubygems hosted by repository name unauthorized response
func (o *PutRepositoriesRubygemsHostedByRepositoryNameUnauthorized) Code() int {
	return 401
}

func (o *PutRepositoriesRubygemsHostedByRepositoryNameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/rubygems/hosted/{repositoryName}][%d] putRepositoriesRubygemsHostedByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesRubygemsHostedByRepositoryNameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/rubygems/hosted/{repositoryName}][%d] putRepositoriesRubygemsHostedByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesRubygemsHostedByRepositoryNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesRubygemsHostedByRepositoryNameForbidden creates a PutRepositoriesRubygemsHostedByRepositoryNameForbidden with default headers values
func NewPutRepositoriesRubygemsHostedByRepositoryNameForbidden() *PutRepositoriesRubygemsHostedByRepositoryNameForbidden {
	return &PutRepositoriesRubygemsHostedByRepositoryNameForbidden{}
}

/*
PutRepositoriesRubygemsHostedByRepositoryNameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutRepositoriesRubygemsHostedByRepositoryNameForbidden struct {
}

// IsSuccess returns true when this put repositories rubygems hosted by repository name forbidden response has a 2xx status code
func (o *PutRepositoriesRubygemsHostedByRepositoryNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories rubygems hosted by repository name forbidden response has a 3xx status code
func (o *PutRepositoriesRubygemsHostedByRepositoryNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories rubygems hosted by repository name forbidden response has a 4xx status code
func (o *PutRepositoriesRubygemsHostedByRepositoryNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories rubygems hosted by repository name forbidden response has a 5xx status code
func (o *PutRepositoriesRubygemsHostedByRepositoryNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories rubygems hosted by repository name forbidden response a status code equal to that given
func (o *PutRepositoriesRubygemsHostedByRepositoryNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put repositories rubygems hosted by repository name forbidden response
func (o *PutRepositoriesRubygemsHostedByRepositoryNameForbidden) Code() int {
	return 403
}

func (o *PutRepositoriesRubygemsHostedByRepositoryNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/rubygems/hosted/{repositoryName}][%d] putRepositoriesRubygemsHostedByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesRubygemsHostedByRepositoryNameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/rubygems/hosted/{repositoryName}][%d] putRepositoriesRubygemsHostedByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesRubygemsHostedByRepositoryNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
