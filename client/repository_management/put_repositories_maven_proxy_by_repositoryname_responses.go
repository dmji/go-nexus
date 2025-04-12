// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRepositoriesMavenProxyByRepositorynameReader is a Reader for the PutRepositoriesMavenProxyByRepositoryname structure.
type PutRepositoriesMavenProxyByRepositorynameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesMavenProxyByRepositorynameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesMavenProxyByRepositorynameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRepositoriesMavenProxyByRepositorynameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRepositoriesMavenProxyByRepositorynameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesMavenProxyByRepositorynameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/maven/proxy/{repositoryName}] PutRepositoriesMavenProxyByRepositoryname", response, response.Code())
	}
}

// NewPutRepositoriesMavenProxyByRepositorynameNoContent creates a PutRepositoriesMavenProxyByRepositorynameNoContent with default headers values
func NewPutRepositoriesMavenProxyByRepositorynameNoContent() *PutRepositoriesMavenProxyByRepositorynameNoContent {
	return &PutRepositoriesMavenProxyByRepositorynameNoContent{}
}

/*
PutRepositoriesMavenProxyByRepositorynameNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutRepositoriesMavenProxyByRepositorynameNoContent struct {
}

// IsSuccess returns true when this put repositories maven proxy by repositoryname no content response has a 2xx status code
func (o *PutRepositoriesMavenProxyByRepositorynameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put repositories maven proxy by repositoryname no content response has a 3xx status code
func (o *PutRepositoriesMavenProxyByRepositorynameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories maven proxy by repositoryname no content response has a 4xx status code
func (o *PutRepositoriesMavenProxyByRepositorynameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put repositories maven proxy by repositoryname no content response has a 5xx status code
func (o *PutRepositoriesMavenProxyByRepositorynameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories maven proxy by repositoryname no content response a status code equal to that given
func (o *PutRepositoriesMavenProxyByRepositorynameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put repositories maven proxy by repositoryname no content response
func (o *PutRepositoriesMavenProxyByRepositorynameNoContent) Code() int {
	return 204
}

func (o *PutRepositoriesMavenProxyByRepositorynameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/proxy/{repositoryName}][%d] putRepositoriesMavenProxyByRepositorynameNoContent", 204)
}

func (o *PutRepositoriesMavenProxyByRepositorynameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/proxy/{repositoryName}][%d] putRepositoriesMavenProxyByRepositorynameNoContent", 204)
}

func (o *PutRepositoriesMavenProxyByRepositorynameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesMavenProxyByRepositorynameBadRequest creates a PutRepositoriesMavenProxyByRepositorynameBadRequest with default headers values
func NewPutRepositoriesMavenProxyByRepositorynameBadRequest() *PutRepositoriesMavenProxyByRepositorynameBadRequest {
	return &PutRepositoriesMavenProxyByRepositorynameBadRequest{}
}

/*
PutRepositoriesMavenProxyByRepositorynameBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutRepositoriesMavenProxyByRepositorynameBadRequest struct {
}

// IsSuccess returns true when this put repositories maven proxy by repositoryname bad request response has a 2xx status code
func (o *PutRepositoriesMavenProxyByRepositorynameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories maven proxy by repositoryname bad request response has a 3xx status code
func (o *PutRepositoriesMavenProxyByRepositorynameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories maven proxy by repositoryname bad request response has a 4xx status code
func (o *PutRepositoriesMavenProxyByRepositorynameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories maven proxy by repositoryname bad request response has a 5xx status code
func (o *PutRepositoriesMavenProxyByRepositorynameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories maven proxy by repositoryname bad request response a status code equal to that given
func (o *PutRepositoriesMavenProxyByRepositorynameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put repositories maven proxy by repositoryname bad request response
func (o *PutRepositoriesMavenProxyByRepositorynameBadRequest) Code() int {
	return 400
}

func (o *PutRepositoriesMavenProxyByRepositorynameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/proxy/{repositoryName}][%d] putRepositoriesMavenProxyByRepositorynameBadRequest", 400)
}

func (o *PutRepositoriesMavenProxyByRepositorynameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/proxy/{repositoryName}][%d] putRepositoriesMavenProxyByRepositorynameBadRequest", 400)
}

func (o *PutRepositoriesMavenProxyByRepositorynameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesMavenProxyByRepositorynameUnauthorized creates a PutRepositoriesMavenProxyByRepositorynameUnauthorized with default headers values
func NewPutRepositoriesMavenProxyByRepositorynameUnauthorized() *PutRepositoriesMavenProxyByRepositorynameUnauthorized {
	return &PutRepositoriesMavenProxyByRepositorynameUnauthorized{}
}

/*
PutRepositoriesMavenProxyByRepositorynameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutRepositoriesMavenProxyByRepositorynameUnauthorized struct {
}

// IsSuccess returns true when this put repositories maven proxy by repositoryname unauthorized response has a 2xx status code
func (o *PutRepositoriesMavenProxyByRepositorynameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories maven proxy by repositoryname unauthorized response has a 3xx status code
func (o *PutRepositoriesMavenProxyByRepositorynameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories maven proxy by repositoryname unauthorized response has a 4xx status code
func (o *PutRepositoriesMavenProxyByRepositorynameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories maven proxy by repositoryname unauthorized response has a 5xx status code
func (o *PutRepositoriesMavenProxyByRepositorynameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories maven proxy by repositoryname unauthorized response a status code equal to that given
func (o *PutRepositoriesMavenProxyByRepositorynameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put repositories maven proxy by repositoryname unauthorized response
func (o *PutRepositoriesMavenProxyByRepositorynameUnauthorized) Code() int {
	return 401
}

func (o *PutRepositoriesMavenProxyByRepositorynameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/proxy/{repositoryName}][%d] putRepositoriesMavenProxyByRepositorynameUnauthorized", 401)
}

func (o *PutRepositoriesMavenProxyByRepositorynameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/proxy/{repositoryName}][%d] putRepositoriesMavenProxyByRepositorynameUnauthorized", 401)
}

func (o *PutRepositoriesMavenProxyByRepositorynameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesMavenProxyByRepositorynameForbidden creates a PutRepositoriesMavenProxyByRepositorynameForbidden with default headers values
func NewPutRepositoriesMavenProxyByRepositorynameForbidden() *PutRepositoriesMavenProxyByRepositorynameForbidden {
	return &PutRepositoriesMavenProxyByRepositorynameForbidden{}
}

/*
PutRepositoriesMavenProxyByRepositorynameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutRepositoriesMavenProxyByRepositorynameForbidden struct {
}

// IsSuccess returns true when this put repositories maven proxy by repositoryname forbidden response has a 2xx status code
func (o *PutRepositoriesMavenProxyByRepositorynameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories maven proxy by repositoryname forbidden response has a 3xx status code
func (o *PutRepositoriesMavenProxyByRepositorynameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories maven proxy by repositoryname forbidden response has a 4xx status code
func (o *PutRepositoriesMavenProxyByRepositorynameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories maven proxy by repositoryname forbidden response has a 5xx status code
func (o *PutRepositoriesMavenProxyByRepositorynameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories maven proxy by repositoryname forbidden response a status code equal to that given
func (o *PutRepositoriesMavenProxyByRepositorynameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put repositories maven proxy by repositoryname forbidden response
func (o *PutRepositoriesMavenProxyByRepositorynameForbidden) Code() int {
	return 403
}

func (o *PutRepositoriesMavenProxyByRepositorynameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/proxy/{repositoryName}][%d] putRepositoriesMavenProxyByRepositorynameForbidden", 403)
}

func (o *PutRepositoriesMavenProxyByRepositorynameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/proxy/{repositoryName}][%d] putRepositoriesMavenProxyByRepositorynameForbidden", 403)
}

func (o *PutRepositoriesMavenProxyByRepositorynameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
