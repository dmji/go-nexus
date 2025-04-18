// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRepositoriesPypiHostedByRepositoryNameReader is a Reader for the PutRepositoriesPypiHostedByRepositoryName structure.
type PutRepositoriesPypiHostedByRepositoryNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesPypiHostedByRepositoryNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesPypiHostedByRepositoryNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRepositoriesPypiHostedByRepositoryNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRepositoriesPypiHostedByRepositoryNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesPypiHostedByRepositoryNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/pypi/hosted/{repositoryName}] PutRepositoriesPypiHostedByRepositoryName", response, response.Code())
	}
}

// NewPutRepositoriesPypiHostedByRepositoryNameNoContent creates a PutRepositoriesPypiHostedByRepositoryNameNoContent with default headers values
func NewPutRepositoriesPypiHostedByRepositoryNameNoContent() *PutRepositoriesPypiHostedByRepositoryNameNoContent {
	return &PutRepositoriesPypiHostedByRepositoryNameNoContent{}
}

/*
PutRepositoriesPypiHostedByRepositoryNameNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutRepositoriesPypiHostedByRepositoryNameNoContent struct {
}

// IsSuccess returns true when this put repositories pypi hosted by repository name no content response has a 2xx status code
func (o *PutRepositoriesPypiHostedByRepositoryNameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put repositories pypi hosted by repository name no content response has a 3xx status code
func (o *PutRepositoriesPypiHostedByRepositoryNameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories pypi hosted by repository name no content response has a 4xx status code
func (o *PutRepositoriesPypiHostedByRepositoryNameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put repositories pypi hosted by repository name no content response has a 5xx status code
func (o *PutRepositoriesPypiHostedByRepositoryNameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories pypi hosted by repository name no content response a status code equal to that given
func (o *PutRepositoriesPypiHostedByRepositoryNameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put repositories pypi hosted by repository name no content response
func (o *PutRepositoriesPypiHostedByRepositoryNameNoContent) Code() int {
	return 204
}

func (o *PutRepositoriesPypiHostedByRepositoryNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/hosted/{repositoryName}][%d] putRepositoriesPypiHostedByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesPypiHostedByRepositoryNameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/hosted/{repositoryName}][%d] putRepositoriesPypiHostedByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesPypiHostedByRepositoryNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesPypiHostedByRepositoryNameBadRequest creates a PutRepositoriesPypiHostedByRepositoryNameBadRequest with default headers values
func NewPutRepositoriesPypiHostedByRepositoryNameBadRequest() *PutRepositoriesPypiHostedByRepositoryNameBadRequest {
	return &PutRepositoriesPypiHostedByRepositoryNameBadRequest{}
}

/*
PutRepositoriesPypiHostedByRepositoryNameBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutRepositoriesPypiHostedByRepositoryNameBadRequest struct {
}

// IsSuccess returns true when this put repositories pypi hosted by repository name bad request response has a 2xx status code
func (o *PutRepositoriesPypiHostedByRepositoryNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories pypi hosted by repository name bad request response has a 3xx status code
func (o *PutRepositoriesPypiHostedByRepositoryNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories pypi hosted by repository name bad request response has a 4xx status code
func (o *PutRepositoriesPypiHostedByRepositoryNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories pypi hosted by repository name bad request response has a 5xx status code
func (o *PutRepositoriesPypiHostedByRepositoryNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories pypi hosted by repository name bad request response a status code equal to that given
func (o *PutRepositoriesPypiHostedByRepositoryNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put repositories pypi hosted by repository name bad request response
func (o *PutRepositoriesPypiHostedByRepositoryNameBadRequest) Code() int {
	return 400
}

func (o *PutRepositoriesPypiHostedByRepositoryNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/hosted/{repositoryName}][%d] putRepositoriesPypiHostedByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesPypiHostedByRepositoryNameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/hosted/{repositoryName}][%d] putRepositoriesPypiHostedByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesPypiHostedByRepositoryNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesPypiHostedByRepositoryNameUnauthorized creates a PutRepositoriesPypiHostedByRepositoryNameUnauthorized with default headers values
func NewPutRepositoriesPypiHostedByRepositoryNameUnauthorized() *PutRepositoriesPypiHostedByRepositoryNameUnauthorized {
	return &PutRepositoriesPypiHostedByRepositoryNameUnauthorized{}
}

/*
PutRepositoriesPypiHostedByRepositoryNameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutRepositoriesPypiHostedByRepositoryNameUnauthorized struct {
}

// IsSuccess returns true when this put repositories pypi hosted by repository name unauthorized response has a 2xx status code
func (o *PutRepositoriesPypiHostedByRepositoryNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories pypi hosted by repository name unauthorized response has a 3xx status code
func (o *PutRepositoriesPypiHostedByRepositoryNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories pypi hosted by repository name unauthorized response has a 4xx status code
func (o *PutRepositoriesPypiHostedByRepositoryNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories pypi hosted by repository name unauthorized response has a 5xx status code
func (o *PutRepositoriesPypiHostedByRepositoryNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories pypi hosted by repository name unauthorized response a status code equal to that given
func (o *PutRepositoriesPypiHostedByRepositoryNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put repositories pypi hosted by repository name unauthorized response
func (o *PutRepositoriesPypiHostedByRepositoryNameUnauthorized) Code() int {
	return 401
}

func (o *PutRepositoriesPypiHostedByRepositoryNameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/hosted/{repositoryName}][%d] putRepositoriesPypiHostedByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesPypiHostedByRepositoryNameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/hosted/{repositoryName}][%d] putRepositoriesPypiHostedByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesPypiHostedByRepositoryNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesPypiHostedByRepositoryNameForbidden creates a PutRepositoriesPypiHostedByRepositoryNameForbidden with default headers values
func NewPutRepositoriesPypiHostedByRepositoryNameForbidden() *PutRepositoriesPypiHostedByRepositoryNameForbidden {
	return &PutRepositoriesPypiHostedByRepositoryNameForbidden{}
}

/*
PutRepositoriesPypiHostedByRepositoryNameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutRepositoriesPypiHostedByRepositoryNameForbidden struct {
}

// IsSuccess returns true when this put repositories pypi hosted by repository name forbidden response has a 2xx status code
func (o *PutRepositoriesPypiHostedByRepositoryNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories pypi hosted by repository name forbidden response has a 3xx status code
func (o *PutRepositoriesPypiHostedByRepositoryNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories pypi hosted by repository name forbidden response has a 4xx status code
func (o *PutRepositoriesPypiHostedByRepositoryNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories pypi hosted by repository name forbidden response has a 5xx status code
func (o *PutRepositoriesPypiHostedByRepositoryNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories pypi hosted by repository name forbidden response a status code equal to that given
func (o *PutRepositoriesPypiHostedByRepositoryNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put repositories pypi hosted by repository name forbidden response
func (o *PutRepositoriesPypiHostedByRepositoryNameForbidden) Code() int {
	return 403
}

func (o *PutRepositoriesPypiHostedByRepositoryNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/hosted/{repositoryName}][%d] putRepositoriesPypiHostedByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesPypiHostedByRepositoryNameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/hosted/{repositoryName}][%d] putRepositoriesPypiHostedByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesPypiHostedByRepositoryNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
