// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRepositoriesCargoProxyByRepositoryNameReader is a Reader for the PutRepositoriesCargoProxyByRepositoryName structure.
type PutRepositoriesCargoProxyByRepositoryNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesCargoProxyByRepositoryNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesCargoProxyByRepositoryNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRepositoriesCargoProxyByRepositoryNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRepositoriesCargoProxyByRepositoryNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesCargoProxyByRepositoryNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/cargo/proxy/{repositoryName}] PutRepositoriesCargoProxyByRepositoryName", response, response.Code())
	}
}

// NewPutRepositoriesCargoProxyByRepositoryNameNoContent creates a PutRepositoriesCargoProxyByRepositoryNameNoContent with default headers values
func NewPutRepositoriesCargoProxyByRepositoryNameNoContent() *PutRepositoriesCargoProxyByRepositoryNameNoContent {
	return &PutRepositoriesCargoProxyByRepositoryNameNoContent{}
}

/*
PutRepositoriesCargoProxyByRepositoryNameNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutRepositoriesCargoProxyByRepositoryNameNoContent struct {
}

// IsSuccess returns true when this put repositories cargo proxy by repository name no content response has a 2xx status code
func (o *PutRepositoriesCargoProxyByRepositoryNameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put repositories cargo proxy by repository name no content response has a 3xx status code
func (o *PutRepositoriesCargoProxyByRepositoryNameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories cargo proxy by repository name no content response has a 4xx status code
func (o *PutRepositoriesCargoProxyByRepositoryNameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put repositories cargo proxy by repository name no content response has a 5xx status code
func (o *PutRepositoriesCargoProxyByRepositoryNameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories cargo proxy by repository name no content response a status code equal to that given
func (o *PutRepositoriesCargoProxyByRepositoryNameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put repositories cargo proxy by repository name no content response
func (o *PutRepositoriesCargoProxyByRepositoryNameNoContent) Code() int {
	return 204
}

func (o *PutRepositoriesCargoProxyByRepositoryNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/cargo/proxy/{repositoryName}][%d] putRepositoriesCargoProxyByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesCargoProxyByRepositoryNameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/cargo/proxy/{repositoryName}][%d] putRepositoriesCargoProxyByRepositoryNameNoContent", 204)
}

func (o *PutRepositoriesCargoProxyByRepositoryNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesCargoProxyByRepositoryNameBadRequest creates a PutRepositoriesCargoProxyByRepositoryNameBadRequest with default headers values
func NewPutRepositoriesCargoProxyByRepositoryNameBadRequest() *PutRepositoriesCargoProxyByRepositoryNameBadRequest {
	return &PutRepositoriesCargoProxyByRepositoryNameBadRequest{}
}

/*
PutRepositoriesCargoProxyByRepositoryNameBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutRepositoriesCargoProxyByRepositoryNameBadRequest struct {
}

// IsSuccess returns true when this put repositories cargo proxy by repository name bad request response has a 2xx status code
func (o *PutRepositoriesCargoProxyByRepositoryNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories cargo proxy by repository name bad request response has a 3xx status code
func (o *PutRepositoriesCargoProxyByRepositoryNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories cargo proxy by repository name bad request response has a 4xx status code
func (o *PutRepositoriesCargoProxyByRepositoryNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories cargo proxy by repository name bad request response has a 5xx status code
func (o *PutRepositoriesCargoProxyByRepositoryNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories cargo proxy by repository name bad request response a status code equal to that given
func (o *PutRepositoriesCargoProxyByRepositoryNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put repositories cargo proxy by repository name bad request response
func (o *PutRepositoriesCargoProxyByRepositoryNameBadRequest) Code() int {
	return 400
}

func (o *PutRepositoriesCargoProxyByRepositoryNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/cargo/proxy/{repositoryName}][%d] putRepositoriesCargoProxyByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesCargoProxyByRepositoryNameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/cargo/proxy/{repositoryName}][%d] putRepositoriesCargoProxyByRepositoryNameBadRequest", 400)
}

func (o *PutRepositoriesCargoProxyByRepositoryNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesCargoProxyByRepositoryNameUnauthorized creates a PutRepositoriesCargoProxyByRepositoryNameUnauthorized with default headers values
func NewPutRepositoriesCargoProxyByRepositoryNameUnauthorized() *PutRepositoriesCargoProxyByRepositoryNameUnauthorized {
	return &PutRepositoriesCargoProxyByRepositoryNameUnauthorized{}
}

/*
PutRepositoriesCargoProxyByRepositoryNameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutRepositoriesCargoProxyByRepositoryNameUnauthorized struct {
}

// IsSuccess returns true when this put repositories cargo proxy by repository name unauthorized response has a 2xx status code
func (o *PutRepositoriesCargoProxyByRepositoryNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories cargo proxy by repository name unauthorized response has a 3xx status code
func (o *PutRepositoriesCargoProxyByRepositoryNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories cargo proxy by repository name unauthorized response has a 4xx status code
func (o *PutRepositoriesCargoProxyByRepositoryNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories cargo proxy by repository name unauthorized response has a 5xx status code
func (o *PutRepositoriesCargoProxyByRepositoryNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories cargo proxy by repository name unauthorized response a status code equal to that given
func (o *PutRepositoriesCargoProxyByRepositoryNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put repositories cargo proxy by repository name unauthorized response
func (o *PutRepositoriesCargoProxyByRepositoryNameUnauthorized) Code() int {
	return 401
}

func (o *PutRepositoriesCargoProxyByRepositoryNameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/cargo/proxy/{repositoryName}][%d] putRepositoriesCargoProxyByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesCargoProxyByRepositoryNameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/cargo/proxy/{repositoryName}][%d] putRepositoriesCargoProxyByRepositoryNameUnauthorized", 401)
}

func (o *PutRepositoriesCargoProxyByRepositoryNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesCargoProxyByRepositoryNameForbidden creates a PutRepositoriesCargoProxyByRepositoryNameForbidden with default headers values
func NewPutRepositoriesCargoProxyByRepositoryNameForbidden() *PutRepositoriesCargoProxyByRepositoryNameForbidden {
	return &PutRepositoriesCargoProxyByRepositoryNameForbidden{}
}

/*
PutRepositoriesCargoProxyByRepositoryNameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutRepositoriesCargoProxyByRepositoryNameForbidden struct {
}

// IsSuccess returns true when this put repositories cargo proxy by repository name forbidden response has a 2xx status code
func (o *PutRepositoriesCargoProxyByRepositoryNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories cargo proxy by repository name forbidden response has a 3xx status code
func (o *PutRepositoriesCargoProxyByRepositoryNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories cargo proxy by repository name forbidden response has a 4xx status code
func (o *PutRepositoriesCargoProxyByRepositoryNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories cargo proxy by repository name forbidden response has a 5xx status code
func (o *PutRepositoriesCargoProxyByRepositoryNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories cargo proxy by repository name forbidden response a status code equal to that given
func (o *PutRepositoriesCargoProxyByRepositoryNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put repositories cargo proxy by repository name forbidden response
func (o *PutRepositoriesCargoProxyByRepositoryNameForbidden) Code() int {
	return 403
}

func (o *PutRepositoriesCargoProxyByRepositoryNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/cargo/proxy/{repositoryName}][%d] putRepositoriesCargoProxyByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesCargoProxyByRepositoryNameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/cargo/proxy/{repositoryName}][%d] putRepositoriesCargoProxyByRepositoryNameForbidden", 403)
}

func (o *PutRepositoriesCargoProxyByRepositoryNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
