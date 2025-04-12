// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRepositoriesComposerProxyByRepositorynameReader is a Reader for the PutRepositoriesComposerProxyByRepositoryname structure.
type PutRepositoriesComposerProxyByRepositorynameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesComposerProxyByRepositorynameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesComposerProxyByRepositorynameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRepositoriesComposerProxyByRepositorynameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRepositoriesComposerProxyByRepositorynameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesComposerProxyByRepositorynameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/composer/proxy/{repositoryName}] PutRepositoriesComposerProxyByRepositoryname", response, response.Code())
	}
}

// NewPutRepositoriesComposerProxyByRepositorynameNoContent creates a PutRepositoriesComposerProxyByRepositorynameNoContent with default headers values
func NewPutRepositoriesComposerProxyByRepositorynameNoContent() *PutRepositoriesComposerProxyByRepositorynameNoContent {
	return &PutRepositoriesComposerProxyByRepositorynameNoContent{}
}

/*
PutRepositoriesComposerProxyByRepositorynameNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutRepositoriesComposerProxyByRepositorynameNoContent struct {
}

// IsSuccess returns true when this put repositories composer proxy by repositoryname no content response has a 2xx status code
func (o *PutRepositoriesComposerProxyByRepositorynameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put repositories composer proxy by repositoryname no content response has a 3xx status code
func (o *PutRepositoriesComposerProxyByRepositorynameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories composer proxy by repositoryname no content response has a 4xx status code
func (o *PutRepositoriesComposerProxyByRepositorynameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put repositories composer proxy by repositoryname no content response has a 5xx status code
func (o *PutRepositoriesComposerProxyByRepositorynameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories composer proxy by repositoryname no content response a status code equal to that given
func (o *PutRepositoriesComposerProxyByRepositorynameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put repositories composer proxy by repositoryname no content response
func (o *PutRepositoriesComposerProxyByRepositorynameNoContent) Code() int {
	return 204
}

func (o *PutRepositoriesComposerProxyByRepositorynameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/composer/proxy/{repositoryName}][%d] putRepositoriesComposerProxyByRepositorynameNoContent", 204)
}

func (o *PutRepositoriesComposerProxyByRepositorynameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/composer/proxy/{repositoryName}][%d] putRepositoriesComposerProxyByRepositorynameNoContent", 204)
}

func (o *PutRepositoriesComposerProxyByRepositorynameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesComposerProxyByRepositorynameBadRequest creates a PutRepositoriesComposerProxyByRepositorynameBadRequest with default headers values
func NewPutRepositoriesComposerProxyByRepositorynameBadRequest() *PutRepositoriesComposerProxyByRepositorynameBadRequest {
	return &PutRepositoriesComposerProxyByRepositorynameBadRequest{}
}

/*
PutRepositoriesComposerProxyByRepositorynameBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutRepositoriesComposerProxyByRepositorynameBadRequest struct {
}

// IsSuccess returns true when this put repositories composer proxy by repositoryname bad request response has a 2xx status code
func (o *PutRepositoriesComposerProxyByRepositorynameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories composer proxy by repositoryname bad request response has a 3xx status code
func (o *PutRepositoriesComposerProxyByRepositorynameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories composer proxy by repositoryname bad request response has a 4xx status code
func (o *PutRepositoriesComposerProxyByRepositorynameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories composer proxy by repositoryname bad request response has a 5xx status code
func (o *PutRepositoriesComposerProxyByRepositorynameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories composer proxy by repositoryname bad request response a status code equal to that given
func (o *PutRepositoriesComposerProxyByRepositorynameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put repositories composer proxy by repositoryname bad request response
func (o *PutRepositoriesComposerProxyByRepositorynameBadRequest) Code() int {
	return 400
}

func (o *PutRepositoriesComposerProxyByRepositorynameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/composer/proxy/{repositoryName}][%d] putRepositoriesComposerProxyByRepositorynameBadRequest", 400)
}

func (o *PutRepositoriesComposerProxyByRepositorynameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/composer/proxy/{repositoryName}][%d] putRepositoriesComposerProxyByRepositorynameBadRequest", 400)
}

func (o *PutRepositoriesComposerProxyByRepositorynameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesComposerProxyByRepositorynameUnauthorized creates a PutRepositoriesComposerProxyByRepositorynameUnauthorized with default headers values
func NewPutRepositoriesComposerProxyByRepositorynameUnauthorized() *PutRepositoriesComposerProxyByRepositorynameUnauthorized {
	return &PutRepositoriesComposerProxyByRepositorynameUnauthorized{}
}

/*
PutRepositoriesComposerProxyByRepositorynameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutRepositoriesComposerProxyByRepositorynameUnauthorized struct {
}

// IsSuccess returns true when this put repositories composer proxy by repositoryname unauthorized response has a 2xx status code
func (o *PutRepositoriesComposerProxyByRepositorynameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories composer proxy by repositoryname unauthorized response has a 3xx status code
func (o *PutRepositoriesComposerProxyByRepositorynameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories composer proxy by repositoryname unauthorized response has a 4xx status code
func (o *PutRepositoriesComposerProxyByRepositorynameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories composer proxy by repositoryname unauthorized response has a 5xx status code
func (o *PutRepositoriesComposerProxyByRepositorynameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories composer proxy by repositoryname unauthorized response a status code equal to that given
func (o *PutRepositoriesComposerProxyByRepositorynameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put repositories composer proxy by repositoryname unauthorized response
func (o *PutRepositoriesComposerProxyByRepositorynameUnauthorized) Code() int {
	return 401
}

func (o *PutRepositoriesComposerProxyByRepositorynameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/composer/proxy/{repositoryName}][%d] putRepositoriesComposerProxyByRepositorynameUnauthorized", 401)
}

func (o *PutRepositoriesComposerProxyByRepositorynameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/composer/proxy/{repositoryName}][%d] putRepositoriesComposerProxyByRepositorynameUnauthorized", 401)
}

func (o *PutRepositoriesComposerProxyByRepositorynameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesComposerProxyByRepositorynameForbidden creates a PutRepositoriesComposerProxyByRepositorynameForbidden with default headers values
func NewPutRepositoriesComposerProxyByRepositorynameForbidden() *PutRepositoriesComposerProxyByRepositorynameForbidden {
	return &PutRepositoriesComposerProxyByRepositorynameForbidden{}
}

/*
PutRepositoriesComposerProxyByRepositorynameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutRepositoriesComposerProxyByRepositorynameForbidden struct {
}

// IsSuccess returns true when this put repositories composer proxy by repositoryname forbidden response has a 2xx status code
func (o *PutRepositoriesComposerProxyByRepositorynameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories composer proxy by repositoryname forbidden response has a 3xx status code
func (o *PutRepositoriesComposerProxyByRepositorynameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories composer proxy by repositoryname forbidden response has a 4xx status code
func (o *PutRepositoriesComposerProxyByRepositorynameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories composer proxy by repositoryname forbidden response has a 5xx status code
func (o *PutRepositoriesComposerProxyByRepositorynameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories composer proxy by repositoryname forbidden response a status code equal to that given
func (o *PutRepositoriesComposerProxyByRepositorynameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put repositories composer proxy by repositoryname forbidden response
func (o *PutRepositoriesComposerProxyByRepositorynameForbidden) Code() int {
	return 403
}

func (o *PutRepositoriesComposerProxyByRepositorynameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/composer/proxy/{repositoryName}][%d] putRepositoriesComposerProxyByRepositorynameForbidden", 403)
}

func (o *PutRepositoriesComposerProxyByRepositorynameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/composer/proxy/{repositoryName}][%d] putRepositoriesComposerProxyByRepositorynameForbidden", 403)
}

func (o *PutRepositoriesComposerProxyByRepositorynameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
