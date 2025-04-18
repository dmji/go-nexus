// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRepositoriesPypiGroupByRepositoryNameReader is a Reader for the PutRepositoriesPypiGroupByRepositoryName structure.
type PutRepositoriesPypiGroupByRepositoryNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesPypiGroupByRepositoryNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesPypiGroupByRepositoryNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRepositoriesPypiGroupByRepositoryNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRepositoriesPypiGroupByRepositoryNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesPypiGroupByRepositoryNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/pypi/group/{repositoryName}] PutRepositoriesPypiGroupByRepositoryName", response, response.Code())
	}
}

// NewPutRepositoriesPypiGroupByRepositoryNameNoContent creates a PutRepositoriesPypiGroupByRepositoryNameNoContent with default headers values
func NewPutRepositoriesPypiGroupByRepositoryNameNoContent() *PutRepositoriesPypiGroupByRepositoryNameNoContent {
	return &PutRepositoriesPypiGroupByRepositoryNameNoContent{}
}

/*
PutRepositoriesPypiGroupByRepositoryNameNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutRepositoriesPypiGroupByRepositoryNameNoContent struct {
}

// IsSuccess returns true when this put repositories pypi group by repository name no content response has a 2xx status code
func (o *PutRepositoriesPypiGroupByRepositoryNameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put repositories pypi group by repository name no content response has a 3xx status code
func (o *PutRepositoriesPypiGroupByRepositoryNameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories pypi group by repository name no content response has a 4xx status code
func (o *PutRepositoriesPypiGroupByRepositoryNameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put repositories pypi group by repository name no content response has a 5xx status code
func (o *PutRepositoriesPypiGroupByRepositoryNameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories pypi group by repository name no content response a status code equal to that given
func (o *PutRepositoriesPypiGroupByRepositoryNameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put repositories pypi group by repository name no content response
func (o *PutRepositoriesPypiGroupByRepositoryNameNoContent) Code() int {
	return 204
}

func (o *PutRepositoriesPypiGroupByRepositoryNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/group/{repositoryName}][%d] putRepositoriesPypiGroupByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesPypiGroupByRepositoryNameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/group/{repositoryName}][%d] putRepositoriesPypiGroupByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesPypiGroupByRepositoryNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesPypiGroupByRepositoryNameBadRequest creates a PutRepositoriesPypiGroupByRepositoryNameBadRequest with default headers values
func NewPutRepositoriesPypiGroupByRepositoryNameBadRequest() *PutRepositoriesPypiGroupByRepositoryNameBadRequest {
	return &PutRepositoriesPypiGroupByRepositoryNameBadRequest{}
}

/*
PutRepositoriesPypiGroupByRepositoryNameBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutRepositoriesPypiGroupByRepositoryNameBadRequest struct {
}

// IsSuccess returns true when this put repositories pypi group by repository name bad request response has a 2xx status code
func (o *PutRepositoriesPypiGroupByRepositoryNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories pypi group by repository name bad request response has a 3xx status code
func (o *PutRepositoriesPypiGroupByRepositoryNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories pypi group by repository name bad request response has a 4xx status code
func (o *PutRepositoriesPypiGroupByRepositoryNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories pypi group by repository name bad request response has a 5xx status code
func (o *PutRepositoriesPypiGroupByRepositoryNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories pypi group by repository name bad request response a status code equal to that given
func (o *PutRepositoriesPypiGroupByRepositoryNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put repositories pypi group by repository name bad request response
func (o *PutRepositoriesPypiGroupByRepositoryNameBadRequest) Code() int {
	return 400
}

func (o *PutRepositoriesPypiGroupByRepositoryNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/group/{repositoryName}][%d] putRepositoriesPypiGroupByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesPypiGroupByRepositoryNameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/group/{repositoryName}][%d] putRepositoriesPypiGroupByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesPypiGroupByRepositoryNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesPypiGroupByRepositoryNameUnauthorized creates a PutRepositoriesPypiGroupByRepositoryNameUnauthorized with default headers values
func NewPutRepositoriesPypiGroupByRepositoryNameUnauthorized() *PutRepositoriesPypiGroupByRepositoryNameUnauthorized {
	return &PutRepositoriesPypiGroupByRepositoryNameUnauthorized{}
}

/*
PutRepositoriesPypiGroupByRepositoryNameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutRepositoriesPypiGroupByRepositoryNameUnauthorized struct {
}

// IsSuccess returns true when this put repositories pypi group by repository name unauthorized response has a 2xx status code
func (o *PutRepositoriesPypiGroupByRepositoryNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories pypi group by repository name unauthorized response has a 3xx status code
func (o *PutRepositoriesPypiGroupByRepositoryNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories pypi group by repository name unauthorized response has a 4xx status code
func (o *PutRepositoriesPypiGroupByRepositoryNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories pypi group by repository name unauthorized response has a 5xx status code
func (o *PutRepositoriesPypiGroupByRepositoryNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories pypi group by repository name unauthorized response a status code equal to that given
func (o *PutRepositoriesPypiGroupByRepositoryNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put repositories pypi group by repository name unauthorized response
func (o *PutRepositoriesPypiGroupByRepositoryNameUnauthorized) Code() int {
	return 401
}

func (o *PutRepositoriesPypiGroupByRepositoryNameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/group/{repositoryName}][%d] putRepositoriesPypiGroupByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesPypiGroupByRepositoryNameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/group/{repositoryName}][%d] putRepositoriesPypiGroupByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesPypiGroupByRepositoryNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesPypiGroupByRepositoryNameForbidden creates a PutRepositoriesPypiGroupByRepositoryNameForbidden with default headers values
func NewPutRepositoriesPypiGroupByRepositoryNameForbidden() *PutRepositoriesPypiGroupByRepositoryNameForbidden {
	return &PutRepositoriesPypiGroupByRepositoryNameForbidden{}
}

/*
PutRepositoriesPypiGroupByRepositoryNameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutRepositoriesPypiGroupByRepositoryNameForbidden struct {
}

// IsSuccess returns true when this put repositories pypi group by repository name forbidden response has a 2xx status code
func (o *PutRepositoriesPypiGroupByRepositoryNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories pypi group by repository name forbidden response has a 3xx status code
func (o *PutRepositoriesPypiGroupByRepositoryNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories pypi group by repository name forbidden response has a 4xx status code
func (o *PutRepositoriesPypiGroupByRepositoryNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories pypi group by repository name forbidden response has a 5xx status code
func (o *PutRepositoriesPypiGroupByRepositoryNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories pypi group by repository name forbidden response a status code equal to that given
func (o *PutRepositoriesPypiGroupByRepositoryNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put repositories pypi group by repository name forbidden response
func (o *PutRepositoriesPypiGroupByRepositoryNameForbidden) Code() int {
	return 403
}

func (o *PutRepositoriesPypiGroupByRepositoryNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/group/{repositoryName}][%d] putRepositoriesPypiGroupByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesPypiGroupByRepositoryNameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/pypi/group/{repositoryName}][%d] putRepositoriesPypiGroupByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesPypiGroupByRepositoryNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
