// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRepositoriesYumProxyByRepositorynameReader is a Reader for the PutRepositoriesYumProxyByRepositoryname structure.
type PutRepositoriesYumProxyByRepositorynameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesYumProxyByRepositorynameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesYumProxyByRepositorynameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRepositoriesYumProxyByRepositorynameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRepositoriesYumProxyByRepositorynameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesYumProxyByRepositorynameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/yum/proxy/{repositoryName}] PutRepositoriesYumProxyByRepositoryname", response, response.Code())
	}
}

// NewPutRepositoriesYumProxyByRepositorynameNoContent creates a PutRepositoriesYumProxyByRepositorynameNoContent with default headers values
func NewPutRepositoriesYumProxyByRepositorynameNoContent() *PutRepositoriesYumProxyByRepositorynameNoContent {
	return &PutRepositoriesYumProxyByRepositorynameNoContent{}
}

/*
PutRepositoriesYumProxyByRepositorynameNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutRepositoriesYumProxyByRepositorynameNoContent struct {
}

// IsSuccess returns true when this put repositories yum proxy by repositoryname no content response has a 2xx status code
func (o *PutRepositoriesYumProxyByRepositorynameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put repositories yum proxy by repositoryname no content response has a 3xx status code
func (o *PutRepositoriesYumProxyByRepositorynameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories yum proxy by repositoryname no content response has a 4xx status code
func (o *PutRepositoriesYumProxyByRepositorynameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put repositories yum proxy by repositoryname no content response has a 5xx status code
func (o *PutRepositoriesYumProxyByRepositorynameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories yum proxy by repositoryname no content response a status code equal to that given
func (o *PutRepositoriesYumProxyByRepositorynameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put repositories yum proxy by repositoryname no content response
func (o *PutRepositoriesYumProxyByRepositorynameNoContent) Code() int {
	return 204
}

func (o *PutRepositoriesYumProxyByRepositorynameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/proxy/{repositoryName}][%d] putRepositoriesYumProxyByRepositorynameNoContent", 204)
}

func (o *PutRepositoriesYumProxyByRepositorynameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/proxy/{repositoryName}][%d] putRepositoriesYumProxyByRepositorynameNoContent", 204)
}

func (o *PutRepositoriesYumProxyByRepositorynameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesYumProxyByRepositorynameBadRequest creates a PutRepositoriesYumProxyByRepositorynameBadRequest with default headers values
func NewPutRepositoriesYumProxyByRepositorynameBadRequest() *PutRepositoriesYumProxyByRepositorynameBadRequest {
	return &PutRepositoriesYumProxyByRepositorynameBadRequest{}
}

/*
PutRepositoriesYumProxyByRepositorynameBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutRepositoriesYumProxyByRepositorynameBadRequest struct {
}

// IsSuccess returns true when this put repositories yum proxy by repositoryname bad request response has a 2xx status code
func (o *PutRepositoriesYumProxyByRepositorynameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories yum proxy by repositoryname bad request response has a 3xx status code
func (o *PutRepositoriesYumProxyByRepositorynameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories yum proxy by repositoryname bad request response has a 4xx status code
func (o *PutRepositoriesYumProxyByRepositorynameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories yum proxy by repositoryname bad request response has a 5xx status code
func (o *PutRepositoriesYumProxyByRepositorynameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories yum proxy by repositoryname bad request response a status code equal to that given
func (o *PutRepositoriesYumProxyByRepositorynameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put repositories yum proxy by repositoryname bad request response
func (o *PutRepositoriesYumProxyByRepositorynameBadRequest) Code() int {
	return 400
}

func (o *PutRepositoriesYumProxyByRepositorynameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/proxy/{repositoryName}][%d] putRepositoriesYumProxyByRepositorynameBadRequest", 400)
}

func (o *PutRepositoriesYumProxyByRepositorynameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/proxy/{repositoryName}][%d] putRepositoriesYumProxyByRepositorynameBadRequest", 400)
}

func (o *PutRepositoriesYumProxyByRepositorynameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesYumProxyByRepositorynameUnauthorized creates a PutRepositoriesYumProxyByRepositorynameUnauthorized with default headers values
func NewPutRepositoriesYumProxyByRepositorynameUnauthorized() *PutRepositoriesYumProxyByRepositorynameUnauthorized {
	return &PutRepositoriesYumProxyByRepositorynameUnauthorized{}
}

/*
PutRepositoriesYumProxyByRepositorynameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutRepositoriesYumProxyByRepositorynameUnauthorized struct {
}

// IsSuccess returns true when this put repositories yum proxy by repositoryname unauthorized response has a 2xx status code
func (o *PutRepositoriesYumProxyByRepositorynameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories yum proxy by repositoryname unauthorized response has a 3xx status code
func (o *PutRepositoriesYumProxyByRepositorynameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories yum proxy by repositoryname unauthorized response has a 4xx status code
func (o *PutRepositoriesYumProxyByRepositorynameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories yum proxy by repositoryname unauthorized response has a 5xx status code
func (o *PutRepositoriesYumProxyByRepositorynameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories yum proxy by repositoryname unauthorized response a status code equal to that given
func (o *PutRepositoriesYumProxyByRepositorynameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put repositories yum proxy by repositoryname unauthorized response
func (o *PutRepositoriesYumProxyByRepositorynameUnauthorized) Code() int {
	return 401
}

func (o *PutRepositoriesYumProxyByRepositorynameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/proxy/{repositoryName}][%d] putRepositoriesYumProxyByRepositorynameUnauthorized", 401)
}

func (o *PutRepositoriesYumProxyByRepositorynameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/proxy/{repositoryName}][%d] putRepositoriesYumProxyByRepositorynameUnauthorized", 401)
}

func (o *PutRepositoriesYumProxyByRepositorynameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesYumProxyByRepositorynameForbidden creates a PutRepositoriesYumProxyByRepositorynameForbidden with default headers values
func NewPutRepositoriesYumProxyByRepositorynameForbidden() *PutRepositoriesYumProxyByRepositorynameForbidden {
	return &PutRepositoriesYumProxyByRepositorynameForbidden{}
}

/*
PutRepositoriesYumProxyByRepositorynameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutRepositoriesYumProxyByRepositorynameForbidden struct {
}

// IsSuccess returns true when this put repositories yum proxy by repositoryname forbidden response has a 2xx status code
func (o *PutRepositoriesYumProxyByRepositorynameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories yum proxy by repositoryname forbidden response has a 3xx status code
func (o *PutRepositoriesYumProxyByRepositorynameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories yum proxy by repositoryname forbidden response has a 4xx status code
func (o *PutRepositoriesYumProxyByRepositorynameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories yum proxy by repositoryname forbidden response has a 5xx status code
func (o *PutRepositoriesYumProxyByRepositorynameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories yum proxy by repositoryname forbidden response a status code equal to that given
func (o *PutRepositoriesYumProxyByRepositorynameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put repositories yum proxy by repositoryname forbidden response
func (o *PutRepositoriesYumProxyByRepositorynameForbidden) Code() int {
	return 403
}

func (o *PutRepositoriesYumProxyByRepositorynameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/proxy/{repositoryName}][%d] putRepositoriesYumProxyByRepositorynameForbidden", 403)
}

func (o *PutRepositoriesYumProxyByRepositorynameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/proxy/{repositoryName}][%d] putRepositoriesYumProxyByRepositorynameForbidden", 403)
}

func (o *PutRepositoriesYumProxyByRepositorynameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
