// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRepositoriesCargoProxyByRepositorynameReader is a Reader for the PutRepositoriesCargoProxyByRepositoryname structure.
type PutRepositoriesCargoProxyByRepositorynameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesCargoProxyByRepositorynameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesCargoProxyByRepositorynameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRepositoriesCargoProxyByRepositorynameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRepositoriesCargoProxyByRepositorynameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesCargoProxyByRepositorynameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/cargo/proxy/{repositoryName}] PutRepositoriesCargoProxyByRepositoryname", response, response.Code())
	}
}

// NewPutRepositoriesCargoProxyByRepositorynameNoContent creates a PutRepositoriesCargoProxyByRepositorynameNoContent with default headers values
func NewPutRepositoriesCargoProxyByRepositorynameNoContent() *PutRepositoriesCargoProxyByRepositorynameNoContent {
	return &PutRepositoriesCargoProxyByRepositorynameNoContent{}
}

/*
PutRepositoriesCargoProxyByRepositorynameNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutRepositoriesCargoProxyByRepositorynameNoContent struct {
}

// IsSuccess returns true when this put repositories cargo proxy by repositoryname no content response has a 2xx status code
func (o *PutRepositoriesCargoProxyByRepositorynameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put repositories cargo proxy by repositoryname no content response has a 3xx status code
func (o *PutRepositoriesCargoProxyByRepositorynameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories cargo proxy by repositoryname no content response has a 4xx status code
func (o *PutRepositoriesCargoProxyByRepositorynameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put repositories cargo proxy by repositoryname no content response has a 5xx status code
func (o *PutRepositoriesCargoProxyByRepositorynameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories cargo proxy by repositoryname no content response a status code equal to that given
func (o *PutRepositoriesCargoProxyByRepositorynameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put repositories cargo proxy by repositoryname no content response
func (o *PutRepositoriesCargoProxyByRepositorynameNoContent) Code() int {
	return 204
}

func (o *PutRepositoriesCargoProxyByRepositorynameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/cargo/proxy/{repositoryName}][%d] putRepositoriesCargoProxyByRepositorynameNoContent", 204)
}

func (o *PutRepositoriesCargoProxyByRepositorynameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/cargo/proxy/{repositoryName}][%d] putRepositoriesCargoProxyByRepositorynameNoContent", 204)
}

func (o *PutRepositoriesCargoProxyByRepositorynameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesCargoProxyByRepositorynameBadRequest creates a PutRepositoriesCargoProxyByRepositorynameBadRequest with default headers values
func NewPutRepositoriesCargoProxyByRepositorynameBadRequest() *PutRepositoriesCargoProxyByRepositorynameBadRequest {
	return &PutRepositoriesCargoProxyByRepositorynameBadRequest{}
}

/*
PutRepositoriesCargoProxyByRepositorynameBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutRepositoriesCargoProxyByRepositorynameBadRequest struct {
}

// IsSuccess returns true when this put repositories cargo proxy by repositoryname bad request response has a 2xx status code
func (o *PutRepositoriesCargoProxyByRepositorynameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories cargo proxy by repositoryname bad request response has a 3xx status code
func (o *PutRepositoriesCargoProxyByRepositorynameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories cargo proxy by repositoryname bad request response has a 4xx status code
func (o *PutRepositoriesCargoProxyByRepositorynameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories cargo proxy by repositoryname bad request response has a 5xx status code
func (o *PutRepositoriesCargoProxyByRepositorynameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories cargo proxy by repositoryname bad request response a status code equal to that given
func (o *PutRepositoriesCargoProxyByRepositorynameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put repositories cargo proxy by repositoryname bad request response
func (o *PutRepositoriesCargoProxyByRepositorynameBadRequest) Code() int {
	return 400
}

func (o *PutRepositoriesCargoProxyByRepositorynameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/cargo/proxy/{repositoryName}][%d] putRepositoriesCargoProxyByRepositorynameBadRequest", 400)
}

func (o *PutRepositoriesCargoProxyByRepositorynameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/cargo/proxy/{repositoryName}][%d] putRepositoriesCargoProxyByRepositorynameBadRequest", 400)
}

func (o *PutRepositoriesCargoProxyByRepositorynameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesCargoProxyByRepositorynameUnauthorized creates a PutRepositoriesCargoProxyByRepositorynameUnauthorized with default headers values
func NewPutRepositoriesCargoProxyByRepositorynameUnauthorized() *PutRepositoriesCargoProxyByRepositorynameUnauthorized {
	return &PutRepositoriesCargoProxyByRepositorynameUnauthorized{}
}

/*
PutRepositoriesCargoProxyByRepositorynameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutRepositoriesCargoProxyByRepositorynameUnauthorized struct {
}

// IsSuccess returns true when this put repositories cargo proxy by repositoryname unauthorized response has a 2xx status code
func (o *PutRepositoriesCargoProxyByRepositorynameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories cargo proxy by repositoryname unauthorized response has a 3xx status code
func (o *PutRepositoriesCargoProxyByRepositorynameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories cargo proxy by repositoryname unauthorized response has a 4xx status code
func (o *PutRepositoriesCargoProxyByRepositorynameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories cargo proxy by repositoryname unauthorized response has a 5xx status code
func (o *PutRepositoriesCargoProxyByRepositorynameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories cargo proxy by repositoryname unauthorized response a status code equal to that given
func (o *PutRepositoriesCargoProxyByRepositorynameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put repositories cargo proxy by repositoryname unauthorized response
func (o *PutRepositoriesCargoProxyByRepositorynameUnauthorized) Code() int {
	return 401
}

func (o *PutRepositoriesCargoProxyByRepositorynameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/cargo/proxy/{repositoryName}][%d] putRepositoriesCargoProxyByRepositorynameUnauthorized", 401)
}

func (o *PutRepositoriesCargoProxyByRepositorynameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/cargo/proxy/{repositoryName}][%d] putRepositoriesCargoProxyByRepositorynameUnauthorized", 401)
}

func (o *PutRepositoriesCargoProxyByRepositorynameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesCargoProxyByRepositorynameForbidden creates a PutRepositoriesCargoProxyByRepositorynameForbidden with default headers values
func NewPutRepositoriesCargoProxyByRepositorynameForbidden() *PutRepositoriesCargoProxyByRepositorynameForbidden {
	return &PutRepositoriesCargoProxyByRepositorynameForbidden{}
}

/*
PutRepositoriesCargoProxyByRepositorynameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutRepositoriesCargoProxyByRepositorynameForbidden struct {
}

// IsSuccess returns true when this put repositories cargo proxy by repositoryname forbidden response has a 2xx status code
func (o *PutRepositoriesCargoProxyByRepositorynameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories cargo proxy by repositoryname forbidden response has a 3xx status code
func (o *PutRepositoriesCargoProxyByRepositorynameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories cargo proxy by repositoryname forbidden response has a 4xx status code
func (o *PutRepositoriesCargoProxyByRepositorynameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories cargo proxy by repositoryname forbidden response has a 5xx status code
func (o *PutRepositoriesCargoProxyByRepositorynameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories cargo proxy by repositoryname forbidden response a status code equal to that given
func (o *PutRepositoriesCargoProxyByRepositorynameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put repositories cargo proxy by repositoryname forbidden response
func (o *PutRepositoriesCargoProxyByRepositorynameForbidden) Code() int {
	return 403
}

func (o *PutRepositoriesCargoProxyByRepositorynameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/cargo/proxy/{repositoryName}][%d] putRepositoriesCargoProxyByRepositorynameForbidden", 403)
}

func (o *PutRepositoriesCargoProxyByRepositorynameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/cargo/proxy/{repositoryName}][%d] putRepositoriesCargoProxyByRepositorynameForbidden", 403)
}

func (o *PutRepositoriesCargoProxyByRepositorynameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
