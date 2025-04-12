// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRepositoriesHelmHostedByRepositorynameReader is a Reader for the PutRepositoriesHelmHostedByRepositoryname structure.
type PutRepositoriesHelmHostedByRepositorynameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesHelmHostedByRepositorynameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesHelmHostedByRepositorynameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRepositoriesHelmHostedByRepositorynameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRepositoriesHelmHostedByRepositorynameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesHelmHostedByRepositorynameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/helm/hosted/{repositoryName}] PutRepositoriesHelmHostedByRepositoryname", response, response.Code())
	}
}

// NewPutRepositoriesHelmHostedByRepositorynameNoContent creates a PutRepositoriesHelmHostedByRepositorynameNoContent with default headers values
func NewPutRepositoriesHelmHostedByRepositorynameNoContent() *PutRepositoriesHelmHostedByRepositorynameNoContent {
	return &PutRepositoriesHelmHostedByRepositorynameNoContent{}
}

/*
PutRepositoriesHelmHostedByRepositorynameNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutRepositoriesHelmHostedByRepositorynameNoContent struct {
}

// IsSuccess returns true when this put repositories helm hosted by repositoryname no content response has a 2xx status code
func (o *PutRepositoriesHelmHostedByRepositorynameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put repositories helm hosted by repositoryname no content response has a 3xx status code
func (o *PutRepositoriesHelmHostedByRepositorynameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories helm hosted by repositoryname no content response has a 4xx status code
func (o *PutRepositoriesHelmHostedByRepositorynameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put repositories helm hosted by repositoryname no content response has a 5xx status code
func (o *PutRepositoriesHelmHostedByRepositorynameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories helm hosted by repositoryname no content response a status code equal to that given
func (o *PutRepositoriesHelmHostedByRepositorynameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put repositories helm hosted by repositoryname no content response
func (o *PutRepositoriesHelmHostedByRepositorynameNoContent) Code() int {
	return 204
}

func (o *PutRepositoriesHelmHostedByRepositorynameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/helm/hosted/{repositoryName}][%d] putRepositoriesHelmHostedByRepositorynameNoContent", 204)
}

func (o *PutRepositoriesHelmHostedByRepositorynameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/helm/hosted/{repositoryName}][%d] putRepositoriesHelmHostedByRepositorynameNoContent", 204)
}

func (o *PutRepositoriesHelmHostedByRepositorynameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesHelmHostedByRepositorynameBadRequest creates a PutRepositoriesHelmHostedByRepositorynameBadRequest with default headers values
func NewPutRepositoriesHelmHostedByRepositorynameBadRequest() *PutRepositoriesHelmHostedByRepositorynameBadRequest {
	return &PutRepositoriesHelmHostedByRepositorynameBadRequest{}
}

/*
PutRepositoriesHelmHostedByRepositorynameBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutRepositoriesHelmHostedByRepositorynameBadRequest struct {
}

// IsSuccess returns true when this put repositories helm hosted by repositoryname bad request response has a 2xx status code
func (o *PutRepositoriesHelmHostedByRepositorynameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories helm hosted by repositoryname bad request response has a 3xx status code
func (o *PutRepositoriesHelmHostedByRepositorynameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories helm hosted by repositoryname bad request response has a 4xx status code
func (o *PutRepositoriesHelmHostedByRepositorynameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories helm hosted by repositoryname bad request response has a 5xx status code
func (o *PutRepositoriesHelmHostedByRepositorynameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories helm hosted by repositoryname bad request response a status code equal to that given
func (o *PutRepositoriesHelmHostedByRepositorynameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put repositories helm hosted by repositoryname bad request response
func (o *PutRepositoriesHelmHostedByRepositorynameBadRequest) Code() int {
	return 400
}

func (o *PutRepositoriesHelmHostedByRepositorynameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/helm/hosted/{repositoryName}][%d] putRepositoriesHelmHostedByRepositorynameBadRequest", 400)
}

func (o *PutRepositoriesHelmHostedByRepositorynameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/helm/hosted/{repositoryName}][%d] putRepositoriesHelmHostedByRepositorynameBadRequest", 400)
}

func (o *PutRepositoriesHelmHostedByRepositorynameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesHelmHostedByRepositorynameUnauthorized creates a PutRepositoriesHelmHostedByRepositorynameUnauthorized with default headers values
func NewPutRepositoriesHelmHostedByRepositorynameUnauthorized() *PutRepositoriesHelmHostedByRepositorynameUnauthorized {
	return &PutRepositoriesHelmHostedByRepositorynameUnauthorized{}
}

/*
PutRepositoriesHelmHostedByRepositorynameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutRepositoriesHelmHostedByRepositorynameUnauthorized struct {
}

// IsSuccess returns true when this put repositories helm hosted by repositoryname unauthorized response has a 2xx status code
func (o *PutRepositoriesHelmHostedByRepositorynameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories helm hosted by repositoryname unauthorized response has a 3xx status code
func (o *PutRepositoriesHelmHostedByRepositorynameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories helm hosted by repositoryname unauthorized response has a 4xx status code
func (o *PutRepositoriesHelmHostedByRepositorynameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories helm hosted by repositoryname unauthorized response has a 5xx status code
func (o *PutRepositoriesHelmHostedByRepositorynameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories helm hosted by repositoryname unauthorized response a status code equal to that given
func (o *PutRepositoriesHelmHostedByRepositorynameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put repositories helm hosted by repositoryname unauthorized response
func (o *PutRepositoriesHelmHostedByRepositorynameUnauthorized) Code() int {
	return 401
}

func (o *PutRepositoriesHelmHostedByRepositorynameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/helm/hosted/{repositoryName}][%d] putRepositoriesHelmHostedByRepositorynameUnauthorized", 401)
}

func (o *PutRepositoriesHelmHostedByRepositorynameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/helm/hosted/{repositoryName}][%d] putRepositoriesHelmHostedByRepositorynameUnauthorized", 401)
}

func (o *PutRepositoriesHelmHostedByRepositorynameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesHelmHostedByRepositorynameForbidden creates a PutRepositoriesHelmHostedByRepositorynameForbidden with default headers values
func NewPutRepositoriesHelmHostedByRepositorynameForbidden() *PutRepositoriesHelmHostedByRepositorynameForbidden {
	return &PutRepositoriesHelmHostedByRepositorynameForbidden{}
}

/*
PutRepositoriesHelmHostedByRepositorynameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutRepositoriesHelmHostedByRepositorynameForbidden struct {
}

// IsSuccess returns true when this put repositories helm hosted by repositoryname forbidden response has a 2xx status code
func (o *PutRepositoriesHelmHostedByRepositorynameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories helm hosted by repositoryname forbidden response has a 3xx status code
func (o *PutRepositoriesHelmHostedByRepositorynameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories helm hosted by repositoryname forbidden response has a 4xx status code
func (o *PutRepositoriesHelmHostedByRepositorynameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories helm hosted by repositoryname forbidden response has a 5xx status code
func (o *PutRepositoriesHelmHostedByRepositorynameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories helm hosted by repositoryname forbidden response a status code equal to that given
func (o *PutRepositoriesHelmHostedByRepositorynameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put repositories helm hosted by repositoryname forbidden response
func (o *PutRepositoriesHelmHostedByRepositorynameForbidden) Code() int {
	return 403
}

func (o *PutRepositoriesHelmHostedByRepositorynameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/helm/hosted/{repositoryName}][%d] putRepositoriesHelmHostedByRepositorynameForbidden", 403)
}

func (o *PutRepositoriesHelmHostedByRepositorynameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/helm/hosted/{repositoryName}][%d] putRepositoriesHelmHostedByRepositorynameForbidden", 403)
}

func (o *PutRepositoriesHelmHostedByRepositorynameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
