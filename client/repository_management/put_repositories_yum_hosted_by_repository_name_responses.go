// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRepositoriesYumHostedByRepositoryNameReader is a Reader for the PutRepositoriesYumHostedByRepositoryName structure.
type PutRepositoriesYumHostedByRepositoryNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesYumHostedByRepositoryNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesYumHostedByRepositoryNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRepositoriesYumHostedByRepositoryNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRepositoriesYumHostedByRepositoryNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesYumHostedByRepositoryNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/yum/hosted/{repositoryName}] PutRepositoriesYumHostedByRepositoryName", response, response.Code())
	}
}

// NewPutRepositoriesYumHostedByRepositoryNameNoContent creates a PutRepositoriesYumHostedByRepositoryNameNoContent with default headers values
func NewPutRepositoriesYumHostedByRepositoryNameNoContent() *PutRepositoriesYumHostedByRepositoryNameNoContent {
	return &PutRepositoriesYumHostedByRepositoryNameNoContent{}
}

/*
PutRepositoriesYumHostedByRepositoryNameNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutRepositoriesYumHostedByRepositoryNameNoContent struct {
}

// IsSuccess returns true when this put repositories yum hosted by repository name no content response has a 2xx status code
func (o *PutRepositoriesYumHostedByRepositoryNameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put repositories yum hosted by repository name no content response has a 3xx status code
func (o *PutRepositoriesYumHostedByRepositoryNameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories yum hosted by repository name no content response has a 4xx status code
func (o *PutRepositoriesYumHostedByRepositoryNameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put repositories yum hosted by repository name no content response has a 5xx status code
func (o *PutRepositoriesYumHostedByRepositoryNameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories yum hosted by repository name no content response a status code equal to that given
func (o *PutRepositoriesYumHostedByRepositoryNameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put repositories yum hosted by repository name no content response
func (o *PutRepositoriesYumHostedByRepositoryNameNoContent) Code() int {
	return 204
}

func (o *PutRepositoriesYumHostedByRepositoryNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/hosted/{repositoryName}][%d] putRepositoriesYumHostedByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesYumHostedByRepositoryNameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/hosted/{repositoryName}][%d] putRepositoriesYumHostedByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesYumHostedByRepositoryNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesYumHostedByRepositoryNameBadRequest creates a PutRepositoriesYumHostedByRepositoryNameBadRequest with default headers values
func NewPutRepositoriesYumHostedByRepositoryNameBadRequest() *PutRepositoriesYumHostedByRepositoryNameBadRequest {
	return &PutRepositoriesYumHostedByRepositoryNameBadRequest{}
}

/*
PutRepositoriesYumHostedByRepositoryNameBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutRepositoriesYumHostedByRepositoryNameBadRequest struct {
}

// IsSuccess returns true when this put repositories yum hosted by repository name bad request response has a 2xx status code
func (o *PutRepositoriesYumHostedByRepositoryNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories yum hosted by repository name bad request response has a 3xx status code
func (o *PutRepositoriesYumHostedByRepositoryNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories yum hosted by repository name bad request response has a 4xx status code
func (o *PutRepositoriesYumHostedByRepositoryNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories yum hosted by repository name bad request response has a 5xx status code
func (o *PutRepositoriesYumHostedByRepositoryNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories yum hosted by repository name bad request response a status code equal to that given
func (o *PutRepositoriesYumHostedByRepositoryNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put repositories yum hosted by repository name bad request response
func (o *PutRepositoriesYumHostedByRepositoryNameBadRequest) Code() int {
	return 400
}

func (o *PutRepositoriesYumHostedByRepositoryNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/hosted/{repositoryName}][%d] putRepositoriesYumHostedByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesYumHostedByRepositoryNameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/hosted/{repositoryName}][%d] putRepositoriesYumHostedByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesYumHostedByRepositoryNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesYumHostedByRepositoryNameUnauthorized creates a PutRepositoriesYumHostedByRepositoryNameUnauthorized with default headers values
func NewPutRepositoriesYumHostedByRepositoryNameUnauthorized() *PutRepositoriesYumHostedByRepositoryNameUnauthorized {
	return &PutRepositoriesYumHostedByRepositoryNameUnauthorized{}
}

/*
PutRepositoriesYumHostedByRepositoryNameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutRepositoriesYumHostedByRepositoryNameUnauthorized struct {
}

// IsSuccess returns true when this put repositories yum hosted by repository name unauthorized response has a 2xx status code
func (o *PutRepositoriesYumHostedByRepositoryNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories yum hosted by repository name unauthorized response has a 3xx status code
func (o *PutRepositoriesYumHostedByRepositoryNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories yum hosted by repository name unauthorized response has a 4xx status code
func (o *PutRepositoriesYumHostedByRepositoryNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories yum hosted by repository name unauthorized response has a 5xx status code
func (o *PutRepositoriesYumHostedByRepositoryNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories yum hosted by repository name unauthorized response a status code equal to that given
func (o *PutRepositoriesYumHostedByRepositoryNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put repositories yum hosted by repository name unauthorized response
func (o *PutRepositoriesYumHostedByRepositoryNameUnauthorized) Code() int {
	return 401
}

func (o *PutRepositoriesYumHostedByRepositoryNameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/hosted/{repositoryName}][%d] putRepositoriesYumHostedByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesYumHostedByRepositoryNameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/hosted/{repositoryName}][%d] putRepositoriesYumHostedByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesYumHostedByRepositoryNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesYumHostedByRepositoryNameForbidden creates a PutRepositoriesYumHostedByRepositoryNameForbidden with default headers values
func NewPutRepositoriesYumHostedByRepositoryNameForbidden() *PutRepositoriesYumHostedByRepositoryNameForbidden {
	return &PutRepositoriesYumHostedByRepositoryNameForbidden{}
}

/*
PutRepositoriesYumHostedByRepositoryNameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutRepositoriesYumHostedByRepositoryNameForbidden struct {
}

// IsSuccess returns true when this put repositories yum hosted by repository name forbidden response has a 2xx status code
func (o *PutRepositoriesYumHostedByRepositoryNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories yum hosted by repository name forbidden response has a 3xx status code
func (o *PutRepositoriesYumHostedByRepositoryNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories yum hosted by repository name forbidden response has a 4xx status code
func (o *PutRepositoriesYumHostedByRepositoryNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories yum hosted by repository name forbidden response has a 5xx status code
func (o *PutRepositoriesYumHostedByRepositoryNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories yum hosted by repository name forbidden response a status code equal to that given
func (o *PutRepositoriesYumHostedByRepositoryNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put repositories yum hosted by repository name forbidden response
func (o *PutRepositoriesYumHostedByRepositoryNameForbidden) Code() int {
	return 403
}

func (o *PutRepositoriesYumHostedByRepositoryNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/hosted/{repositoryName}][%d] putRepositoriesYumHostedByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesYumHostedByRepositoryNameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/hosted/{repositoryName}][%d] putRepositoriesYumHostedByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesYumHostedByRepositoryNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
